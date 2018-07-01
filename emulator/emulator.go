package emulator

import (
	"io/ioutil"
	"os"
)

type Emulator struct {
	AF Register
	BC Register
	DE Register
	HL Register

	CartridgeMemory [0x200000]uint8
}

func (e *Emulator) LoadCartridge(filename string) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}

	copy(e.CartridgeMemory[:], dat)
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
