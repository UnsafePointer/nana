package emulator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

const MaxCyclesPerSecond = 4194304
const MaxCyclesPerEmulationCycle = MaxCyclesPerSecond / 60 // Target is 60 FPS

type CartridgeType int

const (
	CartridgeTypeROMOnly CartridgeType = 0
	CartridgeTypeMBC1    CartridgeType = 1
	CartridgeTypeMBC2    CartridgeType = 2
	// CartridgeTypeMBC3 MemoryBankController = 3
	// CartridgeTypeMBC5 MemoryBankController = 5
)

type Emulator struct {
	AF Register16Bit
	BC Register16Bit
	DE Register16Bit
	HL Register16Bit

	CartridgeMemory          [0x200000]uint8
	ROM                      [0x10000]uint8
	RAM                      [0x8000]uint8
	ProgramCounter           Register16Bit
	StackPointer             Register16Bit
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
	ScreenData                  [160][144][3]uint8

	EnableDebug        bool
	LogBuffer          bytes.Buffer
	MaxCycles          int
	TotalCycles        int
	InstructionCounter map[uint8]int

	CartridgeType CartridgeType
}

func NewEmulator(enableDebug bool, maxCycles int) *Emulator {
	e := new(Emulator)
	if enableDebug {
		e.SetupLogFile()
	}
	e.EnableDebug = enableDebug
	e.MaxCycles = maxCycles
	e.InstructionCounter = make(map[uint8]int)
	e.ProgramCounter.SetValue(0x100)
	e.AF.SetValue(0x01B0)
	e.BC.SetValue(0x0013)
	e.DE.SetValue(0x00D8)
	e.HL.SetValue(0x014D)
	e.StackPointer.SetValue(0xFFFE)
	e.CurrentROMBank = 1 // Should never be 1, ROM bank 0 is fixed
	e.CurrentRAMBank = 0
	e.EnableRAMBank = false
	e.EnableROMBank = false
	e.Halted = false
	e.DisableInterrupts = true
	e.PendingDisableInterrupts = false
	e.PendingEnableInterrupts = false
	e.DividerRegisterCyclesCounter = 0
	e.TimerCyclesCounter = 0
	e.ScanlineRenderCyclesCounter = 456
	e.CartridgeType = CartridgeTypeROMOnly
	e.ROM[0xFF05] = 0x00
	e.ROM[0xFF06] = 0x00
	e.ROM[0xFF07] = 0x00
	e.ROM[0xFF10] = 0x80
	e.ROM[0xFF11] = 0xBF
	e.ROM[0xFF12] = 0xF3
	e.ROM[0xFF14] = 0xBF
	e.ROM[0xFF16] = 0x3F
	e.ROM[0xFF17] = 0x00
	e.ROM[0xFF19] = 0xBF
	e.ROM[0xFF1A] = 0x7F
	e.ROM[0xFF1B] = 0xFF
	e.ROM[0xFF1C] = 0x9F
	e.ROM[0xFF1E] = 0xBF
	e.ROM[0xFF20] = 0xFF
	e.ROM[0xFF21] = 0x00
	e.ROM[0xFF22] = 0x00
	e.ROM[0xFF23] = 0xBF
	e.ROM[0xFF24] = 0x77
	e.ROM[0xFF25] = 0xF3
	e.ROM[0xFF26] = 0xF1
	e.ROM[0xFF40] = 0x91
	e.ROM[0xFF42] = 0x00
	e.ROM[0xFF43] = 0x00
	e.ROM[0xFF45] = 0x00
	e.ROM[0xFF47] = 0xFC
	e.ROM[0xFF48] = 0xFF
	e.ROM[0xFF49] = 0xFF
	e.ROM[0xFF4A] = 0x00
	e.ROM[0xFF4B] = 0x00
	e.ROM[0xFFFF] = 0x00
	return e
}

func (e *Emulator) LoadCartridge(filename string) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}

	copy(e.CartridgeMemory[:], dat)
	for i := 0; i < 0x8000; i++ {
		e.ROM[i] = e.CartridgeMemory[i]
	}

	cartridgeTypeDefinition := e.CartridgeMemory[0x0147]
	switch cartridgeTypeDefinition {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		e.CartridgeType = CartridgeTypeMBC1
	case 5:
		fallthrough
	case 6:
		e.CartridgeType = CartridgeTypeMBC2
	}
}

func (e *Emulator) EmulateFrame() {
	cyclesThisUpdate := 0
	for cyclesThisUpdate < MaxCyclesPerEmulationCycle {
		cycles := e.executeNextOpcode()
		cyclesThisUpdate += cycles
		e.UpdateTimers(cycles)
		e.UpdateScreen(cycles)
		e.ExecuteInterrupts()
		if e.MaxCycles != 0 {
			e.TotalCycles += cycles
			if e.TotalCycles >= e.MaxCycles {
				os.Exit(0)
			}
		}
	}
}

func (e *Emulator) executeNextOpcode() int {
	opCode := e.ReadMemory8Bit(e.ProgramCounter.Value())
	e.CountOperationCode(opCode)
	e.ProgramCounter.Increment()
	var cycles int
	if e.Halted {
		cycles = 4
	} else {
		cycles = e.ExecuteOpCode(opCode)
		e.LogMessage(fmt.Sprintf("OP: %#02x, Cycles: %02d, Program Counter: %#04x", opCode, cycles, e.ProgramCounter.Value()-1))
	}
	// 0xF3: disable interrupts but only after next instruction, so
	// no immediatly after we return from 0xF3
	if e.PendingDisableInterrupts && opCode != 0xF3 {
		e.PendingDisableInterrupts = false
		e.DisableInterrupts = true
	}
	// 0xFB: enable interrupts but only after next instruction
	if e.PendingEnableInterrupts && opCode != 0xFB {
		e.PendingEnableInterrupts = false
		e.DisableInterrupts = false
	}
	return cycles
}

func (e *Emulator) CountOperationCode(opcode uint8) {
	count, ok := e.InstructionCounter[opcode]
	if ok == false {
		count = 0
	}
	count++
	e.InstructionCounter[opcode] = count
}

func testBit(n uint8, pos uint) bool {
	mask := uint8(1) << pos
	return n&mask > 0
}

func setBit(n uint8, pos uint) uint8 {
	mask := uint8(1) << pos
	n |= mask
	return n
}

func clearBit(n uint8, pos uint) uint8 {
	mask := ^(uint8(1) << pos)
	n &= mask
	return n
}

func getBit(n uint8, pos uint) uint8 {
	mask := uint8(1) << pos
	return n & mask
}
