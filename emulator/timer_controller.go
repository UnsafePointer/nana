package emulator

const dividerRegisterAddress = 0xFF04
const timerCounterAddress = 0xFF05
const timerModulatorAddress = 0xFF06
const timerControllerAddress = 0xFF07

func (e *Emulator) IsTimerEnabled() bool {
	value := e.ReadMemory8Bit(timerControllerAddress)
	return testBit(value, 2)
}

func (e *Emulator) UpdateTimers(cycles int) {
	e.updateDividerRegister(cycles)

	if !e.IsTimerEnabled() {
		return
	}

	e.TimerCyclesCounter -= cycles
	if e.TimerCyclesCounter > 0 {
		return
	}

	e.SetTimerCycleCounter()

	timerCounter := e.ReadMemory8Bit(timerCounterAddress)
	if timerCounter == 255 {
		timerModulator := e.ReadMemory8Bit(timerModulatorAddress)
		e.WriteMemory(timerCounterAddress, timerModulator)
		e.RequestInterrupt(2)
	} else {
		e.WriteMemory(timerCounterAddress, timerCounter+1)
	}
}

func (e *Emulator) ClockFrequency() uint8 {
	value := e.ReadMemory8Bit(timerControllerAddress)
	return value & 0x3
}

func (e *Emulator) SetTimerCycleCounter() {
	frequency := e.ClockFrequency()
	// Frequency map:
	// 00: 4096 Hz
	// 01: 262144 Hz
	// 10: 65536 Hz
	// 11: 16384 Hz
	switch frequency {
	case 0:
		e.TimerCyclesCounter = 1024
		break
	case 1:
		e.TimerCyclesCounter = 16
		break
	case 2:
		e.TimerCyclesCounter = 64
		break
	case 3:
		e.TimerCyclesCounter = 256
		break
	}
}

func (e *Emulator) updateDividerRegister(cycles int) {
	e.DividerRegisterCyclesCounter += cycles
	if e.DividerRegisterCyclesCounter >= 256 {
		e.DividerRegisterCyclesCounter = 0
		// WriteMemory has a trap for this address
		// so we can't use it
		e.ROM[dividerRegisterAddress]++
	}
}
