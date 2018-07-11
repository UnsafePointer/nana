package emulator

const ROMBankSize = 0x4000
const RAMBankSize = 0x2000

func (e *Emulator) WriteMemory(address uint16, data uint8) {
	// Memory map:
	// 0000-7FFF ROM
	// E000-FDFF Same as C000-DDFF (ECHO) (typically not used)
	// FEA0-FEFF Not Usable
	if address <= 0x7FFF {
		return
	} else if address >= 0xE000 && address <= 0xFDFF {
		e.ROM[address] = data
		e.WriteMemory(address-0x2000, data)
	} else if address >= 0xFEA0 && address <= 0xFEFF {
		return
	} else {
		e.ROM[address] = data
	}
}

func (e Emulator) ReadMemory(address uint16) uint8 {
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
