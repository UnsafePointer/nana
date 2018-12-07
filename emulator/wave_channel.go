package emulator

type WaveChannel struct {
	VolumeData   uint8
	OutputVolume uint8

	Timer           uint16
	FrequencyData   uint16
	PositionCounter uint8

	LengthData    uint8
	LengthEnabled bool

	Trigger bool

	Enabled    bool
	DACEnabled bool

	WaveTable [16]uint8
}

// Wave
// NR30 FF1A E--- ---- DAC power
// NR31 FF1B LLLL LLLL Length load (256-L)
// NR32 FF1C -VV- ---- Volume code (00=0%, 01=100%, 10=50%, 11=25%)
// NR33 FF1D FFFF FFFF Frequency LSB
// NR34 FF1E TL-- -FFF Trigger, Length enable, Frequency MSB

func (c *WaveChannel) SetValue(address uint16, data uint8) {
	switch address {
	case 0xFF1A:
		c.DACEnabled = (data & 0x80) == 0x80
		break
	case 0xFF1B:
		c.LengthData = data
		break
	case 0xFF1C:
		c.VolumeData = (data >> 5) & 0x3
		break
	case 0xFF1D:
		c.FrequencyData = (c.FrequencyData & 0x700) | uint16(data)
		break
	case 0xFF1E:
		c.FrequencyData = (c.FrequencyData & 0xFF) | (uint16(data)&0x7)<<8
		c.LengthEnabled = (data & 0x40) == 0x40
		c.Trigger = (data & 0x80) == 0x80
		if c.Trigger {
			c.executeTrigger()
		}
		break
	}
	if address >= 0xFF30 && address <= 0xFF3F {
		position := address & 0xF
		c.WaveTable[position] = data
	}
}

func (c *WaveChannel) Step() {
	c.Timer--
	if c.Timer <= 0 {
		c.Timer = (2048 - c.FrequencyData) * 2
		c.PositionCounter = (c.PositionCounter + 1) & 0x1F

		// Unsure if this happens only on the loop back!
		if !c.Enabled || !c.DACEnabled {
			c.OutputVolume = 0
			return
		}
		// Table has 32 4-bit values stored in a
		// 16 8-bit array
		// Counter is counting 32 positions
		position := c.PositionCounter / 2
		valueByte := c.WaveTable[position]
		// Select either the 4 high or lower bits
		if position&0x1 == 0 {
			valueByte >>= 4
		}
		valueByte &= 0xF
		// Volume code (00=0%, 01=100%, 10=50%, 11=25%)
		if c.VolumeData > 0 {
			valueByte >>= (c.VolumeData - 1)
		} else {
			valueByte = 0
		}
		c.OutputVolume = valueByte
	}
}

func (c *WaveChannel) executeTrigger() {
	c.Enabled = true
	c.Timer = (2048 - c.FrequencyData) * 2
	c.PositionCounter = 0
}
