package emulator

func testPanic(shouldPanic bool, message string) {
	if shouldPanic {
		panic(message)
	}
}
