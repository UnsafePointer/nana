package emulator

type Register struct {
	bits [2]uint8
}

func (r Register) Low() uint8 {
	return r.bits[0]
}

func (r Register) High() uint8 {
	return r.bits[1]
}

func (r *Register) SetLow(value uint8) {
	r.bits[0] = value
}

func (r *Register) SetHigh(value uint8) {
	r.bits[1] = value
}

func (r Register) Value() uint16 {
	low := uint16(r.bits[0])
	low <<= 0
	high := uint16(r.bits[1])
	high <<= 8
	return low | high
}

type Emulator struct {
	AF Register
	BC Register
	DE Register
	HL Register
}

func (e Emulator) FlagZ() bool {
	return testBit(e.AF.Low(), 7)
}

func (e Emulator) FlagN() bool {
	return testBit(e.AF.Low(), 6)
}

func (e Emulator) FlagH() bool {
	return testBit(e.AF.Low(), 5)
}

func (e Emulator) FlagC() bool {
	return testBit(e.AF.Low(), 4)
}

func (e *Emulator) SetFlagZ() {
	e.AF.SetLow(setBit(e.AF.Low(), 7))
}

func (e *Emulator) SetFlagN() {
	e.AF.SetLow(setBit(e.AF.Low(), 6))
}

func (e *Emulator) SetFlagH() {
	e.AF.SetLow(setBit(e.AF.Low(), 5))
}

func (e *Emulator) SetFlagC() {
	e.AF.SetLow(setBit(e.AF.Low(), 4))
}

func (e *Emulator) ClearFlagZ() {
	e.AF.SetLow(clearBit(e.AF.Low(), 7))
}

func (e *Emulator) ClearFlagN() {
	e.AF.SetLow(clearBit(e.AF.Low(), 6))
}

func (e *Emulator) ClearFlagH() {
	e.AF.SetLow(clearBit(e.AF.Low(), 5))
}

func (e *Emulator) ClearFlagC() {
	e.AF.SetLow(clearBit(e.AF.Low(), 4))
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
