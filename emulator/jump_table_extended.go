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
	// SRL n
	case 0x2F:
		return e.CPU8BitRegisterSRL(&e.AF.High)
	case 0x28:
		return e.CPU8BitRegisterSRL(&e.BC.High)
	case 0x29:
		return e.CPU8BitRegisterSRL(&e.BC.Low)
	case 0x2A:
		return e.CPU8BitRegisterSRL(&e.DE.High)
	case 0x2B:
		return e.CPU8BitRegisterSRL(&e.DE.Low)
	case 0x2C:
		return e.CPU8BitRegisterSRL(&e.HL.High)
	case 0x2D:
		return e.CPU8BitRegisterSRL(&e.HL.Low)
	case 0x2E:
		return e.CPU8BitSRLMemoryAddress(e.HL.Value())

	}

	return 0
}
