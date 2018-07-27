package emulator

func (e *Emulator) PushIntoStack(value uint16) {
	low := uint8(value & 0xFF)
	high := uint8(value >> 8)
	e.StackPointer.Decrement()
	e.WriteMemory(e.StackPointer.Value(), high)
	e.StackPointer.Decrement()
	e.WriteMemory(e.StackPointer.Value(), low)
}

func (e *Emulator) PopFromStack() uint16 {
	value := e.ReadMemory16Bit(e.StackPointer.Value())
	e.StackPointer.Increment()
	e.StackPointer.Increment()
	return value
}
