package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Ruenzuo/nana/emulator"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	width  = 160
	height = 144
	scale  = 5
)

var (
	e *emulator.Emulator
)

func update(screen *ebiten.Image) error {

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		e.PressKey(4)
	} else {
		e.ReleaseKey(4)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		e.PressKey(5)
	} else {
		e.ReleaseKey(5)
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		e.PressKey(6)
	} else {
		e.ReleaseKey(6)
	}
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		e.PressKey(7)
	} else {
		e.ReleaseKey(7)
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		e.PressKey(1)
	} else {
		e.ReleaseKey(1)
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		e.PressKey(0)
	} else {
		e.ReleaseKey(0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		e.PressKey(2)
	} else {
		e.ReleaseKey(2)
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		e.PressKey(3)
	} else {
		e.ReleaseKey(3)
	}

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
	if e.EnableDebug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %f", ebiten.CurrentFPS()))
	}
	return nil
}

func main() {
	gameArg := os.Args[1]
	_, okDebug := os.LookupEnv("DEBUG")
	_, okEnableTestPanics := os.LookupEnv("ENABLE_TEST_PANICS")
	maxCyclesEnv, okMaxCycles := os.LookupEnv("MAX_CYCLES")
	maxCycles := 0
	if okMaxCycles {
		maxCyclesInt, err := strconv.Atoi(maxCyclesEnv)
		if err != nil {
			panic(err)
		}
		maxCycles = maxCyclesInt
	}
	e = emulator.NewEmulator(okDebug, okEnableTestPanics, maxCycles)
	e.LoadCartridge(gameArg)
	if err := ebiten.Run(update, width, height, scale, fmt.Sprintf("ナナ - %s", gameArg)); err != nil {
		panic(err)
	}
}
