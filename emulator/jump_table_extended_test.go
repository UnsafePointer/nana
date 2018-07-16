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
		emulator = *NewEmulator()
	})

	Describe("verifying operation codes work", func() {
		Context("when using SWAP", func() {
			BeforeEach(func() {
				emulator.BC.SetHigh(0xF0)
				emulator.CPU8BitRegisterSwap(&emulator.BC.High)
			})

			It("should SWAP the value", func() {
				Expect(emulator.BC.High.Value()).To(Equal(uint8(0x0F)))
			})
		})
		Context("when using BIT", func() {
			BeforeEach(func() {
				emulator.BC.SetLow(0xF0)
				emulator.ClearAllFlags()
				emulator.CPU8BitRegisterBit(&emulator.BC.Low, 3)
			})

			It("should set the right flags", func() {
				Expect(emulator.FlagZ()).To(Equal(true))
				Expect(emulator.FlagN()).To(Equal(false))
				Expect(emulator.FlagH()).To(Equal(true))
				Expect(emulator.FlagC()).To(Equal(false))
			})
		})
	})
})
