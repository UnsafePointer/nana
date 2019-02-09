package emulator

type Register8Bit struct {
	Bits uint8
}

type Register16Bit struct {
	Low  Register8Bit
	High Register8Bit
}

func (r Register8Bit) Value() uint8 {
	return r.Bits
}

func (r *Register8Bit) SetValue(value uint8) {
	r.Bits = value
}

func (r *Register16Bit) SetLow(value uint8) {
	r.Low.SetValue(value)
}

func (r *Register16Bit) SetHigh(value uint8) {
	r.High.SetValue(value)
}

func (r Register16Bit) Value() uint16 {
	low := uint16(r.Low.Bits)
	low <<= 0
	high := uint16(r.High.Bits)
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
