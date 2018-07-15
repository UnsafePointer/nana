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
		Context("when using 8-Bit ADD", func() {
			BeforeEach(func() {
				emulator.AF.SetLow(0x01)
				emulator.AF.SetHigh(0x02)
				emulator.CPU8BitAdd(&emulator.AF.High, emulator.AF.Low.Value(), false)
			})

			It("should add the values and set the right flags", func() {
				Expect(emulator.AF.High.Value()).To(Equal(uint8(0x03)))
				Expect(emulator.FlagC()).To(Equal(false))
				Expect(emulator.FlagH()).To(Equal(false))
				Expect(emulator.FlagN()).To(Equal(false))
				Expect(emulator.FlagZ()).To(Equal(false))
			})
		})
		Context("when using 8-Bit ADD", func() {
			BeforeEach(func() {
				emulator.AF.SetLow(0xFF)
				emulator.AF.SetHigh(0x01)
				emulator.CPU8BitAdd(&emulator.AF.High, emulator.AF.Low.Value(), false)
			})

			It("should add the values and set the right flags", func() {
				Expect(emulator.AF.High.Value()).To(Equal(uint8(0x00)))
				Expect(emulator.FlagC()).To(Equal(true))
				Expect(emulator.FlagH()).To(Equal(true))
				Expect(emulator.FlagN()).To(Equal(false))
				Expect(emulator.FlagZ()).To(Equal(true))
			})
		})
		Context("when using 8-Bit SUB", func() {
			BeforeEach(func() {
				emulator.AF.SetLow(0x01)
				emulator.AF.SetHigh(0x02)
				emulator.CPU8BitSub(&emulator.AF.High, emulator.AF.Low.Value(), false)
			})

			It("should add the values and set the right flags", func() {
				Expect(emulator.AF.High.Value()).To(Equal(uint8(0x01)))
				Expect(emulator.FlagC()).To(Equal(false))
				Expect(emulator.FlagH()).To(Equal(false))
				Expect(emulator.FlagN()).To(Equal(true))
				Expect(emulator.FlagZ()).To(Equal(false))
			})
		})
		Context("when using 8-Bit SUB", func() {
			BeforeEach(func() {
				emulator.AF.SetLow(0xFF)
				emulator.AF.SetHigh(0x01)
				emulator.CPU8BitSub(&emulator.AF.High, emulator.AF.Low.Value(), false)
			})

			It("should add the values and set the right flags", func() {
				Expect(emulator.AF.High.Value()).To(Equal(uint8(0x02)))
				Expect(emulator.FlagC()).To(Equal(true))
				Expect(emulator.FlagH()).To(Equal(true))
				Expect(emulator.FlagN()).To(Equal(true))
				Expect(emulator.FlagZ()).To(Equal(false))
			})
		})
		Context("when using 8-Bit AND", func() {
			BeforeEach(func() {
				emulator.AF.SetHigh(0xF5)
				emulator.BC.SetHigh(0x0)
				emulator.CPU8BitAnd(emulator.BC.High.Value())
			})

			It("should AND the values and set the right flags", func() {
				Expect(emulator.AF.High.Value()).To(Equal(uint8(0x0)))
				Expect(emulator.FlagC()).To(Equal(false))
				Expect(emulator.FlagH()).To(Equal(true))
				Expect(emulator.FlagN()).To(Equal(false))
				Expect(emulator.FlagZ()).To(Equal(true))
			})
		})
		Context("when using 8-Bit OR", func() {
			BeforeEach(func() {
				emulator.AF.SetHigh(0x0)
				emulator.BC.SetHigh(0x0)
				emulator.CPU8BitOr(emulator.BC.High.Value())
			})

			It("should OR the values and set the right flags", func() {
				Expect(emulator.AF.High.Value()).To(Equal(uint8(0x0)))
				Expect(emulator.FlagC()).To(Equal(false))
				Expect(emulator.FlagH()).To(Equal(false))
				Expect(emulator.FlagN()).To(Equal(false))
				Expect(emulator.FlagZ()).To(Equal(true))
			})
		})
		Context("when values overflow", func() {
			BeforeEach(func() {
				emulator.ProgramCounter.SetValue(0x0)
				emulator.StackPointer.SetValue(0xFFFF)
				emulator.ROM[0] = 0x0F
				emulator.ExecuteOpCode(0xF8)
			})

			It("should set the right flags", func() {
				Expect(emulator.FlagC()).To(Equal(true))
				Expect(emulator.FlagH()).To(Equal(true))
				Expect(emulator.FlagN()).To(Equal(false))
				Expect(emulator.FlagZ()).To(Equal(false))
			})
		})

		Context("when values don't overflow", func() {
			BeforeEach(func() {
				emulator.ProgramCounter.SetValue(0x0)
				emulator.StackPointer.SetValue(0xFFF0)
				emulator.ROM[0] = 0x01
				emulator.ExecuteOpCode(0xF8)
			})

			It("should set the right flags", func() {
				Expect(emulator.FlagC()).To(Equal(false))
				Expect(emulator.FlagH()).To(Equal(false))
				Expect(emulator.FlagN()).To(Equal(false))
				Expect(emulator.FlagZ()).To(Equal(false))
			})
		})
	})
})
