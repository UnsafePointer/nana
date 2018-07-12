package emulator

func (e *Emulator) ExecuteOpCode(opcode uint8) int {
	switch opcode {
	// no-op
	case 0x00:
		return 4
	// 8-Bit Loads
	case 0x06:
		return e.CPU8BitLoad(&e.BC, true)
	case 0x0E:
		return e.CPU8BitLoad(&e.BC, false)
	case 0x16:
		return e.CPU8BitLoad(&e.DE, true)
	case 0x1E:
		return e.CPU8BitLoad(&e.DE, false)
	case 0x26:
		return e.CPU8BitLoad(&e.HL, true)
	case 0x2E:
		return e.CPU8BitLoad(&e.HL, false)
	}

	return 0
}
