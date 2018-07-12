package emulator

func (e *Emulator) ExecuteOpCode(opcode uint8) int {
	switch opcode {
	// no-op
	case 0x00:
		return 4
	}

	return 0
}
