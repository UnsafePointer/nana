package emulator

import "fmt"

func testPanic(shouldPanic bool, message string) {
	if shouldPanic {
		panic(message)
	}
}

func (e *Emulator) LogMessage(message string) {
	if e.EnableDebug {
		fmt.Println(message)
	}
}
