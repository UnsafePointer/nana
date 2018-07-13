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
	}

	return 0
}
