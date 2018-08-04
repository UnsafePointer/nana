package emulator

import (
	"fmt"
)

const lcdControllerAddress = 0xFF40
const lcdStatusRegisterAddress = 0xFF41
const currentScanlineRegisterAddress = 0xFF44
const coincidenceFlagAddress = 0xFF45

func (e *Emulator) UpdateScreen(cycles int) {
	e.updateLCDStatus()

	if !e.IsLCDEnabled() {
		return
	}
	e.ScanlineRenderCyclesCounter -= cycles
	if e.ScanlineRenderCyclesCounter > 0 {
		return
	}
	// WriteMemory has a trap for this address
	// so we can't use it
	e.ROM[currentScanlineRegisterAddress]++
	currentScanline := e.ReadMemory8Bit(currentScanlineRegisterAddress)
	e.ScanlineRenderCyclesCounter = 456
	e.LogMessage(fmt.Sprintf("Current scanline: %d", currentScanline))

	if currentScanline == 144 {
		e.RequestInterrupt(0)
	} else if currentScanline > 153 {
		e.ROM[currentScanlineRegisterAddress] = 0
	} else if currentScanline < 144 {
		e.DrawScanline()
	}
}

func (e *Emulator) DrawScanline() {
	value := e.ReadMemory8Bit(lcdControllerAddress)
	if testBit(value, 0) {
		e.RenderTiles()
	}
	if testBit(value, 1) {
		e.RenderSprites()
	}
}

//
// LCD Status Register map
// Bit 0-1:
// 00: H-Blank (Mode 0)
// 01: V-Blank (Mode 1)
// 10: Searching Sprites Atts (Mode 2)
// 11: Transfering Data to LCD Driver (Mode 3)
// Bit 3: Mode 0 Interupt Enabled - Cycles 251-455
// Bit 4: Mode 1 Interupt Enabled - Cycles 80-251
// Bit 5: Mode 2 Interupt Enabled - Cycles 0-79
// Bit 6: Coincidence Interrupt Enabled
func (e *Emulator) updateLCDStatus() {
	status := e.ReadMemory8Bit(lcdStatusRegisterAddress)
	if !e.IsLCDEnabled() {
		e.ScanlineRenderCyclesCounter = 456
		e.ROM[currentScanlineRegisterAddress] = 0
		status &= 252
		status = setBit(status, 0)
		e.WriteMemory(lcdStatusRegisterAddress, status)
		return
	}

	currentScanline := e.ReadMemory8Bit(currentScanlineRegisterAddress)
	currentMode := status & 0x3

	mode := uint8(0)
	requestInterrupt := false

	if currentScanline >= 144 {
		mode = 1
		status = setBit(status, 0)
		status = clearBit(status, 1)
		requestInterrupt = testBit(status, 4)
	} else {
		mode2bounds := 456 - 80
		mode3bounds := mode2bounds - 172

		if e.ScanlineRenderCyclesCounter >= mode2bounds {
			mode = 2
			status = clearBit(status, 0)
			status = setBit(status, 1)
			requestInterrupt = testBit(status, 5)
		} else if e.ScanlineRenderCyclesCounter >= mode3bounds {
			mode = 3
			status = setBit(status, 0)
			status = setBit(status, 1)
		} else {
			mode = 0
			status = clearBit(status, 0)
			status = clearBit(status, 1)
			requestInterrupt = testBit(status, 3)
		}
	}

	if requestInterrupt && mode != currentMode {
		e.RequestInterrupt(1)
	}

	coincidenceScanline := e.ReadMemory8Bit(coincidenceFlagAddress)
	if currentScanline == coincidenceScanline {
		status = setBit(status, 2)
		if testBit(status, 6) {
			e.RequestInterrupt(1)
		}
	} else {
		status = clearBit(status, 2)
	}

	e.WriteMemory(lcdStatusRegisterAddress, status)
}

func (e *Emulator) IsLCDEnabled() bool {
	value := e.ReadMemory8Bit(lcdControllerAddress)
	return testBit(value, 7)
}
