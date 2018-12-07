package emulator

var divisors [8]uint16

// Divisor code   Divisor
// -----------------------
//    0             8
//    1            16
//    2            32
//    3            48
//    4            64
//    5            80
//    6            96
//    7           112

func init() {
	divisors = [8]uint16{8, 16, 32, 48, 64, 80, 96, 112}
}

type NoiseChannel struct {
	Volume       uint8
	VolumeData   uint8
	OutputVolume uint8

	Timer uint16

	envelopeAddMode    bool
	envelopePeriodData uint8

	ClockShift uint8

	WidthEnabled bool

	DivisorData uint8

	LengthData    uint8
	LengthEnabled bool

	Trigger bool

	Enabled    bool
	DACEnabled bool

	LFSR int16
}

// Noise
// FF1F ---- ---- Not used
// NR41 FF20 --LL LLLL Length load (64-L)
// NR42 FF21 VVVV APPP Starting volume, Envelope add mode, period
// NR43 FF22 SSSS WDDD Clock shift, Width mode of LFSR, Divisor code
// NR44 FF23 TL-- ---- Trigger, Length enable

func (c *NoiseChannel) SetValue(address uint16, data uint8) {
	switch address {
	case 0xFF20:
		c.LengthData = data
		break
	case 0xFF21:
		// DAC
		c.DACEnabled = (data & 0xF8) != 0
		c.VolumeData = (data >> 4) & 0xF
		c.envelopeAddMode = (data & 0x8) == 0x8
		c.envelopePeriodData = (data & 0x7)
		break
	case 0xFF22:
		c.ClockShift = (data >> 4) & 0xF
		c.WidthEnabled = (data & 0x8) == 0x8
		c.DivisorData = data & 0x7
		break
	case 0xFF23:
		c.LengthEnabled = (data & 0x40) == 0x40
		c.Trigger = (data & 0x80) == 0x80
		if c.Trigger {
			c.executeTrigger()
		}
		break
	}
}

func (c *NoiseChannel) Step() {
	c.Timer--

	if c.Timer <= 0 {
		c.Timer = divisors[c.DivisorData] << c.ClockShift

		value := (c.LFSR & 0x1) ^ ((c.LFSR >> 1) & 0x1)
		c.LFSR >>= 1
		c.LFSR |= value << 14
		if c.WidthEnabled {
			c.LFSR &= ^0x40
			c.LFSR |= value << 6
		}
		if c.Enabled && c.DACEnabled && (c.LFSR&0x1) == 0 {
			c.OutputVolume = c.Volume
		} else {
			c.OutputVolume = 0
		}
	}
}

func (c *NoiseChannel) executeTrigger() {
	c.Enabled = true
	c.Timer = divisors[c.DivisorData] << c.ClockShift
	c.Volume = c.VolumeData
	c.LFSR = 0x7FFF
}
