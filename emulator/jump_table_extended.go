package emulator

func (e *Emulator) ExecuteExtendedOpCode(opcode uint8) int {
	switch opcode {
	// Miscellaneous
	// SWAP n
	case 0x37:
		return e.CPU8BitRegisterSwap(&e.AF.High)
	case 0x30:
		return e.CPU8BitRegisterSwap(&e.BC.High)
	case 0x31:
		return e.CPU8BitRegisterSwap(&e.BC.Low)
	case 0x32:
		return e.CPU8BitRegisterSwap(&e.DE.High)
	case 0x33:
		return e.CPU8BitRegisterSwap(&e.DE.Low)
	case 0x34:
		return e.CPU8BitRegisterSwap(&e.HL.High)
	case 0x35:
		return e.CPU8BitRegisterSwap(&e.HL.Low)
	case 0x36:
		return e.CPU8BitSwapMemoryAddress(e.HL.Value())
	// Rotates & Shifts
	// RLC n
	case 0x07:
		return e.CPU8BitRegisterRLC(&e.AF.High)
	case 0x00:
		return e.CPU8BitRegisterRLC(&e.BC.High)
	case 0x01:
		return e.CPU8BitRegisterRLC(&e.BC.Low)
	case 0x02:
		return e.CPU8BitRegisterRLC(&e.DE.High)
	case 0x03:
		return e.CPU8BitRegisterRLC(&e.DE.Low)
	case 0x04:
		return e.CPU8BitRegisterRLC(&e.HL.High)
	case 0x05:
		return e.CPU8BitRegisterRLC(&e.HL.Low)
	case 0x06:
		return e.CPU8BitRLCMemoryAddress(e.HL.Value())
	// RL n
	case 0x17:
		return e.CPU8BitRegisterRL(&e.AF.High)
	case 0x10:
		return e.CPU8BitRegisterRL(&e.BC.High)
	case 0x11:
		return e.CPU8BitRegisterRL(&e.BC.Low)
	case 0x12:
		return e.CPU8BitRegisterRL(&e.DE.High)
	case 0x13:
		return e.CPU8BitRegisterRL(&e.DE.Low)
	case 0x14:
		return e.CPU8BitRegisterRL(&e.HL.High)
	case 0x15:
		return e.CPU8BitRegisterRL(&e.HL.Low)
	case 0x16:
		return e.CPU8BitRLMemoryAddress(e.HL.Value())
	// RRC n
	case 0x0F:
		return e.CPU8BitRegisterRRC(&e.AF.High)
	case 0x08:
		return e.CPU8BitRegisterRRC(&e.BC.High)
	case 0x09:
		return e.CPU8BitRegisterRRC(&e.BC.Low)
	case 0x0A:
		return e.CPU8BitRegisterRRC(&e.DE.High)
	case 0x0B:
		return e.CPU8BitRegisterRRC(&e.DE.Low)
	case 0x0C:
		return e.CPU8BitRegisterRRC(&e.HL.High)
	case 0x0D:
		return e.CPU8BitRegisterRRC(&e.HL.Low)
	case 0x0E:
		return e.CPU8BitRRCMemoryAddress(e.HL.Value())
	// RR n
	case 0x1F:
		return e.CPU8BitRegisterRR(&e.AF.High)
	case 0x18:
		return e.CPU8BitRegisterRR(&e.BC.High)
	case 0x19:
		return e.CPU8BitRegisterRR(&e.BC.Low)
	case 0x1A:
		return e.CPU8BitRegisterRR(&e.DE.High)
	case 0x1B:
		return e.CPU8BitRegisterRR(&e.DE.Low)
	case 0x1C:
		return e.CPU8BitRegisterRR(&e.HL.High)
	case 0x1D:
		return e.CPU8BitRegisterRR(&e.HL.Low)
	case 0x1E:
		return e.CPU8BitRRMemoryAddress(e.HL.Value())
	// SLA n
	case 0x27:
		return e.CPU8BitRegisterSLA(&e.AF.High)
	case 0x20:
		return e.CPU8BitRegisterSLA(&e.BC.High)
	case 0x21:
		return e.CPU8BitRegisterSLA(&e.BC.Low)
	case 0x22:
		return e.CPU8BitRegisterSLA(&e.DE.High)
	case 0x23:
		return e.CPU8BitRegisterSLA(&e.DE.Low)
	case 0x24:
		return e.CPU8BitRegisterSLA(&e.HL.High)
	case 0x25:
		return e.CPU8BitRegisterSLA(&e.HL.Low)
	case 0x26:
		return e.CPU8BitSLAMemoryAddress(e.HL.Value())
	// SRA n
	case 0x2F:
		return e.CPU8BitRegisterSRA(&e.AF.High)
	case 0x28:
		return e.CPU8BitRegisterSRA(&e.BC.High)
	case 0x29:
		return e.CPU8BitRegisterSRA(&e.BC.Low)
	case 0x2A:
		return e.CPU8BitRegisterSRA(&e.DE.High)
	case 0x2B:
		return e.CPU8BitRegisterSRA(&e.DE.Low)
	case 0x2C:
		return e.CPU8BitRegisterSRA(&e.HL.High)
	case 0x2D:
		return e.CPU8BitRegisterSRA(&e.HL.Low)
	case 0x2E:
		return e.CPU8BitSRAMemoryAddress(e.HL.Value())
	// SRL n
	case 0x3F:
		return e.CPU8BitRegisterSRL(&e.AF.High)
	case 0x38:
		return e.CPU8BitRegisterSRL(&e.BC.High)
	case 0x39:
		return e.CPU8BitRegisterSRL(&e.BC.Low)
	case 0x3A:
		return e.CPU8BitRegisterSRL(&e.DE.High)
	case 0x3B:
		return e.CPU8BitRegisterSRL(&e.DE.Low)
	case 0x3C:
		return e.CPU8BitRegisterSRL(&e.HL.High)
	case 0x3D:
		return e.CPU8BitRegisterSRL(&e.HL.Low)
	case 0x3E:
		return e.CPU8BitSRLMemoryAddress(e.HL.Value())
	// BIT b,r
	// b=0
	case 0x47:
		return e.CPU8BitRegisterBit(&e.AF.High, 0)
	case 0x40:
		return e.CPU8BitRegisterBit(&e.BC.High, 0)
	case 0x41:
		return e.CPU8BitRegisterBit(&e.BC.Low, 0)
	case 0x42:
		return e.CPU8BitRegisterBit(&e.DE.High, 0)
	case 0x43:
		return e.CPU8BitRegisterBit(&e.DE.Low, 0)
	case 0x44:
		return e.CPU8BitRegisterBit(&e.HL.High, 0)
	case 0x45:
		return e.CPU8BitRegisterBit(&e.HL.Low, 0)
	case 0x46:
		return e.CPU8BitBitMemoryAddress(e.HL.Value(), 0)
	// b=1
	case 0x4F:
		return e.CPU8BitRegisterBit(&e.AF.High, 1)
	case 0x48:
		return e.CPU8BitRegisterBit(&e.BC.High, 1)
	case 0x49:
		return e.CPU8BitRegisterBit(&e.BC.Low, 1)
	case 0x4A:
		return e.CPU8BitRegisterBit(&e.DE.High, 1)
	case 0x4B:
		return e.CPU8BitRegisterBit(&e.DE.Low, 1)
	case 0x4C:
		return e.CPU8BitRegisterBit(&e.HL.High, 1)
	case 0x4D:
		return e.CPU8BitRegisterBit(&e.HL.Low, 1)
	case 0x4E:
		return e.CPU8BitBitMemoryAddress(e.HL.Value(), 1)
	// b=2
	case 0x57:
		return e.CPU8BitRegisterBit(&e.AF.High, 2)
	case 0x50:
		return e.CPU8BitRegisterBit(&e.BC.High, 2)
	case 0x51:
		return e.CPU8BitRegisterBit(&e.BC.Low, 2)
	case 0x52:
		return e.CPU8BitRegisterBit(&e.DE.High, 2)
	case 0x53:
		return e.CPU8BitRegisterBit(&e.DE.Low, 2)
	case 0x54:
		return e.CPU8BitRegisterBit(&e.HL.High, 2)
	case 0x55:
		return e.CPU8BitRegisterBit(&e.HL.Low, 2)
	case 0x56:
		return e.CPU8BitBitMemoryAddress(e.HL.Value(), 2)
	// b=3
	case 0x5F:
		return e.CPU8BitRegisterBit(&e.AF.High, 3)
	case 0x58:
		return e.CPU8BitRegisterBit(&e.BC.High, 3)
	case 0x59:
		return e.CPU8BitRegisterBit(&e.BC.Low, 3)
	case 0x5A:
		return e.CPU8BitRegisterBit(&e.DE.High, 3)
	case 0x5B:
		return e.CPU8BitRegisterBit(&e.DE.Low, 3)
	case 0x5C:
		return e.CPU8BitRegisterBit(&e.HL.High, 3)
	case 0x5D:
		return e.CPU8BitRegisterBit(&e.HL.Low, 3)
	case 0x5E:
		return e.CPU8BitBitMemoryAddress(e.HL.Value(), 3)
	// b=4
	case 0x67:
		return e.CPU8BitRegisterBit(&e.AF.High, 4)
	case 0x60:
		return e.CPU8BitRegisterBit(&e.BC.High, 4)
	case 0x61:
		return e.CPU8BitRegisterBit(&e.BC.Low, 4)
	case 0x62:
		return e.CPU8BitRegisterBit(&e.DE.High, 4)
	case 0x63:
		return e.CPU8BitRegisterBit(&e.DE.Low, 4)
	case 0x64:
		return e.CPU8BitRegisterBit(&e.HL.High, 4)
	case 0x65:
		return e.CPU8BitRegisterBit(&e.HL.Low, 4)
	case 0x66:
		return e.CPU8BitBitMemoryAddress(e.HL.Value(), 4)
	// b=5
	case 0x6F:
		return e.CPU8BitRegisterBit(&e.AF.High, 5)
	case 0x68:
		return e.CPU8BitRegisterBit(&e.BC.High, 5)
	case 0x69:
		return e.CPU8BitRegisterBit(&e.BC.Low, 5)
	case 0x6A:
		return e.CPU8BitRegisterBit(&e.DE.High, 5)
	case 0x6B:
		return e.CPU8BitRegisterBit(&e.DE.Low, 5)
	case 0x6C:
		return e.CPU8BitRegisterBit(&e.HL.High, 5)
	case 0x6D:
		return e.CPU8BitRegisterBit(&e.HL.Low, 5)
	case 0x6E:
		return e.CPU8BitBitMemoryAddress(e.HL.Value(), 5)
	// b=6
	case 0x77:
		return e.CPU8BitRegisterBit(&e.AF.High, 6)
	case 0x70:
		return e.CPU8BitRegisterBit(&e.BC.High, 6)
	case 0x71:
		return e.CPU8BitRegisterBit(&e.BC.Low, 6)
	case 0x72:
		return e.CPU8BitRegisterBit(&e.DE.High, 6)
	case 0x73:
		return e.CPU8BitRegisterBit(&e.DE.Low, 6)
	case 0x74:
		return e.CPU8BitRegisterBit(&e.HL.High, 6)
	case 0x75:
		return e.CPU8BitRegisterBit(&e.HL.Low, 6)
	case 0x76:
		return e.CPU8BitBitMemoryAddress(e.HL.Value(), 6)
	// b=7
	case 0x7F:
		return e.CPU8BitRegisterBit(&e.AF.High, 7)
	case 0x78:
		return e.CPU8BitRegisterBit(&e.BC.High, 7)
	case 0x79:
		return e.CPU8BitRegisterBit(&e.BC.Low, 7)
	case 0x7A:
		return e.CPU8BitRegisterBit(&e.DE.High, 7)
	case 0x7B:
		return e.CPU8BitRegisterBit(&e.DE.Low, 7)
	case 0x7C:
		return e.CPU8BitRegisterBit(&e.HL.High, 7)
	case 0x7D:
		return e.CPU8BitRegisterBit(&e.HL.Low, 7)
	case 0x7E:
		return e.CPU8BitBitMemoryAddress(e.HL.Value(), 7)
	// SET b,r
	// b=0
	case 0xC7:
		return e.CPU8BitRegisterSet(&e.AF.High, 0)
	case 0xC0:
		return e.CPU8BitRegisterSet(&e.BC.High, 0)
	case 0xC1:
		return e.CPU8BitRegisterSet(&e.BC.Low, 0)
	case 0xC2:
		return e.CPU8BitRegisterSet(&e.DE.High, 0)
	case 0xC3:
		return e.CPU8BitRegisterSet(&e.DE.Low, 0)
	case 0xC4:
		return e.CPU8BitRegisterSet(&e.HL.High, 0)
	case 0xC5:
		return e.CPU8BitRegisterSet(&e.HL.Low, 0)
	case 0xC6:
		return e.CPU8BitSetMemoryAddress(e.HL.Value(), 0)
	// b=1
	case 0xCF:
		return e.CPU8BitRegisterSet(&e.AF.High, 1)
	case 0xC8:
		return e.CPU8BitRegisterSet(&e.BC.High, 1)
	case 0xC9:
		return e.CPU8BitRegisterSet(&e.BC.Low, 1)
	case 0xCA:
		return e.CPU8BitRegisterSet(&e.DE.High, 1)
	case 0xCB:
		return e.CPU8BitRegisterSet(&e.DE.Low, 1)
	case 0xCC:
		return e.CPU8BitRegisterSet(&e.HL.High, 1)
	case 0xCD:
		return e.CPU8BitRegisterSet(&e.HL.Low, 1)
	case 0xCE:
		return e.CPU8BitSetMemoryAddress(e.HL.Value(), 1)
	// b=2
	case 0xD7:
		return e.CPU8BitRegisterSet(&e.AF.High, 2)
	case 0xD0:
		return e.CPU8BitRegisterSet(&e.BC.High, 2)
	case 0xD1:
		return e.CPU8BitRegisterSet(&e.BC.Low, 2)
	case 0xD2:
		return e.CPU8BitRegisterSet(&e.DE.High, 2)
	case 0xD3:
		return e.CPU8BitRegisterSet(&e.DE.Low, 2)
	case 0xD4:
		return e.CPU8BitRegisterSet(&e.HL.High, 2)
	case 0xD5:
		return e.CPU8BitRegisterSet(&e.HL.Low, 2)
	case 0xD6:
		return e.CPU8BitSetMemoryAddress(e.HL.Value(), 2)
	// b=3
	case 0xDF:
		return e.CPU8BitRegisterSet(&e.AF.High, 3)
	case 0xD8:
		return e.CPU8BitRegisterSet(&e.BC.High, 3)
	case 0xD9:
		return e.CPU8BitRegisterSet(&e.BC.Low, 3)
	case 0xDA:
		return e.CPU8BitRegisterSet(&e.DE.High, 3)
	case 0xDB:
		return e.CPU8BitRegisterSet(&e.DE.Low, 3)
	case 0xDC:
		return e.CPU8BitRegisterSet(&e.HL.High, 3)
	case 0xDD:
		return e.CPU8BitRegisterSet(&e.HL.Low, 3)
	case 0xDE:
		return e.CPU8BitSetMemoryAddress(e.HL.Value(), 3)
	// b=4
	case 0xE7:
		return e.CPU8BitRegisterSet(&e.AF.High, 4)
	case 0xE0:
		return e.CPU8BitRegisterSet(&e.BC.High, 4)
	case 0xE1:
		return e.CPU8BitRegisterSet(&e.BC.Low, 4)
	case 0xE2:
		return e.CPU8BitRegisterSet(&e.DE.High, 4)
	case 0xE3:
		return e.CPU8BitRegisterSet(&e.DE.Low, 4)
	case 0xE4:
		return e.CPU8BitRegisterSet(&e.HL.High, 4)
	case 0xE5:
		return e.CPU8BitRegisterSet(&e.HL.Low, 4)
	case 0xE6:
		return e.CPU8BitSetMemoryAddress(e.HL.Value(), 4)
	// b=5
	case 0xEF:
		return e.CPU8BitRegisterSet(&e.AF.High, 5)
	case 0xE8:
		return e.CPU8BitRegisterSet(&e.BC.High, 5)
	case 0xE9:
		return e.CPU8BitRegisterSet(&e.BC.Low, 5)
	case 0xEA:
		return e.CPU8BitRegisterSet(&e.DE.High, 5)
	case 0xEB:
		return e.CPU8BitRegisterSet(&e.DE.Low, 5)
	case 0xEC:
		return e.CPU8BitRegisterSet(&e.HL.High, 5)
	case 0xED:
		return e.CPU8BitRegisterSet(&e.HL.Low, 5)
	case 0xEE:
		return e.CPU8BitSetMemoryAddress(e.HL.Value(), 5)
	// b=6
	case 0xF7:
		return e.CPU8BitRegisterSet(&e.AF.High, 6)
	case 0xF0:
		return e.CPU8BitRegisterSet(&e.BC.High, 6)
	case 0xF1:
		return e.CPU8BitRegisterSet(&e.BC.Low, 6)
	case 0xF2:
		return e.CPU8BitRegisterSet(&e.DE.High, 6)
	case 0xF3:
		return e.CPU8BitRegisterSet(&e.DE.Low, 6)
	case 0xF4:
		return e.CPU8BitRegisterSet(&e.HL.High, 6)
	case 0xF5:
		return e.CPU8BitRegisterSet(&e.HL.Low, 6)
	case 0xF6:
		return e.CPU8BitSetMemoryAddress(e.HL.Value(), 6)
	// b=7
	case 0xFF:
		return e.CPU8BitRegisterSet(&e.AF.High, 7)
	case 0xF8:
		return e.CPU8BitRegisterSet(&e.BC.High, 7)
	case 0xF9:
		return e.CPU8BitRegisterSet(&e.BC.Low, 7)
	case 0xFA:
		return e.CPU8BitRegisterSet(&e.DE.High, 7)
	case 0xFB:
		return e.CPU8BitRegisterSet(&e.DE.Low, 7)
	case 0xFC:
		return e.CPU8BitRegisterSet(&e.HL.High, 7)
	case 0xFD:
		return e.CPU8BitRegisterSet(&e.HL.Low, 7)
	case 0xFE:
		return e.CPU8BitSetMemoryAddress(e.HL.Value(), 7)
	}

	return 0
}
