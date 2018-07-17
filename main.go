package main

import (
	"os"
	"time"

	"github.com/Ruenzuo/nana/emulator"
)

func main() {
	gameArg := os.Args[1]
	_, okDebug := os.LookupEnv("DEBUG")
	e := emulator.NewEmulator(okDebug)
	e.LoadCartridge(gameArg)
	ticker := time.NewTicker(time.Second / 60)
	go func() {
		for range ticker.C {
			e.EmulateFrame()
		}
	}()
	time.Sleep(time.Millisecond * 10000)
	ticker.Stop()
}
