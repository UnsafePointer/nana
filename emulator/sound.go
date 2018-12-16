package emulator

import (
	"github.com/veandco/go-sdl2/sdl"
)

const SampleRate = 44100
const MaxCyclesPerSample = MaxCyclesPerSecond / SampleRate
const SoundBufferSize = 4096

// Control/Status
// NR50 FF24 ALLL BRRR Vin L enable, Left vol, Vin R enable, Right vol
// NR51 FF25 NW21 NW21 Left enables, Right enables
// NR52 FF26 P--- NW21 Power control/status, Channel length statuses

func (e *Emulator) HandleSound(address uint16, data uint8) {
	if address >= 0xFF10 && address <= 0xFF14 {
		e.SquareOne.SetValue(address, data)
	} else if address >= 0xFF16 && address <= 0xFF19 {
		e.SquareTwo.SetValue(address, data)
	} else if address >= 0xFF1A && address <= 0xFF1E {
		e.Wave.SetValue(address, data)
	} else if address >= 0xFF20 && address <= 0xFF23 {
		e.Noise.SetValue(address, data)
	} else if address >= 0xFF24 && address <= 0xFF26 {
		switch address {
		case 0xFF24:
			e.RightVolume = data & 0x7
			e.LeftVolume = (data >> 4) & 0x7
			// A Vin L enable and B Vil R enable left unimplemented intentionally
			// I have no idea what they do
			break
		case 0xFF25:
			for i := uint8(0); i < 4; i++ {
				e.RightChannelEnable[i] = ((data >> i) & 0x1) == 0x1
			}
			for i := uint8(0); i < 4; i++ {
				e.LeftChannelEnable[i] = (data >> (i + 4) & 0x1) == 0x1
			}
			break
		case 0xFF26:
			e.SquareOne.LengthEnabled = ((data >> 0) & 0x1) == 0x1
			e.SquareTwo.LengthEnabled = ((data >> 1) & 0x1) == 0x1
			e.Wave.LengthEnabled = ((data >> 2) & 0x1) == 0x1
			e.Noise.LengthEnabled = ((data >> 3) & 0x1) == 0x1
			if (data & 0x80) != 0x80 {
				for i := uint16(0xFF10); i <= 0xFF25; i++ {
					e.WriteMemory(i, 0x00)
				}
				e.SoundEnabled = false
			} else if !e.SoundEnabled {
				for i := uint16(0); i < 16; i++ {
					e.Wave.SetValue(0xFF30|i, 0)
				}
				e.SoundEnabled = true
			}
			break
		}
	} else if address >= 0xFF30 && address <= 0xFF3F {
		e.Wave.SetValue(address, data)
	}
}

func (e *Emulator) UpdateSound(cycles int) {
	for i := 0; i < cycles; i++ {
		e.SquareOne.Step()
		e.SquareTwo.Step()
		e.Wave.Step()
		e.Noise.Step()

		e.SoundSampleCounter--
		if e.SoundSampleCounter <= 0 {
			e.SoundSampleCounter = MaxCyclesPerSample

			volume := uint8(0)
			// NW21
			if e.LeftChannelEnable[0] {
				volume = MixSoundSignal(volume, e.SquareOne.OutputVolume)
			}
			if e.LeftChannelEnable[1] {
				volume = MixSoundSignal(volume, e.SquareTwo.OutputVolume)
			}
			if e.LeftChannelEnable[2] {
				volume = MixSoundSignal(volume, e.Wave.OutputVolume)
			}
			if e.LeftChannelEnable[3] {
				volume = MixSoundSignal(volume, e.Noise.OutputVolume)
			}
			leftVolue := float32(e.LeftVolume) / float32(7)
			volume = uint8(float32(volume) * leftVolue)
			e.SoundBuffer = append(e.SoundBuffer, volume)

			volume = uint8(0)
			if e.RightChannelEnable[0] {
				volume = MixSoundSignal(volume, e.SquareOne.OutputVolume)
			}
			if e.RightChannelEnable[1] {
				volume = MixSoundSignal(volume, e.SquareTwo.OutputVolume)
			}
			if e.RightChannelEnable[2] {
				volume = MixSoundSignal(volume, e.Wave.OutputVolume)
			}
			if e.RightChannelEnable[3] {
				volume = MixSoundSignal(volume, e.Noise.OutputVolume)
			}
			rightVolume := float32(e.RightVolume) / float32(7)
			volume = uint8(float32(volume) * rightVolume)
			e.SoundBuffer = append(e.SoundBuffer, volume)
		}

		if len(e.SoundBuffer) >= SoundBufferSize {
			for sdl.GetQueuedAudioSize(1) > SoundBufferSize {
				sdl.Delay(1)
			}
			sdl.QueueAudio(1, e.SoundBuffer)
			e.SoundBuffer = e.SoundBuffer[:0]
		}
	}
}

// Credits for this go to Viktor T. Toth
// http://www.vttoth.com/CMS/index.php/technical-notes/68

func MixSoundSignal(a uint8, b uint8) uint8 {
	return (a + b) - a*b/255
}
