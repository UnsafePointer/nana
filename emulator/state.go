package emulator

import (
	"encoding/gob"
	"os"
)

type State struct {
	AF Register16Bit
	BC Register16Bit
	DE Register16Bit
	HL Register16Bit

	SquareOne          SquareChannel
	SquareTwo          SquareChannel
	Wave               WaveChannel
	Noise              NoiseChannel
	RightVolume        uint8
	LeftVolume         uint8
	RightChannelEnable [4]bool
	LeftChannelEnable  [4]bool
	SoundEnabled       bool
	SoundSampleCounter int
	SoundBuffer        []uint8

	ROM            [0x10000]uint8
	RAM            [0x8000]uint8
	ProgramCounter Register16Bit
	StackPointer   Register16Bit

	CurrentROMBank           uint16
	CurrentRAMBank           uint16
	EnableRAMBank            bool
	EnableROMBank            bool
	Halted                   bool
	DisableInterrupts        bool
	PendingDisableInterrupts bool
	PendingEnableInterrupts  bool

	DividerRegisterCyclesCounter int
	TimerCyclesCounter           int

	ScanlineRenderCyclesCounter int
}

func NewState(e *Emulator) *State {
	s := new(State)
	s.AF = e.AF
	s.BC = e.BC
	s.DE = e.DE
	s.HL = e.HL
	s.SquareOne = e.SquareOne
	s.SquareTwo = e.SquareTwo
	s.Wave = e.Wave
	s.Noise = e.Noise
	s.RightVolume = e.RightVolume
	s.LeftVolume = e.LeftVolume
	s.RightChannelEnable = e.RightChannelEnable
	s.LeftChannelEnable = e.LeftChannelEnable
	s.SoundEnabled = e.SoundEnabled
	s.SoundSampleCounter = e.SoundSampleCounter
	s.SoundBuffer = e.SoundBuffer
	s.ROM = e.ROM
	s.RAM = e.RAM
	s.ProgramCounter = e.ProgramCounter
	s.StackPointer = e.StackPointer
	s.CurrentROMBank = e.CurrentROMBank
	s.CurrentRAMBank = e.CurrentRAMBank
	s.EnableRAMBank = e.EnableRAMBank
	s.EnableROMBank = e.EnableROMBank
	s.Halted = e.Halted
	s.DisableInterrupts = e.DisableInterrupts
	s.PendingDisableInterrupts = e.PendingDisableInterrupts
	s.PendingEnableInterrupts = e.PendingEnableInterrupts
	s.DividerRegisterCyclesCounter = e.DividerRegisterCyclesCounter
	s.TimerCyclesCounter = e.TimerCyclesCounter
	s.ScanlineRenderCyclesCounter = e.ScanlineRenderCyclesCounter
	return s
}

func (e *Emulator) CopyState(s *State) {
	e.AF = s.AF
	e.BC = s.BC
	e.DE = s.DE
	e.HL = s.HL
	e.SquareOne = s.SquareOne
	e.SquareTwo = s.SquareTwo
	e.Wave = s.Wave
	e.Noise = s.Noise
	e.RightVolume = s.RightVolume
	e.LeftVolume = s.LeftVolume
	e.RightChannelEnable = s.RightChannelEnable
	e.LeftChannelEnable = s.LeftChannelEnable
	e.SoundEnabled = s.SoundEnabled
	e.SoundSampleCounter = s.SoundSampleCounter
	e.SoundBuffer = s.SoundBuffer
	e.ROM = s.ROM
	e.RAM = s.RAM
	e.ProgramCounter = s.ProgramCounter
	e.StackPointer = s.StackPointer
	e.CurrentROMBank = s.CurrentROMBank
	e.CurrentRAMBank = s.CurrentRAMBank
	e.EnableRAMBank = s.EnableRAMBank
	e.EnableROMBank = s.EnableROMBank
	e.Halted = s.Halted
	e.DisableInterrupts = s.DisableInterrupts
	e.PendingDisableInterrupts = s.PendingDisableInterrupts
	e.PendingEnableInterrupts = s.PendingEnableInterrupts
	e.DividerRegisterCyclesCounter = s.DividerRegisterCyclesCounter
	e.TimerCyclesCounter = s.TimerCyclesCounter
	e.ScanlineRenderCyclesCounter = s.ScanlineRenderCyclesCounter
}

func (e *Emulator) SaveState() {
	s := NewState(e)
	if _, err := os.Stat("./save.sav"); err == nil {
		os.Remove("./save.sav")
	}
	file, err := os.Create("./save.sav")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(&s)
	if err != nil {
		panic(err)
	}
}

func (e *Emulator) LoadState() {
	s := new(State)
	if _, err := os.Stat("./save.sav"); err != nil {
		return
	}
	file, err := os.Open("./save.sav")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(s)
	if err != nil {
		panic(err)
	}
	e.CopyState(s)
}
