package emulator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/Ruenzuo/nana/emulator"
)

var _ = Describe("Emulator", func() {
	var (
		emulator Emulator
	)

	BeforeEach(func() {
		emulator = *NewEmulator(false)
	})

	Describe("verifying push/pop 16-Bit values into stack", func() {
		Context("when pushing a new value", func() {
			BeforeEach(func() {
				emulator.PushIntoStack(0xEEAA)
				emulator.PushIntoStack(0xFFBB)
				emulator.PushIntoStack(0xDDCC)
			})

			It("should pop the same value", func() {
				Expect(emulator.PopFromStack()).To(Equal(uint16(0xDDCC)))
				Expect(emulator.PopFromStack()).To(Equal(uint16(0xFFBB)))
				Expect(emulator.PopFromStack()).To(Equal(uint16(0xEEAA)))
			})
		})
	})
})
