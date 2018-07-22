package emulator

import "fmt"

func (e *Emulator) PushIntoStack(value uint16) {
	e.LogMessage(fmt.Sprintf("Push into stack: %#X", value))
	low := uint8(value & 0xFF)
	high := uint8(value >> 8)
	e.StackPointer.Decrement()
	e.WriteMemory(e.StackPointer.Value(), high)
	e.StackPointer.Decrement()
	e.WriteMemory(e.StackPointer.Value(), low)
}

func (e *Emulator) PopFromStack() uint16 {
	value := e.ReadMemory16Bit(e.StackPointer.Value())
	e.LogMessage(fmt.Sprintf("Pop from stack: %#X", value))
	e.StackPointer.Increment()
	e.StackPointer.Increment()
	return value
}
