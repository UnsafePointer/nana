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

func (r *Register) SetValue(value uint16) {
	r.SetLow(uint8(value))
	r.SetHigh(uint8((value >> 8)))
}
