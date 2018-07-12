package emulator

func (e *Emulator) CPU8BitRegisterMemoryLoad(r *Register8Bit) int {
	value := e.ReadMemory(e.ProgramCounter.Value())
	r.SetValue(value)
	e.ProgramCounter.Increment()
	return 8
}

func (e Emulator) CPU8BitRegisterLoad(r *Register8Bit, v Register8Bit) int {
	r.SetValue(v.Value())
	return 4
}
