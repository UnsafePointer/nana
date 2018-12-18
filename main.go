package main

import (
	"os"
	"strconv"

	"github.com/Ruenzuo/nana/emulator"
)

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

	emulator.StartEmulation(gameArg, okDebug, okLCDState, okMemoryAccess, okEnableTestPanics, maxCycles)
}
