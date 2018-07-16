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

func (e *Emulator) CPU16BitRegisterMemoryLoad(r *Register16Bit) int {
	value := e.ReadMemory16Bit(e.ProgramCounter.Value())
	e.ProgramCounter.Increment()
	e.ProgramCounter.Increment()
	r.SetValue(value)
	return 12
}

func (e *Emulator) CPU8BitRegisterAdd(r1 *Register8Bit, addend uint8, useCarry bool) int {
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

func (e *Emulator) CPU8BitRegisterSubtract(r1 *Register8Bit, minuend uint8, useCarry bool) int {
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

func (e *Emulator) CPU8BitRegisterIncrement(r *Register8Bit) int {
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

func (e *Emulator) CPU8BitRegisterDecrement(r *Register8Bit) int {
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

func (e *Emulator) CPU16BitRegisterAdd(r1 *Register16Bit, r2 Register16Bit) int {
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

func (e *Emulator) CPU16BitRegisterIncrement(r *Register16Bit) int {
	r.SetValue(r.Value() + 1)
	return 8
}

func (e *Emulator) CPU16BitRegisterDecrement(r *Register16Bit) int {
	r.SetValue(r.Value() - 1)
	return 8
}

func (e *Emulator) CPU8BitRegisterSwap(r *Register8Bit) int {
	result := (r.Value()&0xF0)>>4 | (r.Value()&0x0F)<<4
	r.SetValue(result)
	e.ClearAllFlags()
	if result == 0x0 {
		e.SetFlagZ()
	}
	return 8
}

func (e *Emulator) CPU8BitSwapMemoryAddress(address uint16) int {
	value := e.ReadMemory8Bit(address)
	result := (value&0xF0)>>4 | (value&0x0F)<<4
	e.WriteMemory(address, result)
	e.ClearAllFlags()
	if result == 0x0 {
		e.SetFlagZ()
	}
	return 16
}

// There's no much information about this on the Game Boy CPU Manual
// This implementation is adapted from other emulators.
func (e *Emulator) CPUDDA() int {
	value := uint16(e.AF.High.Value())
	if e.FlagN() {
		if e.FlagH() {
			value = (value - 0x06) & 0xFF
		}
		if e.FlagC() {
			value = value - 0x60
		}
	} else {
		if e.FlagH() || ((value & 0x0F) > 9) {
			value = value + 0x06
		}
		if e.FlagC() || value > 0x9F {
			value = value + 0x60
		}
	}
	if value >= 0x100 {
		e.SetFlagC()
	}
	e.AF.High.SetValue(uint8(value & 0xFF))
	if e.AF.High.Value() == 0 {
		e.SetFlagZ()
	} else {
		e.ClearFlagZ()
	}
	e.ClearFlagH()
	return 4
}

func (e *Emulator) CPU8BitRegisterRLC(r *Register8Bit) int {
	test := testBit(r.Value(), 7)
	r.SetValue(r.Value() << 1)
	e.ClearAllFlags()
	if test {
		e.SetFlagC()
		r.SetValue(setBit(r.Value(), 0))
	}
	if r.Value() == 0x0 {
		e.SetFlagZ()
	}
	return 8
}

func (e *Emulator) CPU8BitRegisterRL(r *Register8Bit) int {
	testCarry := e.FlagC()
	test := testBit(r.Value(), 7)
	r.SetValue(r.Value() << 1)
	e.ClearAllFlags()
	if test {
		e.SetFlagC()
	}
	if testCarry {
		r.SetValue(setBit(r.Value(), 0))
	}
	if r.Value() == 0x0 {
		e.SetFlagZ()
	}
	return 8
}

func (e *Emulator) CPU8BitRegisterRRC(r *Register8Bit) int {
	test := testBit(r.Value(), 0)
	r.SetValue(r.Value() >> 1)
	e.ClearAllFlags()
	if test {
		e.SetFlagC()
		r.SetValue(setBit(r.Value(), 7))
	}
	if r.Value() == 0x0 {
		e.SetFlagZ()
	}
	return 8
}

func (e *Emulator) CPU8BitRegisterRR(r *Register8Bit) int {
	testCarry := e.FlagC()
	test := testBit(r.Value(), 0)
	r.SetValue(r.Value() >> 1)
	e.ClearAllFlags()
	if test {
		e.SetFlagC()
	}
	if testCarry {
		r.SetValue(setBit(r.Value(), 7))
	}
	if r.Value() == 0x0 {
		e.SetFlagZ()
	}
	return 8
}

func (e *Emulator) CPU8BitRLCMemoryAddress(address uint16) int {
	value := e.ReadMemory8Bit(address)
	test := testBit(value, 7)
	value <<= 1
	e.ClearAllFlags()
	if test {
		e.SetFlagC()
		value = setBit(value, 0)
	}
	if value == 0x0 {
		e.SetFlagZ()
	}
	e.WriteMemory(address, value)
	return 16
}

func (e *Emulator) CPU8BitRLMemoryAddress(address uint16) int {
	value := e.ReadMemory8Bit(address)
	testCarry := e.FlagC()
	test := testBit(value, 7)
	value <<= 1
	e.ClearAllFlags()
	if test {
		e.SetFlagC()
	}
	if testCarry {
		value = setBit(value, 0)
	}
	if value == 0x0 {
		e.SetFlagZ()
	}
	e.WriteMemory(address, value)
	return 16
}

func (e *Emulator) CPU8BitRRCMemoryAddress(address uint16) int {
	value := e.ReadMemory8Bit(address)
	test := testBit(value, 0)
	value >>= 1
	e.ClearAllFlags()
	if test {
		e.SetFlagC()
		value = setBit(value, 7)
	}
	if value == 0x0 {
		e.SetFlagZ()
	}
	e.WriteMemory(address, value)
	return 16
}

func (e *Emulator) CPU8BitRRMemoryAddress(address uint16) int {
	value := e.ReadMemory8Bit(address)
	testCarry := e.FlagC()
	test := testBit(value, 0)
	value >>= 1
	e.ClearAllFlags()
	if test {
		e.SetFlagC()
	}
	if testCarry {
		value = setBit(value, 7)
	}
	if value == 0x0 {
		e.SetFlagZ()
	}
	e.WriteMemory(address, value)
	return 16
}

func (e *Emulator) CPU8BitRegisterSLA(r *Register8Bit) int {
	test := testBit(r.Value(), 7)
	r.SetValue(r.Value() << 1)
	e.ClearAllFlags()
	if test {
		e.SetFlagC()
	}
	if r.Value() == 0x0 {
		e.SetFlagZ()
	}
	return 8
}

func (e *Emulator) CPU8BitSLAMemoryAddress(address uint16) int {
	value := e.ReadMemory8Bit(address)
	test := testBit(value, 7)
	value <<= 1
	e.ClearAllFlags()
	if test {
		e.SetFlagC()
	}
	if value == 0x0 {
		e.SetFlagZ()
	}
	e.WriteMemory(address, value)
	return 16
}

func (e *Emulator) CPU8BitRegisterSRA(r *Register8Bit) int {
	testMSB := testBit(r.Value(), 7)
	testLSB := testBit(r.Value(), 0)
	r.SetValue(r.Value() >> 1)
	e.ClearAllFlags()
	if testMSB {
		r.SetValue(setBit(r.Value(), 7))
	}
	if testLSB {
		e.SetFlagC()
	}
	if r.Value() == 0x0 {
		e.SetFlagZ()
	}
	return 8
}

func (e *Emulator) CPU8BitSRAMemoryAddress(address uint16) int {
	value := e.ReadMemory8Bit(address)
	testMSB := testBit(value, 7)
	testLSB := testBit(value, 0)
	value >>= 1
	e.ClearAllFlags()
	if testMSB {
		value = setBit(value, 7)
	}
	if testLSB {
		e.SetFlagC()
	}
	if value == 0x0 {
		e.SetFlagZ()
	}
	e.WriteMemory(address, value)
	return 16
}

func (e *Emulator) CPU8BitRegisterSRL(r *Register8Bit) int {
	test := testBit(r.Value(), 0)
	r.SetValue(r.Value() >> 1)
	e.ClearAllFlags()
	if test {
		e.SetFlagC()
	}
	if r.Value() == 0x0 {
		e.SetFlagZ()
	}
	return 8
}

func (e *Emulator) CPU8BitSRLMemoryAddress(address uint16) int {
	value := e.ReadMemory8Bit(address)
	test := testBit(value, 0)
	value >>= 1
	e.ClearAllFlags()
	if test {
		e.SetFlagC()
	}
	if value == 0x0 {
		e.SetFlagZ()
	}
	e.WriteMemory(address, value)
	return 16
}

func (e *Emulator) CPU8BitRegisterBit(r *Register8Bit, position uint) int {
	test := testBit(r.Value(), position)
	if test {
		e.ClearFlagZ()
	} else {
		e.SetFlagZ()
	}
	e.ClearFlagN()
	e.SetFlagH()
	return 8
}

func (e *Emulator) CPU8BitBitMemoryAddress(address uint16, position uint) int {
	value := e.ReadMemory8Bit(address)
	test := testBit(value, position)
	if test {
		e.SetFlagZ()
	} else {
		e.ClearFlagZ()
	}
	e.ClearFlagN()
	e.SetFlagH()
	return 16
}
