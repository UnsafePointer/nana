// package name: nanavm
package main

import "C"
import "github.com/Ruenzuo/nana/emulator"

//export StartEmulation
func StartEmulation(gameArg string, enableDebug bool, enableLCDStateDebug bool, enableMemoryAccessDebug bool, enableTestPanics bool, maxCycles int) {
	emulator.StartEmulation(gameArg, enableDebug, enableLCDStateDebug, enableMemoryAccessDebug, enableTestPanics, maxCycles)
}

func main() {
}
