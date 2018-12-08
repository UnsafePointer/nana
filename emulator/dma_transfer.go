package emulator

import "fmt"

func (e *Emulator) DMATransfer(data uint8) {
	if e.EnableDebug {
		e.LogMessage(fmt.Sprintf("DMA Transfer: %#02x", data))
	}
	address := uint16(data) << 8
	for i := uint16(0); i < 0xA0; i++ {
		data := e.ReadMemory8Bit(address + i)
		e.WriteMemory(0xFE00+i, data)
	}
}
