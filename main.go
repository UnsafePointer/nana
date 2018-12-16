package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Ruenzuo/nana/emulator"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	width  = 160
	height = 144
	scale  = 5
)

var (
	e      *emulator.Emulator
	ticker *time.Ticker
	done   chan struct{}
)

func update(r *sdl.Renderer, t *sdl.Texture) error {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			ticker.Stop()
			close(done)
			break
		case *sdl.KeyboardEvent:
			event := event.(*sdl.KeyboardEvent)
			switch event.Keysym.Sym {
			case 'a':
				switch event.State {
				case sdl.PRESSED:
					e.PressKey(4)
					break
				case sdl.RELEASED:
					e.ReleaseKey(4)
					break
				}
				break
			case 's':
				switch event.State {
				case sdl.PRESSED:
					e.PressKey(5)
					break
				case sdl.RELEASED:
					e.ReleaseKey(5)
					break
				}
				break
			case sdl.K_UP:
				switch event.State {
				case sdl.PRESSED:
					e.PressKey(2)
					break
				case sdl.RELEASED:
					e.ReleaseKey(2)
					break
				}
				break
			case sdl.K_DOWN:
				switch event.State {
				case sdl.PRESSED:
					e.PressKey(3)
					break
				case sdl.RELEASED:
					e.ReleaseKey(3)
					break
				}
				break
			case sdl.K_LEFT:
				switch event.State {
				case sdl.PRESSED:
					e.PressKey(1)
					break
				case sdl.RELEASED:
					e.ReleaseKey(1)
					break
				}
				break
			case sdl.K_RIGHT:
				switch event.State {
				case sdl.PRESSED:
					e.PressKey(0)
					break
				case sdl.RELEASED:
					e.ReleaseKey(0)
					break
				}
				break
			case sdl.K_RETURN:
				switch event.State {
				case sdl.PRESSED:
					e.PressKey(7)
					break
				case sdl.RELEASED:
					e.ReleaseKey(7)
					break
				}
				break
			case sdl.K_SPACE:
				switch event.State {
				case sdl.PRESSED:
					e.PressKey(6)
					break
				case sdl.RELEASED:
					e.ReleaseKey(6)
					break
				}
				break
			}
		}
	}

	e.EmulateFrame()

	pixels, _, _ := t.Lock(nil)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			position := y*width + x
			pixels[(position*4 + 0)] = e.ScreenData[x][y][0]
			pixels[(position*4 + 1)] = e.ScreenData[x][y][1]
			pixels[(position*4 + 2)] = e.ScreenData[x][y][2]
			pixels[(position*4 + 3)] = 255
		}
	}
	t.Unlock()

	r.Clear()
	r.Copy(t, nil, nil)
	r.Present()
	return nil
}

func main() {
	gameArg := os.Args[1]
	_, okDebug := os.LookupEnv("NANA_DEBUG")
	_, okLCDState := os.LookupEnv("NANA_LCD_STATE_DEBUG")
	_, okMemoryAccess := os.LookupEnv("NANA_MEMORY_ACCESS_DEBUG")
	_, okEnableTestPanics := os.LookupEnv("NANA_ENABLE_TEST_PANICS")
	maxCyclesEnv, okMaxCycles := os.LookupEnv("NANA_MAX_CYCLES")
	maxCycles := 0
	if okMaxCycles {
		maxCyclesInt, err := strconv.Atoi(maxCyclesEnv)
		if err != nil {
			panic(err)
		}
		maxCycles = maxCyclesInt
	}
	e = emulator.NewEmulator(okDebug, okLCDState, okMemoryAccess, okEnableTestPanics, maxCycles)
	e.LoadCartridge(gameArg)

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(fmt.Sprintf("ナナ - %s", gameArg), sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width*scale, height*scale, sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	renderer.SetScale(scale, scale)

	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, width, height)
	if err != nil {
		panic(err)
	}
	defer texture.Destroy()

	spec := &sdl.AudioSpec{
		Freq:     emulator.SampleRate,
		Format:   sdl.AUDIO_U8,
		Channels: 2,
		Samples:  emulator.SoundBufferSize,
	}
	if err := sdl.OpenAudio(spec, nil); err != nil {
		panic(err)
	}
	sdl.PauseAudio(false)

	done = make(chan struct{})
	ticker = time.NewTicker(1000 * time.Millisecond / 60)

loop:
	for {
		select {
		case <-ticker.C:
			update(renderer, texture)
		case <-done:
			break loop
		}
	}

	sdl.CloseAudio()
}
