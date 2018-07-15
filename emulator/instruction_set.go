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
		testPanic(addend == 0, "TODO: Verify what happens in specification when this overflows. What's the right order?")
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

func (e *Emulator) CPU8BitSubtract(r1 *Register8Bit, minuend uint8, useCarry bool) int {
	subtrahend := r1.Value()
	if useCarry && e.FlagC() {
		minuend++
		testPanic(minuend == 0, "TODO: Verify what happens in specification when this overflows. What's the right order?")
	}
	result := uint16(subtrahend&0xFF) - uint16(minuend&0xFF)
	r1.SetValue(uint8(result & 0xFF))
	e.ClearAllFlags()
	if r1.Value() == 0 {
		e.SetFlagZ()
	}
	e.SetFlagN()
	if int16(subtrahend&0xF)-int16(minuend&0xF) < 0x0 {
		e.SetFlagH()
	}
	if subtrahend < minuend {
		e.SetFlagC()
	}
	return 4
}

func (e *Emulator) CPU8BitAnd(operand uint8) int {
	result := e.AF.High.Value() & operand
	e.AF.SetHigh(result)
	e.ClearAllFlags()
	if e.AF.High.Value() == 0x0 {
		e.SetFlagZ()
	}
	e.SetFlagH()
	return 4
}

func (e *Emulator) CPU8BitOr(operand uint8) int {
	result := e.AF.High.Value() | operand
	e.AF.High.SetValue(result)
	e.ClearAllFlags()
	if e.AF.High.Value() == 0x0 {
		e.SetFlagZ()
	}
	return 4
}

func (e *Emulator) CPU8BitXor(operand uint8) int {
	result := e.AF.High.Value() ^ operand
	e.AF.High.SetValue(result)
	e.ClearAllFlags()
	if e.AF.High.Value() == 0x0 {
		e.SetFlagZ()
	}
	return 4
}

func (e *Emulator) CPU8BitCompare(operand uint8) int {
	result := e.AF.High.Value() - operand
	e.ClearAllFlags()
	if result == 0x0 {
		e.SetFlagZ()
	}
	e.SetFlagN()
	if int16(e.AF.High.Value()&0xF)-int16(operand&0xF) < 0x0 {
		e.SetFlagH()
	}
	if e.AF.High.Value() < operand {
		e.SetFlagC()
	}
	return 4
}

func (e *Emulator) CPU8BitIncrement(r *Register8Bit) int {
	previous := r.Value()
	r.SetValue(r.Value() + 1)
	if r.Value() == 0x0 {
		e.SetFlagZ()
	} else {
		e.ClearFlagZ()
	}
	e.ClearFlagN()
	if previous&0xF == 0xF {
		e.SetFlagH()
	} else {
		e.ClearFlagH()
	}
	return 4
}

func (e *Emulator) CPU8BitIncrementMemoryAddress(address uint16) int {
	previous := e.ReadMemory8Bit(address)
	current := previous + 1
	e.WriteMemory(address, current)
	if current == 0x0 {
		e.SetFlagZ()
	} else {
		e.ClearFlagZ()
	}
	e.ClearFlagN()
	if previous&0xF == 0xF {
		e.SetFlagH()
	} else {
		e.ClearFlagH()
	}
	return 12
}

func (e *Emulator) CPU8BitDecrement(r *Register8Bit) int {
	previous := r.Value()
	r.SetValue(r.Value() - 1)
	if r.Value() == 0x0 {
		e.SetFlagZ()
	} else {
		e.ClearFlagZ()
	}
	e.SetFlagN()
	if previous&0x0F == 0x0 {
		e.SetFlagH()
	} else {
		e.ClearFlagH()
	}
	return 4
}

func (e *Emulator) CPU8BitDecrementMemoryAddress(address uint16) int {
	previous := e.ReadMemory8Bit(address)
	current := previous - 1
	e.WriteMemory(address, current)
	if current == 0x0 {
		e.SetFlagZ()
	} else {
		e.ClearFlagZ()
	}
	e.SetFlagN()
	if previous&0x0F == 0x0 {
		e.SetFlagH()
	} else {
		e.ClearFlagH()
	}
	return 12
}

func (e *Emulator) CPU16BitAdd(r1 *Register16Bit, r2 Register16Bit) int {
	augend := r1.Value()
	addend := r2.Value()
	result := uint32(augend&0xFFFF) + uint32(addend&0xFFFF)
	r1.SetValue(uint16(result & 0xFFFF))
	e.ClearFlagN()
	if result > 0xFFFF {
		e.SetFlagC()
	} else {
		e.ClearFlagC()
	}
	if ((augend&0xFF00)&0xF)+((addend>>8)&0xF) != 0 {
		e.SetFlagH()
	} else {
		e.ClearFlagH()
	}
	return 8
}
