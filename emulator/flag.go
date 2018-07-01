package emulator

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
