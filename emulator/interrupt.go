package emulator

const interruptRequestRegisterAddress = 0xFF0F
const InterruptEnabledRegisterAddress = 0xFFFF

func (e *Emulator) RequestInterrupt(interruptId uint) {
	value := e.ReadMemory8Bit(interruptRequestRegisterAddress)
	value = setBit(value, interruptId)
	e.WriteMemory(interruptRequestRegisterAddress, value)
}

func (e *Emulator) ClearRequestInterrupt(interruptId uint) {
	value := e.ReadMemory8Bit(interruptRequestRegisterAddress)
	value = clearBit(value, interruptId)
	e.WriteMemory(interruptRequestRegisterAddress, value)
}

func (e *Emulator) ExecuteInterrupts() {
	if e.DisableInterrupts {
		return
	}
	interruptRequest := e.ReadMemory8Bit(interruptRequestRegisterAddress)
	if interruptRequest > 0 {
		for bitPosition := uint(0); bitPosition < 8; bitPosition++ {
			if testBit(interruptRequest, bitPosition) {
				interruptEnabled := e.ReadMemory8Bit(InterruptEnabledRegisterAddress)
				if testBit(interruptEnabled, bitPosition) {
					e.executeInterrupt(bitPosition)
				}
			}
		}
	}
}

func (e *Emulator) executeInterrupt(interruptId uint) {
	e.PushIntoStack(e.ProgramCounter.Value())
	e.Halted = false

	// Interrupt routine map
	// V-Blank: 0x40
	// LCD: 0x48
	// TIMER: 0x50
	// JOYPAD: 0x60
	switch interruptId {
	case 0:
		e.ProgramCounter.SetValue(0x40)
		break
	case 1:
		e.ProgramCounter.SetValue(0x48)
		break
	case 2:
		e.ProgramCounter.SetValue(0x50)
		break
	case 3:
		e.ProgramCounter.SetValue(0x60)
		break
	}

	e.DisableInterrupts = true
	e.ClearRequestInterrupt(interruptId)
}
