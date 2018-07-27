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
		emulator = *NewEmulator(false, 0)
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
		Context("when using SET", func() {
			BeforeEach(func() {
				emulator.BC.SetLow(0x00)
				emulator.ClearAllFlags()
				emulator.CPU8BitRegisterSet(&emulator.BC.Low, 0)
			})

			It("should set the right bits", func() {
				Expect(emulator.BC.Low.Value()).To(Equal(uint8(0x01)))
			})
		})
		Context("when using RES", func() {
			BeforeEach(func() {
				emulator.BC.SetLow(0xFF)
				emulator.ClearAllFlags()
				emulator.CPU8BitRegisterReset(&emulator.BC.Low, 0)
			})

			It("should reset the right bits", func() {
				Expect(emulator.BC.Low.Value()).To(Equal(uint8(0xFE)))
			})
		})
	})
})
