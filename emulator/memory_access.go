package emulator

const ROMBankSize = 0x4000
const RAMBankSize = 0x2000
const DMATransferAddress = 0xFF46

func (e *Emulator) WriteMemory(address uint16, data uint8) {
	// Memory map:
	// 0000-7FFF ROM
	// E000-FDFF Same as C000-DDFF (ECHO) (typically not used)
	// FEA0-FEFF Not Usable
	if address <= 0x7FFF {
		e.HandleMemoryBanking(address, data)
		return
	} else if address >= 0xA000 && address <= 0xBFFF {
		if e.EnableRAMBank {
			bankAddress := address - 0xA000
			e.RAM[bankAddress+(e.CurrentRAMBank*RAMBankSize)] = data
		}
	} else if address >= 0xE000 && address <= 0xFDFF {
		e.ROM[address] = data
		e.WriteMemory(address-0x2000, data)
	} else if address >= 0xFEA0 && address <= 0xFEFF {
		return
	} else if address == dividerRegisterAddress {
		// Divider register trap
		e.ROM[address] = 0
	} else if address == timerControllerAddress {
		currentClockFrequency := e.ClockFrequency()
		e.ROM[address] = data
		newClockFrequency := e.ClockFrequency()
		if currentClockFrequency != newClockFrequency {
			e.SetTimerCycleCounter()
		}
	} else if address == currentScanlineRegisterAddress {
		// Current scanline register trap
		e.ROM[address] = 0
	} else if address == DMATransferAddress {
		e.DMATransfer(data)
	} else {
		e.ROM[address] = data
	}
}

func (e *Emulator) ReadMemory8Bit(address uint16) uint8 {
	// Memory map:
	// 0000-3FFF 16KB ROM Bank 00 (in cartridge, fixed at bank 00)
	// 4000-7FFF 16KB ROM Bank 01..NN (in cartridge, switchable bank number)
	// A000-BFFF 8KB External RAM (in cartridge, switchable bank, if any)
	if address >= 0x4000 && address <= 0x7FFF {
		bankAddress := address - 0x4000
		return e.CartridgeMemory[bankAddress+e.CurrentROMBank*ROMBankSize]
	} else if address >= 0xA000 && address <= 0xBFFF {
		bankAddress := address - 0xA000
		return e.RAM[bankAddress+e.CurrentRAMBank*RAMBankSize]
	}

	return e.ROM[address]
}

func (e *Emulator) ReadMemory16Bit(address uint16) uint16 {
	high := uint16(e.ReadMemory8Bit(address + 1))
	high <<= 8
	low := uint16(e.ReadMemory8Bit(address))
	return low | high
}

func (e *Emulator) HandleMemoryBanking(address uint16, data uint8) {
	if address < 0x2000 {
		if e.CartridgeType != CartridgeTypeROMOnly {
			e.EnableRAMBanking(address, data)
		}
	} else if address >= 0x2000 && address < 0x4000 {
		if e.CartridgeType != CartridgeTypeROMOnly {
			e.ChangeLowROMBank(data)
		}
	} else if address >= 0x4000 && address < 0x6000 {
		if e.CartridgeType == CartridgeTypeMBC1 {
			if e.EnableROMBank {
				e.ChangeHighROMBank(data)
			} else {
				e.ChangeRAMBank(data)
			}
		}
	} else if address >= 0x6000 && address < 0x8000 {
		if e.CartridgeType == CartridgeTypeMBC1 {
			e.SelectMemoryBankingMode(data)
		}
	}
}

func (e *Emulator) EnableRAMBanking(address uint16, data uint8) {
	if e.CartridgeType == CartridgeTypeMBC2 {
		if testBit(uint8(address), 4) {
			return
		}
	}
	test := data & 0xF
	if test == 0xA {
		e.EnableRAMBank = true
	} else {
		e.EnableRAMBank = false
	}
}

func (e *Emulator) ChangeRAMBank(data uint8) {
	e.CurrentRAMBank = uint16(data & 0x3)
}

func (e *Emulator) ChangeLowROMBank(data uint8) {
	if e.CartridgeType == CartridgeTypeMBC2 {
		e.CurrentROMBank = uint16(data & 0xF)
		if e.CurrentROMBank == 0 {
			e.CurrentROMBank++
		}
	} else {
		test := uint16(data & 31)
		e.CurrentROMBank &= 244
		e.CurrentROMBank |= test
		if e.CurrentROMBank == 0 {
			e.CurrentROMBank++
		}
	}
}

func (e *Emulator) ChangeHighROMBank(data uint8) {
	e.CurrentROMBank &= 31
	data &= 224
	e.CurrentROMBank |= uint16(data)
	if e.CurrentROMBank == 0 {
		e.CurrentROMBank++
	}
}

func (e *Emulator) SelectMemoryBankingMode(data uint8) {
	if data&0x1 == 0 {
		e.EnableROMBank = true
	} else {
		e.EnableROMBank = false
	}
	if e.EnableROMBank {
		e.CurrentRAMBank = 0
	}
}
