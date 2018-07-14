package emulator

type Register8Bit struct {
	bits uint8
}

type Register16Bit struct {
	low  Register8Bit
	high Register8Bit
}

func (r Register8Bit) Value() uint8 {
	return r.bits
}

func (r *Register8Bit) SetValue(value uint8) {
	r.bits = value
}

func (r Register16Bit) Low() uint8 {
	return r.low.bits
}

func (r Register16Bit) High() uint8 {
	return r.high.bits
}

func (r *Register16Bit) SetLow(value uint8) {
	r.low.SetValue(value)
}

func (r *Register16Bit) SetHigh(value uint8) {
	r.high.SetValue(value)
}

func (r Register16Bit) Value() uint16 {
	low := uint16(r.low.bits)
	low <<= 0
	high := uint16(r.high.bits)
	high <<= 8
	return low | high
}

func (r *Register16Bit) SetValue(value uint16) {
	r.SetLow(uint8(value & 0xFF))
	r.SetHigh(uint8((value >> 8)))
}

func (r *Register16Bit) Increment() {
	r.SetValue(r.Value() + 1)
}

func (r *Register16Bit) Decrement() {
	r.SetValue(r.Value() - 1)
}
