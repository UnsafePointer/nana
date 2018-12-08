package emulator

import "fmt"

// Viewing area/ Background map:
// 0xFF42: The Y Position of the BACKGROUND where to start drawing the viewing area from
// 0xFF43: The X Position of the BACKGROUND to start drawing the viewing area from
// 0xFF4A: The Y Position of the VIEWING AREA to start drawing the window from
// 0xFF4B: The X Positions -7 of the VIEWING AREA to start drawing the window from
const viewingAreaPositonYAddress = 0xFF42
const viewingAreaPositonXAddress = 0xFF43
const windowPositionYAddress = 0xFF4A
const windowPositionXAddress = 0xFF4B

func (e *Emulator) RenderTiles() {
	tileData := uint16(0)
	backgroundMemory := uint16(0)
	unsigned := true

	viewingAreaPositonY := e.ReadMemory8Bit(viewingAreaPositonYAddress)
	viewingAreaPositonX := e.ReadMemory8Bit(viewingAreaPositonXAddress)
	windowPositionY := e.ReadMemory8Bit(windowPositionYAddress)
	windowPositioX := e.ReadMemory8Bit(windowPositionXAddress) - 7

	lcdController := e.ReadMemory8Bit(lcdControllerAddress)
	currentScanline := e.ReadMemory8Bit(currentScanlineRegisterAddress)

	usingWindow := false

	if testBit(lcdController, 5) {
		if windowPositionY <= currentScanline {
			usingWindow = true
		}
	}

	if testBit(lcdController, 4) {
		tileData = 0x8000
	} else {
		tileData = 0x8800
		unsigned = false
	}

	if usingWindow {
		if testBit(lcdController, 6) {
			backgroundMemory = 0x9C00
		} else {
			backgroundMemory = 0x9800
		}
	} else {
		if testBit(lcdController, 3) {
			backgroundMemory = 0x9C00
		} else {
			backgroundMemory = 0x9800
		}
	}

	positionY := uint8(0)

	if usingWindow {
		positionY = currentScanline - windowPositionY
	} else {
		positionY = viewingAreaPositonY + currentScanline
	}

	tileRow := uint16(positionY/8) * 32

	for pixel := uint8(0); pixel < 160; pixel++ {
		positionX := pixel + viewingAreaPositonX

		if usingWindow {
			if pixel >= windowPositioX {
				positionX = pixel - windowPositioX
			}
		}

		tileColumn := uint16(positionX) / 8
		tileAddress := backgroundMemory + tileRow + tileColumn

		tileLocation := uint16(tileData)
		if unsigned {
			tileNumber := e.ReadMemory8Bit(tileAddress)
			tileLocation += uint16(tileNumber) * 16
		} else {
			tileNumber := int8(e.ReadMemory8Bit(tileAddress))
			tileLocation += uint16((int16(tileNumber) + 128) * 16)
		}

		line := positionY % 8
		line *= 2

		address1 := tileLocation + uint16(line)
		data1 := e.ReadMemory8Bit(address1)
		address2 := tileLocation + uint16(line) + 1
		data2 := e.ReadMemory8Bit(address2)

		if e.EnableLCDStateDebug {
			e.LogMessage(fmt.Sprintf("Background pixel: %d, data1: %#02x (%#04x), data2: %#02x (%#04x)", pixel, data1, address1, data2, address2))
		}

		colorBit := int(positionX) % 8
		colorBit -= 7
		colorBit *= -1

		colorValue := getBit(data2, uint(colorBit))
		colorValue <<= 1
		colorValue |= getBit(data1, uint(colorBit))

		color := e.GetColor(colorValue, 0xFF47)
		e.testPanic(color < 0 || color > 3, "Invalid color")

		red := uint8(0)
		green := uint8(0)
		blue := uint8(0)

		switch color {
		case 0:
			red = 255
			green = 255
			blue = 255
			break
		case 1:
			red = 0xCC
			green = 0xCC
			blue = 0xCC
			break
		case 2:
			red = 0x77
			green = 0x77
			blue = 0x77
			break
		}

		if currentScanline < 0 || currentScanline > 143 || pixel < 0 || pixel > 159 {
			continue
		}

		e.ScreenData[pixel][currentScanline][0] = red
		e.ScreenData[pixel][currentScanline][1] = green
		e.ScreenData[pixel][currentScanline][2] = blue
	}
}

func (e *Emulator) GetColor(color uint8, address uint16) uint8 {
	high := uint(0)
	low := uint(0)
	palette := e.ReadMemory8Bit(address)

	switch color {
	case 0:
		high = 1
		low = 0
		break
	case 1:
		high = 3
		low = 2
		break
	case 2:
		high = 5
		low = 4
		break
	case 3:
		high = 7
		low = 6
	}

	result := getBit(palette, high)
	result <<= 1
	result |= getBit(palette, low)

	return result
}

func (e *Emulator) RenderSprites() {
	lcdController := e.ReadMemory8Bit(lcdControllerAddress)
	currentScanline := e.ReadMemory8Bit(currentScanlineRegisterAddress)
	use8x16 := false
	if testBit(lcdController, 2) {
		use8x16 = true
	}

	for sprite := uint16(0); sprite < 40; sprite++ {
		spriteIndex := sprite * 4
		positionY := e.ReadMemory8Bit(0xFE00+spriteIndex) - 16
		positionX := e.ReadMemory8Bit(0xFE00+spriteIndex+1) - 8
		tileLocation := e.ReadMemory8Bit(0xFE00 + spriteIndex + 2)
		attributes := e.ReadMemory8Bit(0xFE00 + spriteIndex + 3)

		flipY := testBit(attributes, 6)
		flipX := testBit(attributes, 5)

		height := uint8(8)
		if use8x16 {
			height = 16
		}

		if currentScanline >= positionY && currentScanline < positionY+height {
			line := int32(currentScanline - positionY)

			if flipY {
				line -= int32(height)
				line *= -1
			}

			line *= 2
			dataAddress := int32((0x8000 + (int32(tileLocation) * 16))) + line
			data1 := e.ReadMemory8Bit(uint16(dataAddress))
			data2 := e.ReadMemory8Bit(uint16(dataAddress) + 1)

			for tilePixel := 7; tilePixel >= 0; tilePixel-- {
				colorBit := tilePixel
				if flipX {
					colorBit -= 7
					colorBit *= -1
				}

				colorValue := getBit(data2, uint(colorBit))
				colorValue <<= 1
				colorValue |= getBit(data1, uint(colorBit))

				colorAddress := uint16(0xFF48)
				if testBit(attributes, 4) {
					colorAddress = 0xFF49
				}

				color := e.GetColor(colorValue, colorAddress)
				e.testPanic(color < 0 || color > 3, "Invalid color")

				if color == 0 {
					continue
				}

				red := uint8(0)
				green := uint8(0)
				blue := uint8(0)

				switch color {
				case 0:
					red = 255
					green = 255
					blue = 255
					break
				case 1:
					red = 0xCC
					green = 0xCC
					blue = 0xCC
					break
				case 2:
					red = 0x77
					green = 0x77
					blue = 0x77
					break
				}

				pixelX := 0 - tilePixel
				pixelX += 7

				pixel := positionX + uint8(pixelX)

				if currentScanline < 0 || currentScanline > 143 || pixel < 0 || pixel > 159 {
					continue
				}

				if testBit(attributes, 7) {
					if e.ScreenData[pixel][currentScanline][0] != 255 || e.ScreenData[pixel][currentScanline][1] != 255 || e.ScreenData[pixel][currentScanline][2] != 255 {
						continue
					}
				}

				e.ScreenData[pixel][currentScanline][0] = red
				e.ScreenData[pixel][currentScanline][1] = green
				e.ScreenData[pixel][currentScanline][2] = blue
			}
		}
	}
}
