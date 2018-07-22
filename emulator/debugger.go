package emulator

import (
	"fmt"
	"os"
)

func testPanic(shouldPanic bool, message string) {
	if shouldPanic {
		panic(message)
	}
}

func (e *Emulator) SetupLogFile() {
	if _, err := os.Stat("nana.log"); err == nil {
		os.Remove("nana.log")
	}
	f, err := os.OpenFile("nana.log", os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

func (e *Emulator) LogMessage(message string) {
	if e.EnableDebug {
		e.LogBuffer.WriteString(fmt.Sprintf("%s\n", message))
	}
	if e.LogBuffer.Len() > 8192 {
		f, err := os.OpenFile("nana.log", os.O_WRONLY|os.O_APPEND, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if _, err := f.WriteString(e.LogBuffer.String()); err != nil {
			panic(err)
		}
		e.LogBuffer.Reset()
	}
}
