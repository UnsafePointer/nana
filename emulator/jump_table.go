package emulator

func (e *Emulator) ExecuteOpCode(opcode uint8) int {
	switch opcode {
	// no-op
	case 0x00:
		return 4
	// 8-Bit Loads
	// LD nn,n
	case 0x06:
		return e.CPU8BitRegisterMemoryLoad(&e.BC.High)
	case 0x0E:
		return e.CPU8BitRegisterMemoryLoad(&e.BC.Low)
	case 0x16:
		return e.CPU8BitRegisterMemoryLoad(&e.DE.High)
	case 0x1E:
		return e.CPU8BitRegisterMemoryLoad(&e.DE.Low)
	case 0x26:
		return e.CPU8BitRegisterMemoryLoad(&e.HL.High)
	case 0x2E:
		return e.CPU8BitRegisterMemoryLoad(&e.HL.Low)
	// LD r1,r2
	case 0x7F:
		return e.CPU8BitRegisterLoad(&e.AF.High, e.AF.High)
	case 0x78:
		return e.CPU8BitRegisterLoad(&e.AF.High, e.BC.High)
	case 0x79:
		return e.CPU8BitRegisterLoad(&e.AF.High, e.BC.Low)
	case 0x7A:
		return e.CPU8BitRegisterLoad(&e.AF.High, e.DE.High)
	case 0x7B:
		return e.CPU8BitRegisterLoad(&e.AF.High, e.DE.Low)
	case 0x7C:
		return e.CPU8BitRegisterLoad(&e.AF.High, e.HL.High)
	case 0x7D:
		return e.CPU8BitRegisterLoad(&e.AF.High, e.HL.Low)
	case 0x7E:
		return e.CPU8BitRegisterMemoryAddressLoad(&e.AF.High, e.HL.Value())
	case 0x40:
		return e.CPU8BitRegisterLoad(&e.BC.High, e.BC.High)
	case 0x41:
		return e.CPU8BitRegisterLoad(&e.BC.High, e.BC.Low)
	case 0x42:
		return e.CPU8BitRegisterLoad(&e.BC.High, e.DE.High)
	case 0x43:
		return e.CPU8BitRegisterLoad(&e.BC.High, e.DE.Low)
	case 0x44:
		return e.CPU8BitRegisterLoad(&e.BC.High, e.HL.High)
	case 0x45:
		return e.CPU8BitRegisterLoad(&e.BC.High, e.HL.Low)
	case 0x46:
		return e.CPU8BitRegisterMemoryAddressLoad(&e.BC.High, e.HL.Value())
	case 0x48:
		return e.CPU8BitRegisterLoad(&e.BC.Low, e.BC.High)
	case 0x49:
		return e.CPU8BitRegisterLoad(&e.BC.Low, e.BC.Low)
	case 0x4A:
		return e.CPU8BitRegisterLoad(&e.BC.Low, e.DE.High)
	case 0x4B:
		return e.CPU8BitRegisterLoad(&e.BC.Low, e.DE.Low)
	case 0x4C:
		return e.CPU8BitRegisterLoad(&e.BC.Low, e.HL.High)
	case 0x4D:
		return e.CPU8BitRegisterLoad(&e.BC.Low, e.HL.Low)
	case 0x4E:
		return e.CPU8BitRegisterMemoryAddressLoad(&e.BC.Low, e.HL.Value())
	case 0x50:
		return e.CPU8BitRegisterLoad(&e.DE.High, e.BC.High)
	case 0x51:
		return e.CPU8BitRegisterLoad(&e.DE.High, e.BC.Low)
	case 0x52:
		return e.CPU8BitRegisterLoad(&e.DE.High, e.DE.High)
	case 0x53:
		return e.CPU8BitRegisterLoad(&e.DE.High, e.DE.Low)
	case 0x54:
		return e.CPU8BitRegisterLoad(&e.DE.High, e.HL.High)
	case 0x55:
		return e.CPU8BitRegisterLoad(&e.DE.High, e.HL.Low)
	case 0x56:
		return e.CPU8BitRegisterMemoryAddressLoad(&e.DE.High, e.HL.Value())
	case 0x58:
		return e.CPU8BitRegisterLoad(&e.DE.Low, e.BC.High)
	case 0x59:
		return e.CPU8BitRegisterLoad(&e.DE.Low, e.BC.Low)
	case 0x5A:
		return e.CPU8BitRegisterLoad(&e.DE.Low, e.DE.High)
	case 0x5B:
		return e.CPU8BitRegisterLoad(&e.DE.Low, e.DE.Low)
	case 0x5C:
		return e.CPU8BitRegisterLoad(&e.DE.Low, e.HL.High)
	case 0x5D:
		return e.CPU8BitRegisterLoad(&e.DE.Low, e.HL.Low)
	case 0x5E:
		return e.CPU8BitRegisterMemoryAddressLoad(&e.DE.Low, e.HL.Value())
	case 0x60:
		return e.CPU8BitRegisterLoad(&e.HL.High, e.BC.High)
	case 0x61:
		return e.CPU8BitRegisterLoad(&e.HL.High, e.BC.Low)
	case 0x62:
		return e.CPU8BitRegisterLoad(&e.HL.High, e.DE.High)
	case 0x63:
		return e.CPU8BitRegisterLoad(&e.HL.High, e.DE.Low)
	case 0x64:
		return e.CPU8BitRegisterLoad(&e.HL.High, e.HL.High)
	case 0x65:
		return e.CPU8BitRegisterLoad(&e.HL.High, e.HL.Low)
	case 0x66:
		return e.CPU8BitRegisterMemoryAddressLoad(&e.HL.High, e.HL.Value())
	case 0x68:
		return e.CPU8BitRegisterLoad(&e.HL.Low, e.BC.High)
	case 0x69:
		return e.CPU8BitRegisterLoad(&e.HL.Low, e.BC.Low)
	case 0x6A:
		return e.CPU8BitRegisterLoad(&e.HL.Low, e.DE.High)
	case 0x6B:
		return e.CPU8BitRegisterLoad(&e.HL.Low, e.DE.Low)
	case 0x6C:
		return e.CPU8BitRegisterLoad(&e.HL.Low, e.HL.High)
	case 0x6D:
		return e.CPU8BitRegisterLoad(&e.HL.Low, e.HL.Low)
	case 0x6E:
		return e.CPU8BitRegisterMemoryAddressLoad(&e.HL.Low, e.HL.Value())
	case 0x70:
		return e.CPU8BitRegisterMemoryWrite(e.HL.Value(), e.BC.High)
	case 0x71:
		return e.CPU8BitRegisterMemoryWrite(e.HL.Value(), e.BC.Low)
	case 0x72:
		return e.CPU8BitRegisterMemoryWrite(e.HL.Value(), e.DE.High)
	case 0x73:
		return e.CPU8BitRegisterMemoryWrite(e.HL.Value(), e.DE.Low)
	case 0x74:
		return e.CPU8BitRegisterMemoryWrite(e.HL.Value(), e.HL.High)
	case 0x75:
		return e.CPU8BitRegisterMemoryWrite(e.HL.Value(), e.HL.Low)
	case 0x36:
		value := e.ReadMemory8Bit(e.ProgramCounter.Value())
		e.ProgramCounter.Increment()
		e.WriteMemory(e.HL.Value(), value)
		return 12
	// LD A,n
	case 0x0A:
		return e.CPU8BitRegisterMemoryAddressLoad(&e.AF.High, e.BC.Value())
	case 0x1A:
		return e.CPU8BitRegisterMemoryAddressLoad(&e.AF.High, e.DE.Value())
	case 0xFA:
		address := e.ReadMemory16Bit(e.ProgramCounter.Value())
		e.ProgramCounter.Increment()
		e.ProgramCounter.Increment()
		value := e.ReadMemory8Bit(address)
		e.AF.High.SetValue(value)
		return 16
	case 0x3E:
		value := e.ReadMemory8Bit(e.ProgramCounter.Value())
		e.ProgramCounter.Increment()
		e.AF.High.SetValue(value)
		return 8
	// LD n,A
	case 0x47:
		return e.CPU8BitRegisterLoad(&e.BC.High, e.AF.High)
	case 0x4F:
		return e.CPU8BitRegisterLoad(&e.BC.Low, e.AF.High)
	case 0x57:
		return e.CPU8BitRegisterLoad(&e.DE.High, e.AF.High)
	case 0x5F:
		return e.CPU8BitRegisterLoad(&e.DE.Low, e.AF.High)
	case 0x67:
		return e.CPU8BitRegisterLoad(&e.HL.High, e.AF.High)
	case 0x6F:
		return e.CPU8BitRegisterLoad(&e.HL.Low, e.AF.High)
	case 0x02:
		return e.CPU8BitRegisterMemoryWrite(e.BC.Value(), e.AF.High)
	case 0x12:
		return e.CPU8BitRegisterMemoryWrite(e.DE.Value(), e.AF.High)
	case 0x77:
		return e.CPU8BitRegisterMemoryWrite(e.HL.Value(), e.AF.High)
	case 0xEA:
		address := e.ReadMemory16Bit(e.ProgramCounter.Value())
		e.ProgramCounter.Increment()
		e.ProgramCounter.Increment()
		e.WriteMemory(address, e.AF.High.Value())
		return 16
	// LD A,(C)
	case 0xF2:
		return e.CPU8BitRegisterMemoryAddressLoad(&e.AF.High, (0xFF00 + uint16(e.BC.Low.Value())))
	// LD (C),A
	case 0xE2:
		return e.CPU8BitRegisterMemoryWrite((0xFF00 + uint16(e.BC.Low.Value())), e.AF.High)
	// LD A,(HLD)
	// LD A,(HL-)
	// LDD A,(HL)
	case 0x3A:
		cycles := e.CPU8BitRegisterMemoryAddressLoad(&e.AF.High, e.HL.Value())
		e.HL.Decrement()
		return cycles
	// LD (HLD),A
	// LD (HL-),A
	// LDD (HL),A
	case 0x32:
		cycles := e.CPU8BitRegisterMemoryWrite(e.HL.Value(), e.AF.High)
		e.HL.Decrement()
		return cycles
	// LD A,(HLI)
	// LD A,(HL+)
	// LDI A,(HL)
	case 0x2A:
		cycles := e.CPU8BitRegisterMemoryAddressLoad(&e.AF.High, e.HL.Value())
		e.HL.Increment()
		return cycles
	// LD (HLI),A
	// LD (HL+),A
	// LDI (HL),A
	case 0x22:
		cycles := e.CPU8BitRegisterMemoryWrite(e.HL.Value(), e.AF.High)
		e.HL.Increment()
		return cycles
	// LDH (n),A
	case 0xE0:
		value := e.ReadMemory8Bit(e.ProgramCounter.Value())
		e.ProgramCounter.Increment()
		address := 0xFF00 + uint16(value)
		e.WriteMemory(address, e.AF.High.Value())
		return 12
	// LDH A,(n)
	case 0xF0:
		addressValue := e.ReadMemory8Bit(e.ProgramCounter.Value())
		e.ProgramCounter.Increment()
		address := 0xFF00 + uint16(addressValue)
		value := e.ReadMemory8Bit(address)
		e.AF.High.SetValue(value)
		return 12
	// 16-Bit Loads
	// LD n,nn
	case 0x01:
		return e.CPU16BitRegistryMemoryLoad(&e.BC)
	case 0x11:
		return e.CPU16BitRegistryMemoryLoad(&e.DE)
	case 0x21:
		return e.CPU16BitRegistryMemoryLoad(&e.HL)
	case 0x31:
		return e.CPU16BitRegistryMemoryLoad(&e.StackPointer)
	// LD SP,HL
	case 0xF9:
		e.StackPointer.SetValue(e.HL.Value())
		return 8
	// LD HL,SP+n
	// LDHL SP,n
	case 0xF8:
		signedValue := int8(e.ReadMemory8Bit(e.ProgramCounter.Value()))
		e.ProgramCounter.Increment()
		e.ClearFlagZ()
		e.ClearFlagN()

		value := uint32(e.StackPointer.Value()) + uint32(signedValue)

		if value > 0xFFFF {
			e.SetFlagC()
		} else {
			e.ClearFlagC()
		}

		if (uint32(e.StackPointer.Value())&0xF + value&0xF) > 0xF {
			e.SetFlagH()
		} else {
			e.ClearFlagH()
		}

		e.HL.SetValue(uint16(0x00FFFF & value))

		return 12
	// LD (nn),SP
	case 0x08:
		value := e.ReadMemory16Bit(e.ProgramCounter.Value())
		e.ProgramCounter.Increment()
		e.ProgramCounter.Increment()
		e.WriteMemory(value, e.StackPointer.Low.Value())
		value++
		e.WriteMemory(value, e.StackPointer.High.Value())
		return 20
	// PUSH nn
	case 0xF5:
		e.PushIntoStack(e.AF.Value())
		return 16
	case 0xC5:
		e.PushIntoStack(e.BC.Value())
		return 16
	case 0xD5:
		e.PushIntoStack(e.DE.Value())
		return 16
	case 0xE5:
		e.PushIntoStack(e.HL.Value())
		return 16
	// POP nn
	case 0xF1:
		e.AF.SetValue(e.PopFromStack())
		return 12
	case 0xC1:
		e.BC.SetValue(e.PopFromStack())
		return 12
	case 0xD1:
		e.DE.SetValue(e.PopFromStack())
		return 12
	case 0xE1:
		e.HL.SetValue(e.PopFromStack())
		return 12
	// 8-Bit ALU
	// ADD A,n
	case 0x87:
		return e.CPU8BitAdd(&e.AF.High, e.AF.High.Value(), false)
	case 0x80:
		return e.CPU8BitAdd(&e.AF.High, e.BC.High.Value(), false)
	case 0x81:
		return e.CPU8BitAdd(&e.AF.High, e.BC.Low.Value(), false)
	case 0x82:
		return e.CPU8BitAdd(&e.AF.High, e.DE.High.Value(), false)
	case 0x83:
		return e.CPU8BitAdd(&e.AF.High, e.DE.Low.Value(), false)
	case 0x84:
		return e.CPU8BitAdd(&e.AF.High, e.HL.High.Value(), false)
	case 0x85:
		return e.CPU8BitAdd(&e.AF.High, e.HL.Low.Value(), false)
	case 0x86:
		cycles := e.CPU8BitAdd(&e.AF.High, e.ReadMemory8Bit(e.HL.Value()), false)
		cycles += 4 // 8
		return cycles
	case 0xC6:
		value := e.ReadMemory8Bit(e.ProgramCounter.Value())
		e.ProgramCounter.Increment()
		cycles := e.CPU8BitAdd(&e.AF.High, value, false)
		cycles += 4 // 8
		return cycles
	// ADC A,n
	case 0x8F:
		return e.CPU8BitAdd(&e.AF.High, e.AF.High.Value(), true)
	case 0x88:
		return e.CPU8BitAdd(&e.AF.High, e.BC.High.Value(), true)
	case 0x89:
		return e.CPU8BitAdd(&e.AF.High, e.BC.Low.Value(), true)
	case 0x8A:
		return e.CPU8BitAdd(&e.AF.High, e.DE.High.Value(), true)
	case 0x8B:
		return e.CPU8BitAdd(&e.AF.High, e.DE.Low.Value(), true)
	case 0x8C:
		return e.CPU8BitAdd(&e.AF.High, e.HL.High.Value(), true)
	case 0x8D:
		return e.CPU8BitAdd(&e.AF.High, e.HL.Low.Value(), true)
	case 0x8E:
		cycles := e.CPU8BitAdd(&e.AF.High, e.ReadMemory8Bit(e.HL.Value()), true)
		cycles += 4 // 8
		return cycles
	case 0xCE:
		value := e.ReadMemory8Bit(e.ProgramCounter.Value())
		e.ProgramCounter.Increment()
		cycles := e.CPU8BitAdd(&e.AF.High, value, true)
		cycles += 4 // 8
		return cycles
	// SUB A,n
	case 0x97:
		return e.CPU8BitSubtract(&e.AF.High, e.AF.High.Value(), false)
	case 0x90:
		return e.CPU8BitSubtract(&e.AF.High, e.BC.High.Value(), false)
	case 0x91:
		return e.CPU8BitSubtract(&e.AF.High, e.BC.Low.Value(), false)
	case 0x92:
		return e.CPU8BitSubtract(&e.AF.High, e.DE.High.Value(), false)
	case 0x93:
		return e.CPU8BitSubtract(&e.AF.High, e.DE.Low.Value(), false)
	case 0x94:
		return e.CPU8BitSubtract(&e.AF.High, e.HL.High.Value(), false)
	case 0x95:
		return e.CPU8BitSubtract(&e.AF.High, e.HL.Low.Value(), false)
	case 0x96:
		cycles := e.CPU8BitSubtract(&e.AF.High, e.ReadMemory8Bit(e.HL.Value()), false)
		cycles += 4 // 8
		return cycles
	case 0xD6:
		value := e.ReadMemory8Bit(e.ProgramCounter.Value())
		e.ProgramCounter.Increment()
		cycles := e.CPU8BitSubtract(&e.AF.High, value, false)
		cycles += 4 // 8
		return cycles
	// SBC A,n
	case 0x9F:
		return e.CPU8BitSubtract(&e.AF.High, e.AF.High.Value(), true)
	case 0x98:
		return e.CPU8BitSubtract(&e.AF.High, e.BC.High.Value(), true)
	case 0x99:
		return e.CPU8BitSubtract(&e.AF.High, e.BC.Low.Value(), true)
	case 0x9A:
		return e.CPU8BitSubtract(&e.AF.High, e.DE.High.Value(), true)
	case 0x9B:
		return e.CPU8BitSubtract(&e.AF.High, e.DE.Low.Value(), true)
	case 0x9C:
		return e.CPU8BitSubtract(&e.AF.High, e.HL.High.Value(), true)
	case 0x9D:
		return e.CPU8BitSubtract(&e.AF.High, e.HL.Low.Value(), true)
	case 0x9E:
		cycles := e.CPU8BitSubtract(&e.AF.High, e.ReadMemory8Bit(e.HL.Value()), true)
		cycles += 4 // 8
		return cycles
	case 0xDE:
		value := e.ReadMemory8Bit(e.ProgramCounter.Value())
		e.ProgramCounter.Increment()
		cycles := e.CPU8BitSubtract(&e.AF.High, value, true)
		cycles += 4 // 8
		return cycles
	// AND n
	case 0xA7:
		return e.CPU8BitAnd(e.AF.High.Value())
	case 0xA0:
		return e.CPU8BitAnd(e.BC.High.Value())
	case 0xA1:
		return e.CPU8BitAnd(e.BC.Low.Value())
	case 0xA2:
		return e.CPU8BitAnd(e.DE.High.Value())
	case 0xA3:
		return e.CPU8BitAnd(e.DE.Low.Value())
	case 0xA4:
		return e.CPU8BitAnd(e.HL.High.Value())
	case 0xA5:
		return e.CPU8BitAnd(e.HL.Low.Value())
	case 0xA6:
		value := e.ReadMemory8Bit(e.HL.Value())
		cycles := e.CPU8BitAnd(value)
		cycles += 4
		return cycles
	case 0xE6:
		value := e.ReadMemory8Bit(e.ProgramCounter.Value())
		e.ProgramCounter.Increment()
		cycles := e.CPU8BitAnd(value)
		cycles += 4
		return cycles
	// OR n
	case 0xB7:
		return e.CPU8BitOr(e.AF.High.Value())
	case 0xB0:
		return e.CPU8BitOr(e.BC.High.Value())
	case 0xB1:
		return e.CPU8BitOr(e.BC.Low.Value())
	case 0xB2:
		return e.CPU8BitOr(e.DE.High.Value())
	case 0xB3:
		return e.CPU8BitOr(e.DE.Low.Value())
	case 0xB4:
		return e.CPU8BitOr(e.HL.High.Value())
	case 0xB5:
		return e.CPU8BitOr(e.HL.Low.Value())
	case 0xB6:
		value := e.ReadMemory8Bit(e.HL.Value())
		cycles := e.CPU8BitOr(value)
		cycles += 4
		return cycles
	case 0xF6:
		value := e.ReadMemory8Bit(e.ProgramCounter.Value())
		e.ProgramCounter.Increment()
		cycles := e.CPU8BitOr(value)
		cycles += 4
		return cycles
	// XOR n
	case 0xAF:
		return e.CPU8BitXor(e.AF.High.Value())
	case 0xA8:
		return e.CPU8BitXor(e.BC.High.Value())
	case 0xA9:
		return e.CPU8BitXor(e.BC.Low.Value())
	case 0xAA:
		return e.CPU8BitXor(e.DE.High.Value())
	case 0xAB:
		return e.CPU8BitXor(e.DE.Low.Value())
	case 0xAC:
		return e.CPU8BitXor(e.HL.High.Value())
	case 0xAD:
		return e.CPU8BitXor(e.HL.Low.Value())
	case 0xAE:
		value := e.ReadMemory8Bit(e.HL.Value())
		cycles := e.CPU8BitXor(value)
		cycles += 4
		return cycles
	case 0xEE:
		value := e.ReadMemory8Bit(e.ProgramCounter.Value())
		e.ProgramCounter.Increment()
		cycles := e.CPU8BitXor(value)
		cycles += 4
		return cycles
	// CP n
	case 0xBF:
		return e.CPU8BitCompare(e.AF.High.Value())
	case 0xB8:
		return e.CPU8BitCompare(e.BC.High.Value())
	case 0xB9:
		return e.CPU8BitCompare(e.BC.Low.Value())
	case 0xBA:
		return e.CPU8BitCompare(e.DE.High.Value())
	case 0xBB:
		return e.CPU8BitCompare(e.DE.Low.Value())
	case 0xBC:
		return e.CPU8BitCompare(e.HL.High.Value())
	case 0xBD:
		return e.CPU8BitCompare(e.HL.Low.Value())
	case 0xBE:
		value := e.ReadMemory8Bit(e.HL.Value())
		cycles := e.CPU8BitCompare(value)
		cycles += 4
		return cycles
	case 0xFE:
		value := e.ReadMemory8Bit(e.ProgramCounter.Value())
		e.ProgramCounter.Increment()
		cycles := e.CPU8BitCompare(value)
		cycles += 4
		return cycles
	}

	return 0
}
