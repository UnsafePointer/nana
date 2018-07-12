package emulator

func (e *Emulator) ExecuteOpCode(opcode uint8) int {
	switch opcode {
	// no-op
	case 0x00:
		return 4
	// 8-Bit Loads
	// LD nn,n
	case 0x06:
		return e.CPU8BitLoad(&e.BC.high)
	case 0x0E:
		return e.CPU8BitLoad(&e.BC.low)
	case 0x16:
		return e.CPU8BitLoad(&e.DE.high)
	case 0x1E:
		return e.CPU8BitLoad(&e.DE.low)
	case 0x26:
		return e.CPU8BitLoad(&e.HL.high)
	case 0x2E:
		return e.CPU8BitLoad(&e.HL.low)
	// LD r1,r2
	case 0x7F:
		return e.CPURegLoad(&e.AF.high, e.AF.high)
	case 0x78:
		return e.CPURegLoad(&e.AF.high, e.BC.high)
	case 0x79:
		return e.CPURegLoad(&e.AF.high, e.BC.low)
	case 0x7A:
		return e.CPURegLoad(&e.AF.high, e.DE.high)
	case 0x7B:
		return e.CPURegLoad(&e.AF.high, e.DE.low)
	case 0x7C:
		return e.CPURegLoad(&e.AF.high, e.HL.high)
	case 0x7D:
		return e.CPURegLoad(&e.AF.high, e.HL.low)
	case 0x7E:
		return 0
	case 0x40:
		return e.CPURegLoad(&e.BC.high, e.BC.high)
	case 0x41:
		return e.CPURegLoad(&e.BC.high, e.BC.low)
	case 0x42:
		return e.CPURegLoad(&e.BC.high, e.DE.high)
	case 0x43:
		return e.CPURegLoad(&e.BC.high, e.DE.low)
	case 0x44:
		return e.CPURegLoad(&e.BC.high, e.HL.high)
	case 0x45:
		return e.CPURegLoad(&e.BC.high, e.HL.low)
	case 0x46:
		return 0
	case 0x48:
		return e.CPURegLoad(&e.BC.low, e.BC.high)
	case 0x49:
		return e.CPURegLoad(&e.BC.low, e.BC.low)
	case 0x4A:
		return e.CPURegLoad(&e.BC.low, e.DE.high)
	case 0x4B:
		return e.CPURegLoad(&e.BC.low, e.DE.low)
	case 0x4C:
		return e.CPURegLoad(&e.BC.low, e.HL.high)
	case 0x4D:
		return e.CPURegLoad(&e.BC.low, e.HL.low)
	case 0x4E:
		return 0
	case 0x50:
		return e.CPURegLoad(&e.DE.high, e.BC.high)
	case 0x51:
		return e.CPURegLoad(&e.DE.high, e.BC.low)
	case 0x52:
		return e.CPURegLoad(&e.DE.high, e.DE.high)
	case 0x53:
		return e.CPURegLoad(&e.DE.high, e.DE.low)
	case 0x54:
		return e.CPURegLoad(&e.DE.high, e.HL.high)
	case 0x55:
		return e.CPURegLoad(&e.DE.high, e.HL.low)
	case 0x56:
		return 0
	case 0x58:
		return e.CPURegLoad(&e.DE.low, e.BC.high)
	case 0x59:
		return e.CPURegLoad(&e.DE.low, e.BC.low)
	case 0x5A:
		return e.CPURegLoad(&e.DE.low, e.DE.high)
	case 0x5B:
		return e.CPURegLoad(&e.DE.low, e.DE.low)
	case 0x5C:
		return e.CPURegLoad(&e.DE.low, e.HL.high)
	case 0x5D:
		return e.CPURegLoad(&e.DE.low, e.HL.low)
	case 0x5E:
		return 0
	case 0x60:
		return e.CPURegLoad(&e.HL.high, e.BC.high)
	case 0x61:
		return e.CPURegLoad(&e.HL.high, e.BC.low)
	case 0x62:
		return e.CPURegLoad(&e.HL.high, e.DE.high)
	case 0x63:
		return e.CPURegLoad(&e.HL.high, e.DE.low)
	case 0x64:
		return e.CPURegLoad(&e.HL.high, e.HL.high)
	case 0x65:
		return e.CPURegLoad(&e.HL.high, e.HL.low)
	case 0x66:
		return 0
	case 0x68:
		return e.CPURegLoad(&e.HL.low, e.BC.high)
	case 0x69:
		return e.CPURegLoad(&e.HL.low, e.BC.low)
	case 0x6A:
		return e.CPURegLoad(&e.HL.low, e.DE.high)
	case 0x6B:
		return e.CPURegLoad(&e.HL.low, e.DE.low)
	case 0x6C:
		return e.CPURegLoad(&e.HL.low, e.HL.high)
	case 0x6D:
		return e.CPURegLoad(&e.HL.low, e.HL.low)
	case 0x6E:
		return 0
	case 0x70:
		return 0
	case 0x71:
		return 0
	case 0x72:
		return 0
	case 0x73:
		return 0
	case 0x74:
		return 0
	case 0x75:
		return 0
	case 0x36:
		return 0
	}

	return 0
}
