package emulator

type Emulator struct {
	AF Register
	BC Register
	DE Register
	HL Register
}

func testBit(n uint8, pos uint) bool {
	mask := uint8(1) << pos
	return n&mask > 0
}

func setBit(n uint8, pos uint) uint8 {
	mask := uint8(1) << pos
	n |= mask
	return n
}

func clearBit(n uint8, pos uint) uint8 {
	mask := ^(uint8(1) << pos)
	n &= mask
	return n
}
