package emulator

func (e *Emulator) ExecuteOpCode(opcode uint8) int {
	switch opcode {
	// no-op
	case 0x00:
		return 4
	// 8-Bit Loads
	// LD nn,n
	case 0x06:
		return e.CPU8BitRegisterMemoryLoad(&e.BC.high)
	case 0x0E:
		return e.CPU8BitRegisterMemoryLoad(&e.BC.low)
	case 0x16:
		return e.CPU8BitRegisterMemoryLoad(&e.DE.high)
	case 0x1E:
		return e.CPU8BitRegisterMemoryLoad(&e.DE.low)
	case 0x26:
		return e.CPU8BitRegisterMemoryLoad(&e.HL.high)
	case 0x2E:
		return e.CPU8BitRegisterMemoryLoad(&e.HL.low)
	// LD r1,r2
	case 0x7F:
		return e.CPU8BitRegisterLoad(&e.AF.high, e.AF.high)
	case 0x78:
		return e.CPU8BitRegisterLoad(&e.AF.high, e.BC.high)
	case 0x79:
		return e.CPU8BitRegisterLoad(&e.AF.high, e.BC.low)
	case 0x7A:
		return e.CPU8BitRegisterLoad(&e.AF.high, e.DE.high)
	case 0x7B:
		return e.CPU8BitRegisterLoad(&e.AF.high, e.DE.low)
	case 0x7C:
		return e.CPU8BitRegisterLoad(&e.AF.high, e.HL.high)
	case 0x7D:
		return e.CPU8BitRegisterLoad(&e.AF.high, e.HL.low)
	case 0x7E:
		return e.CPU8BitRegisterMemoryAddressLoad(&e.AF.high, e.HL.Value())
	case 0x40:
		return e.CPU8BitRegisterLoad(&e.BC.high, e.BC.high)
	case 0x41:
		return e.CPU8BitRegisterLoad(&e.BC.high, e.BC.low)
	case 0x42:
		return e.CPU8BitRegisterLoad(&e.BC.high, e.DE.high)
	case 0x43:
		return e.CPU8BitRegisterLoad(&e.BC.high, e.DE.low)
	case 0x44:
		return e.CPU8BitRegisterLoad(&e.BC.high, e.HL.high)
	case 0x45:
		return e.CPU8BitRegisterLoad(&e.BC.high, e.HL.low)
	case 0x46:
		return e.CPU8BitRegisterMemoryAddressLoad(&e.BC.high, e.HL.Value())
	case 0x48:
		return e.CPU8BitRegisterLoad(&e.BC.low, e.BC.high)
	case 0x49:
		return e.CPU8BitRegisterLoad(&e.BC.low, e.BC.low)
	case 0x4A:
		return e.CPU8BitRegisterLoad(&e.BC.low, e.DE.high)
	case 0x4B:
		return e.CPU8BitRegisterLoad(&e.BC.low, e.DE.low)
	case 0x4C:
		return e.CPU8BitRegisterLoad(&e.BC.low, e.HL.high)
	case 0x4D:
		return e.CPU8BitRegisterLoad(&e.BC.low, e.HL.low)
	case 0x4E:
		return e.CPU8BitRegisterMemoryAddressLoad(&e.BC.low, e.HL.Value())
	case 0x50:
		return e.CPU8BitRegisterLoad(&e.DE.high, e.BC.high)
	case 0x51:
		return e.CPU8BitRegisterLoad(&e.DE.high, e.BC.low)
	case 0x52:
		return e.CPU8BitRegisterLoad(&e.DE.high, e.DE.high)
	case 0x53:
		return e.CPU8BitRegisterLoad(&e.DE.high, e.DE.low)
	case 0x54:
		return e.CPU8BitRegisterLoad(&e.DE.high, e.HL.high)
	case 0x55:
		return e.CPU8BitRegisterLoad(&e.DE.high, e.HL.low)
	case 0x56:
		return e.CPU8BitRegisterMemoryAddressLoad(&e.DE.high, e.HL.Value())
	case 0x58:
		return e.CPU8BitRegisterLoad(&e.DE.low, e.BC.high)
	case 0x59:
		return e.CPU8BitRegisterLoad(&e.DE.low, e.BC.low)
	case 0x5A:
		return e.CPU8BitRegisterLoad(&e.DE.low, e.DE.high)
	case 0x5B:
		return e.CPU8BitRegisterLoad(&e.DE.low, e.DE.low)
	case 0x5C:
		return e.CPU8BitRegisterLoad(&e.DE.low, e.HL.high)
	case 0x5D:
		return e.CPU8BitRegisterLoad(&e.DE.low, e.HL.low)
	case 0x5E:
		return e.CPU8BitRegisterMemoryAddressLoad(&e.DE.low, e.HL.Value())
	case 0x60:
		return e.CPU8BitRegisterLoad(&e.HL.high, e.BC.high)
	case 0x61:
		return e.CPU8BitRegisterLoad(&e.HL.high, e.BC.low)
	case 0x62:
		return e.CPU8BitRegisterLoad(&e.HL.high, e.DE.high)
	case 0x63:
		return e.CPU8BitRegisterLoad(&e.HL.high, e.DE.low)
	case 0x64:
		return e.CPU8BitRegisterLoad(&e.HL.high, e.HL.high)
	case 0x65:
		return e.CPU8BitRegisterLoad(&e.HL.high, e.HL.low)
	case 0x66:
		return e.CPU8BitRegisterMemoryAddressLoad(&e.HL.high, e.HL.Value())
	case 0x68:
		return e.CPU8BitRegisterLoad(&e.HL.low, e.BC.high)
	case 0x69:
		return e.CPU8BitRegisterLoad(&e.HL.low, e.BC.low)
	case 0x6A:
		return e.CPU8BitRegisterLoad(&e.HL.low, e.DE.high)
	case 0x6B:
		return e.CPU8BitRegisterLoad(&e.HL.low, e.DE.low)
	case 0x6C:
		return e.CPU8BitRegisterLoad(&e.HL.low, e.HL.high)
	case 0x6D:
		return e.CPU8BitRegisterLoad(&e.HL.low, e.HL.low)
	case 0x6E:
		return e.CPU8BitRegisterMemoryAddressLoad(&e.HL.low, e.HL.Value())
	case 0x70:
		return e.CPU8BitRegisterMemoryWrite(e.HL.Value(), e.BC.high)
	case 0x71:
		return e.CPU8BitRegisterMemoryWrite(e.HL.Value(), e.BC.low)
	case 0x72:
		return e.CPU8BitRegisterMemoryWrite(e.HL.Value(), e.DE.high)
	case 0x73:
		return e.CPU8BitRegisterMemoryWrite(e.HL.Value(), e.DE.low)
	case 0x74:
		return e.CPU8BitRegisterMemoryWrite(e.HL.Value(), e.HL.high)
	case 0x75:
		return e.CPU8BitRegisterMemoryWrite(e.HL.Value(), e.HL.low)
	case 0x36:
		value := e.ReadMemory8Bit(e.ProgramCounter.Value())
		e.ProgramCounter.Increment()
		e.WriteMemory(e.HL.Value(), value)
		return 12
	// LD A,n
	case 0x0A:
		return e.CPU8BitRegisterMemoryAddressLoad(&e.AF.high, e.BC.Value())
	case 0x1A:
		return e.CPU8BitRegisterMemoryAddressLoad(&e.AF.high, e.DE.Value())
	case 0xFA:
		address := e.ReadMemory16Bit(e.ProgramCounter.Value())
		e.ProgramCounter.Increment()
		e.ProgramCounter.Increment()
		value := e.ReadMemory8Bit(address)
		e.AF.high.SetValue(value)
		return 16
	case 0x3E:
		value := e.ReadMemory8Bit(e.ProgramCounter.Value())
		e.ProgramCounter.Increment()
		e.AF.high.SetValue(value)
		return 8
	// LD n,A
	case 0x47:
		return e.CPU8BitRegisterLoad(&e.BC.high, e.AF.high)
	case 0x4F:
		return e.CPU8BitRegisterLoad(&e.BC.low, e.AF.high)
	case 0x57:
		return e.CPU8BitRegisterLoad(&e.DE.high, e.AF.high)
	case 0x5F:
		return e.CPU8BitRegisterLoad(&e.DE.low, e.AF.high)
	case 0x67:
		return e.CPU8BitRegisterLoad(&e.HL.high, e.AF.high)
	case 0x6F:
		return e.CPU8BitRegisterLoad(&e.HL.low, e.AF.high)
	case 0x02:
		return e.CPU8BitRegisterMemoryWrite(e.BC.Value(), e.AF.high)
	case 0x12:
		return e.CPU8BitRegisterMemoryWrite(e.DE.Value(), e.AF.high)
	case 0x77:
		return e.CPU8BitRegisterMemoryWrite(e.HL.Value(), e.AF.high)
	case 0xEA:
		address := e.ReadMemory16Bit(e.ProgramCounter.Value())
		e.ProgramCounter.Increment()
		e.ProgramCounter.Increment()
		e.WriteMemory(address, e.AF.high.Value())
		return 16
	// LD A,(C)
	case 0xF2:
		return e.CPU8BitRegisterMemoryAddressLoad(&e.AF.high, (0xFF00 + uint16(e.BC.low.Value())))
	// LD (C),A
	case 0xE2:
		return e.CPU8BitRegisterMemoryWrite((0xFF00 + uint16(e.BC.low.Value())), e.AF.high)
	// LD A,(HLD)
	// LD A,(HL-)
	// LDD A,(HL)
	case 0x3A:
		cycles := e.CPU8BitRegisterMemoryAddressLoad(&e.AF.high, e.HL.Value())
		e.HL.Decrement()
		return cycles
	// LD (HLD),A
	// LD (HL-),A
	// LDD (HL),A
	case 0x32:
		cycles := e.CPU8BitRegisterMemoryWrite(e.HL.Value(), e.AF.high)
		e.HL.Decrement()
		return cycles
	// LD A,(HLI)
	// LD A,(HL+)
	// LDI A,(HL)
	case 0x2A:
		cycles := e.CPU8BitRegisterMemoryAddressLoad(&e.AF.high, e.HL.Value())
		e.HL.Increment()
		return cycles
	// LD (HLI),A
	// LD (HL+),A
	// LDI (HL),A
	case 0x22:
		cycles := e.CPU8BitRegisterMemoryWrite(e.HL.Value(), e.AF.high)
		e.HL.Increment()
		return cycles
	// LDH (n),A
	case 0xE0:
		value := e.ReadMemory8Bit(e.ProgramCounter.Value())
		e.ProgramCounter.Increment()
		address := 0xFF00 + uint16(value)
		e.WriteMemory(address, e.AF.high.Value())
		return 12
	// LDH A,(n)
	case 0xF0:
		addressValue := e.ReadMemory8Bit(e.ProgramCounter.Value())
		e.ProgramCounter.Increment()
		address := 0xFF00 + uint16(addressValue)
		value := e.ReadMemory8Bit(address)
		e.AF.high.SetValue(value)
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
		e.WriteMemory(value, e.StackPointer.Low())
		value++
		e.WriteMemory(value, e.StackPointer.High())
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
	}

	return 0
}
