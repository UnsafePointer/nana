package main

import (
	"os"

	"github.com/Ruenzuo/nana/emulator"
)

func main() {
	gameArg := os.Args[1]
	e := emulator.NewEmulator()
	e.LoadCartridge(gameArg)
}
