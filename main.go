package main

import (
	"fmt"
	"os"

	"github.com/Ruenzuo/nana/emulator"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	width  = 160
	height = 144
	scale  = 10
)

var (
	e *emulator.Emulator
)

func update(screen *ebiten.Image) error {
	e.EmulateFrame()
	pixels := make([]byte, width*height*4)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			position := y*width + x
			pixels[(position*4 + 0)] = e.ScreenData[x][y][0]
			pixels[(position*4 + 1)] = e.ScreenData[x][y][1]
			pixels[(position*4 + 2)] = e.ScreenData[x][y][2]
			pixels[(position*4 + 3)] = 255
		}
	}
	screen.ReplacePixels(pixels)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %f", ebiten.CurrentFPS()))
	return nil
}

func main() {
	gameArg := os.Args[1]
	_, okDebug := os.LookupEnv("DEBUG")
	e = emulator.NewEmulator(okDebug)
	e.LoadCartridge(gameArg)
	if err := ebiten.Run(update, width, height, scale, fmt.Sprintf("nana - %s", gameArg)); err != nil {
		panic(err)
	}
}
