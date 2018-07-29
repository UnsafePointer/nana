package emulator

func (e *Emulator) FlagZ() bool {
	return testBit(e.AF.Low.Value(), 7)
}

func (e *Emulator) FlagN() bool {
	return testBit(e.AF.Low.Value(), 6)
}

func (e *Emulator) FlagH() bool {
	return testBit(e.AF.Low.Value(), 5)
}

func (e *Emulator) FlagC() bool {
	return testBit(e.AF.Low.Value(), 4)
}

func (e *Emulator) SetFlagZ() {
	e.AF.SetLow(setBit(e.AF.Low.Value(), 7))
}

func (e *Emulator) SetFlagN() {
	e.AF.SetLow(setBit(e.AF.Low.Value(), 6))
}

func (e *Emulator) SetFlagH() {
	e.AF.SetLow(setBit(e.AF.Low.Value(), 5))
}

func (e *Emulator) SetFlagC() {
	e.AF.SetLow(setBit(e.AF.Low.Value(), 4))
}

func (e *Emulator) ClearFlagZ() {
	e.AF.SetLow(clearBit(e.AF.Low.Value(), 7))
}

func (e *Emulator) ClearFlagN() {
	e.AF.SetLow(clearBit(e.AF.Low.Value(), 6))
}

func (e *Emulator) ClearFlagH() {
	e.AF.SetLow(clearBit(e.AF.Low.Value(), 5))
}

func (e *Emulator) ClearFlagC() {
	e.AF.SetLow(clearBit(e.AF.Low.Value(), 4))
}

func (e *Emulator) ClearAllFlags() {
	e.ClearFlagZ()
	e.ClearFlagN()
	e.ClearFlagH()
	e.ClearFlagC()
}

func (e *Emulator) DebugFlags() string {
	var output string

	if e.FlagZ() {
		output += "1"
	} else {
		output += "0"
	}
	if e.FlagN() {
		output += "1"
	} else {
		output += "0"
	}
	if e.FlagH() {
		output += "1"
	} else {
		output += "0"
	}
	if e.FlagC() {
		output += "1"
	} else {
		output += "0"
	}

	return output
}
