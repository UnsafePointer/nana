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
