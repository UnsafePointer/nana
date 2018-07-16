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
	ticker := time.NewTicker(1000 * time.Millisecond)
	for range ticker.C {
		e.EmulateSecond()
	}
}
