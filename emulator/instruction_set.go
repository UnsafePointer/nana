package emulator

func (e *Emulator) CPU8BitRegisterMemoryLoad(r *Register8Bit) int {
	value := e.ReadMemory8Bit(e.ProgramCounter.Value())
	r.SetValue(value)
	e.ProgramCounter.Increment()
	return 8
}

func (e Emulator) CPU8BitRegisterLoad(r *Register8Bit, v Register8Bit) int {
	r.SetValue(v.Value())
	return 4
}

func (e *Emulator) CPU8BitRegisterMemoryWrite(address uint16, v Register8Bit) int {
	e.WriteMemory(address, v.Value())
	return 8
}

func (e Emulator) CPU8BitRegisterMemoryAddressLoad(r *Register8Bit, address uint16) int {
	r.SetValue(e.ReadMemory8Bit(address))
	return 8
}
