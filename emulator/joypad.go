package emulator

import "github.com/veandco/go-sdl2/sdl"

// Taken from documentation
// Bit 7 - Not used
// Bit 6 - Not used
// Bit 5 - P15 Select Button Keys (0=Select)
// Bit 4 - P14 Select Direction Keys (0=Select)
// Bit 3 - P13 Input Down or Start (0=Pressed) (Read Only)
// Bit 2 - P12 Input Up or Select (0=Pressed) (Read Only)
// Bit 1 - P11 Input Left or Button B (0=Pressed) (Read Only)
// Bit 0 - P10 Input Right or Button A (0=Pressed) (Read Only)

const joypadRegister = 0xFF00

// If a button in the joypad is pressed it's corresponding
// bit in JoypadState is set to 0
// RIGHT -> 0
// LEFT -> 1
// UP -> 2
// DOWN -> 3
// A -> 4
// B (S Key) -> 5
// SELECT (Space Key) -> 6
// ENTER -> 7

func (e *Emulator) PressKey(key uint) {
	isKeyPressed := !testBit(e.JoypadState, key)
	e.JoypadState = clearBit(e.JoypadState, key)

	isButton := false
	if key > 3 {
		isButton = true
	}

	requested := e.ROM[joypadRegister]
	requestInterrupt := false

	if isButton && !testBit(requested, 5) {
		requestInterrupt = true
	} else if !isButton && !testBit(requested, 4) {
		requestInterrupt = true
	}

	if requestInterrupt && !isKeyPressed {
		e.RequestInterrupt(4)
	}
}

func (e *Emulator) ReleaseKey(key uint) {
	e.JoypadState = setBit(e.JoypadState, key)
}

func (e *Emulator) HandleKeyboardEvent(keyboardEvent *sdl.KeyboardEvent) {
	switch keyboardEvent.State {
	case sdl.PRESSED:
		e.PressKey(e.KeyMap[keyboardEvent.Keysym.Sym])
		break
	case sdl.RELEASED:
		e.ReleaseKey(e.KeyMap[keyboardEvent.Keysym.Sym])
		break
	}
}

func (e *Emulator) GetJoypadState() uint8 {
	register := e.ROM[joypadRegister]
	register ^= 0xFF

	if !testBit(register, 4) {
		top := e.JoypadState >> 4
		top |= 0xF0
		register &= top
	} else if !testBit(register, 5) {
		bottom := e.JoypadState & 0xF
		bottom |= 0xF0
		register &= bottom
	}

	return register
}
