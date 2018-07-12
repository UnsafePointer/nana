package emulator

func (e *Emulator) CPU8BitLoad(r *Register, high bool) int {
	value := e.ReadMemory(e.ProgramCounter)
	if high {
		r.SetHigh(value)
	} else {
		r.SetLow(value)
	}
	e.ProgramCounter++
	return 8
}
