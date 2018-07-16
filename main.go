package main

import (
	"os"

	"github.com/Ruenzuo/nana/emulator"
)

func main() {
	gameArg := os.Args[1]
	_, okDebug := os.LookupEnv("DEBUG")
	e := emulator.NewEmulator(okDebug)
	e.LoadCartridge(gameArg)
	for true {
		e.EmulateSecond()
	}
}
