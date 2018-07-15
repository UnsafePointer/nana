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
	}
	return 0
}
