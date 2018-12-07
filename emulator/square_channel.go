package emulator

var dutyTable [4][8]bool

// 	Duty   Waveform    Ratio
// -------------------------
// 	0      00000001    12.5%
//  1      10000001    25%
//  2      10000111    50%
//  3      01111110    75%

func init() {
	dutyTable = [4][8]bool{
		{false, false, false, false, false, false, false, true},
		{true, false, false, false, false, false, false, true},
		{true, false, false, false, false, true, true, true},
		{false, true, true, true, true, true, true, false},
	}
}

type SquareChannel struct {
	SequenceCounter uint8
	DutyCounter     uint8

	Volume       uint8
	VolumeData   uint8
	OutputVolume uint8

	Timer         uint16
	FrequencyData uint16

	envelopeAddMode    bool
	envelopePeriodData uint8

	SweepShift      uint8
	SweepNegate     bool
	SweepPeriodData uint8

	LengthData    uint8
	LengthEnabled bool

	Trigger bool

	Enabled    bool
	DACEnabled bool
}

// Square 1
// NR10 FF10 -PPP NSSS Sweep period, negate, shift
// NR11 FF11 DDLL LLLL Duty, Length load (64-L)
// NR12 FF12 VVVV APPP Starting volume, Envelope add mode, period
// NR13 FF13 FFFF FFFF Frequency LSB
// NR14 FF14 TL-- -FFF Trigger, Length enable, Frequency MSB
//
// Square 2
//      FF15 ---- ---- Not used
// NR21 FF16 DDLL LLLL Duty, Length load (64-L)
// NR22 FF17 VVVV APPP Starting volume, Envelope add mode, period
// NR23 FF18 FFFF FFFF Frequency LSB
// NR24 FF19 TL-- -FFF Trigger, Length enable, Frequency MSB

func (c *SquareChannel) SetValue(address uint16, data uint8) {
	switch address {
	// Sweep period, negate, shift
	case 0xFF10:
		c.SweepPeriodData = (data >> 4) & 0x7
		c.SweepNegate = (data & 0x8) == 0x8
		c.SweepShift = data & 0x7
		break
	// Duty, Length load (64-L)
	case 0xFF11:
		fallthrough
	case 0xFF16:
		c.DutyCounter = (data >> 6) & 0x3
		c.LengthData = data & 0x3F
		break
	// Starting volume, Envelope add mode, period
	case 0xFF12:
		fallthrough
	case 0xFF17:
		// DAC
		c.DACEnabled = (data & 0xF8) != 0
		c.VolumeData = (data >> 4) & 0xF
		c.envelopeAddMode = (data & 0x8) == 0x8
		c.envelopePeriodData = (data & 0x7)
		c.Volume = c.VolumeData
		break
	// Frequency LSB
	case 0xFF13:
		fallthrough
	case 0xFF18:
		c.FrequencyData = (c.FrequencyData & 0x700) | uint16(data)
		break
	// Trigger, Length enable, Frequency MSB
	case 0xFF14:
		fallthrough
	case 0xFF19:
		c.FrequencyData = (c.FrequencyData & 0xFF) | (uint16(data)&0x7)<<8
		c.LengthEnabled = (data & 0x40) == 0x40
		c.Trigger = (data & 0x80) == 0x80
		if c.Trigger {
			c.executeTrigger()
		}
		break
	}
}

func (c *SquareChannel) Step() {
	c.Timer--
	if c.Timer <= 0 {
		c.Timer = (2048 - c.FrequencyData) * 4
		c.SequenceCounter = (c.SequenceCounter + 1) & 0x7
	}

	if c.Enabled && c.DACEnabled && dutyTable[c.DutyCounter][c.SequenceCounter] {
		c.OutputVolume = c.Volume
	} else {
		c.OutputVolume = 0
	}
}

func (c *SquareChannel) executeTrigger() {
	c.Enabled = true
	c.Timer = (2048 - c.FrequencyData) * 4
	c.Volume = c.VolumeData
}
