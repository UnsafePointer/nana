package emulator

func (e *Emulator) CPU8BitLoad(r *Register8Bit) int {
	value := e.ReadMemory(e.ProgramCounter.Value())
	r.SetValue(value)
	e.ProgramCounter.Increment()
	return 8
}

func (e *Emulator) CPURegLoad(r *Register8Bit, v Register8Bit) int {
	r.SetValue(v.Value())
	return 4
}
