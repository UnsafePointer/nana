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

func (e *Emulator) CPU16BitRegistryMemoryLoad(r *Register16Bit) int {
	value := e.ReadMemory16Bit(e.ProgramCounter.Value())
	e.ProgramCounter.Increment()
	e.ProgramCounter.Increment()
	r.SetValue(value)
	return 12
}

func (e *Emulator) CPU8BitAdd(r1 *Register8Bit, addend uint8, useCarry bool) int {
	augend := r1.Value()
	if useCarry && e.FlagC() {
		addend++
		if addend == 0 {
			panic("TODO: Verify what happens in specification when this overflows. What's the right order?")
		}
	}
	result := uint16(augend&0xFF) + uint16(addend&0xFF)
	r1.SetValue(uint8(result & 0xFF))
	e.ClearAllFlags()
	if r1.Value() == 0 {
		e.SetFlagZ()
	}
	if uint8(augend&0xF)+uint8(addend&0xF) > 0xF {
		e.SetFlagH()
	}
	if result > 0xFF {
		e.SetFlagC()
	}
	return 4
}
