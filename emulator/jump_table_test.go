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

	Describe("verifying operation codes work", func() {
		Context("when using 8-Bit ADD", func() {
			BeforeEach(func() {
				emulator.AF.SetLow(0x01)
				emulator.AF.SetHigh(0x02)
				emulator.CPU8BitRegisterAdd(&emulator.AF.High, emulator.AF.Low.Value(), false)
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
				emulator.CPU8BitRegisterAdd(&emulator.AF.High, emulator.AF.Low.Value(), false)
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
				emulator.CPU8BitRegisterSubtract(&emulator.AF.High, emulator.AF.Low.Value(), false)
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
				emulator.CPU8BitRegisterSubtract(&emulator.AF.High, emulator.AF.Low.Value(), false)
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
		Context("when using 8-Bit XOR", func() {
			BeforeEach(func() {
				emulator.AF.SetHigh(0xF2)
				emulator.BC.SetHigh(0xD3)
				emulator.CPU8BitXor(emulator.BC.High.Value())
			})

			It("should XOR the values and set the right flags", func() {
				Expect(emulator.AF.High.Value()).To(Equal(uint8(0x21)))
				Expect(emulator.FlagC()).To(Equal(false))
				Expect(emulator.FlagH()).To(Equal(false))
				Expect(emulator.FlagN()).To(Equal(false))
				Expect(emulator.FlagZ()).To(Equal(false))
			})
		})
		Context("when using 8-Bit CP", func() {
			BeforeEach(func() {
				emulator.AF.SetHigh(0xF1)
				emulator.BC.SetHigh(0x1F)
				emulator.CPU8BitCompare(emulator.BC.High.Value())
			})

			It("should CP the values and set the right flags", func() {
				Expect(emulator.AF.High.Value()).To(Equal(uint8(0xF1)))
				Expect(emulator.FlagC()).To(Equal(false))
				Expect(emulator.FlagH()).To(Equal(true))
				Expect(emulator.FlagN()).To(Equal(true))
				Expect(emulator.FlagZ()).To(Equal(false))
			})
		})
		Context("when using 8-Bit INC", func() {
			BeforeEach(func() {
				emulator.AF.SetHigh(0x0F)
				emulator.SetFlagC()
				emulator.CPU8BitRegisterIncrement(&emulator.AF.High)
			})

			It("should INC the values and set the right flags", func() {
				Expect(emulator.AF.High.Value()).To(Equal(uint8(0x10)))
				Expect(emulator.FlagC()).To(Equal(true))
				Expect(emulator.FlagH()).To(Equal(true))
				Expect(emulator.FlagN()).To(Equal(false))
				Expect(emulator.FlagZ()).To(Equal(false))
			})
		})
		Context("when using 8-Bit DEC", func() {
			BeforeEach(func() {
				emulator.AF.SetHigh(0x10)
				emulator.SetFlagC()
				emulator.CPU8BitRegisterDecrement(&emulator.AF.High)
			})

			It("should DEC the values and set the right flags", func() {
				Expect(emulator.AF.High.Value()).To(Equal(uint8(0x0F)))
				Expect(emulator.FlagC()).To(Equal(true))
				Expect(emulator.FlagH()).To(Equal(true))
				Expect(emulator.FlagN()).To(Equal(true))
				Expect(emulator.FlagZ()).To(Equal(false))
			})
		})
		Context("when using RLC", func() {
			BeforeEach(func() {
				emulator.AF.High.SetValue(0xFF)
				emulator.CPU8BitRegisterRLC(&emulator.AF.High)
			})

			It("should rotate the value and set the right flags", func() {
				Expect(emulator.AF.High.Value()).To(Equal(uint8(0xFF)))
				Expect(emulator.FlagC()).To(Equal(true))
				Expect(emulator.FlagH()).To(Equal(false))
				Expect(emulator.FlagN()).To(Equal(false))
				Expect(emulator.FlagZ()).To(Equal(false))
			})
		})
		Context("when using RL", func() {
			BeforeEach(func() {
				emulator.AF.High.SetValue(0xFF)
				emulator.ClearFlagC()
				emulator.CPU8BitRegisterRL(&emulator.AF.High)
			})

			It("should rotate the value and set the right flags", func() {
				Expect(emulator.AF.High.Value()).To(Equal(uint8(0xFE)))
				Expect(emulator.FlagC()).To(Equal(true))
				Expect(emulator.FlagH()).To(Equal(false))
				Expect(emulator.FlagN()).To(Equal(false))
				Expect(emulator.FlagZ()).To(Equal(false))
			})
		})
		Context("when using RRC", func() {
			BeforeEach(func() {
				emulator.AF.High.SetValue(0x0F)
				emulator.CPU8BitRegisterRRC(&emulator.AF.High)
			})

			It("should rotate the value and set the right flags", func() {
				Expect(emulator.AF.High.Value()).To(Equal(uint8(0x87)))
				Expect(emulator.FlagC()).To(Equal(true))
				Expect(emulator.FlagH()).To(Equal(false))
				Expect(emulator.FlagN()).To(Equal(false))
				Expect(emulator.FlagZ()).To(Equal(false))
			})
		})
		Context("when using RR", func() {
			BeforeEach(func() {
				emulator.AF.High.SetValue(0x0F)
				emulator.ClearFlagC()
				emulator.CPU8BitRegisterRR(&emulator.AF.High)
			})

			It("should rotate the value and set the right flags", func() {
				Expect(emulator.AF.High.Value()).To(Equal(uint8(0x07)))
				Expect(emulator.FlagC()).To(Equal(true))
				Expect(emulator.FlagH()).To(Equal(false))
				Expect(emulator.FlagN()).To(Equal(false))
				Expect(emulator.FlagZ()).To(Equal(false))
			})
		})
		Context("when using SLA", func() {
			BeforeEach(func() {
				emulator.AF.High.SetValue(0xFF)
				emulator.CPU8BitRegisterSLA(&emulator.AF.High)
			})

			It("should shift the value and set the right flags", func() {
				Expect(emulator.AF.High.Value()).To(Equal(uint8(0xFE)))
				Expect(emulator.FlagC()).To(Equal(true))
				Expect(emulator.FlagH()).To(Equal(false))
				Expect(emulator.FlagN()).To(Equal(false))
				Expect(emulator.FlagZ()).To(Equal(false))
			})
		})
		Context("when using SRL", func() {
			BeforeEach(func() {
				emulator.AF.High.SetValue(0xFF)
				emulator.CPU8BitRegisterSRL(&emulator.AF.High)
			})

			It("should shift the value and set the right flags", func() {
				Expect(emulator.AF.High.Value()).To(Equal(uint8(0x7F)))
				Expect(emulator.FlagC()).To(Equal(true))
				Expect(emulator.FlagH()).To(Equal(false))
				Expect(emulator.FlagN()).To(Equal(false))
				Expect(emulator.FlagZ()).To(Equal(false))
			})
		})
		Context("when using 16-Bit ADD", func() {
			BeforeEach(func() {
				emulator.HL.SetValue(0xFFFF)
				emulator.BC.SetValue(0x0100)
				emulator.SetFlagZ()
				emulator.CPU16BitRegisterAdd(&emulator.HL, emulator.BC)
			})

			It("should ADD the values and set the right flags", func() {
				Expect(emulator.HL.Value()).To(Equal(uint16(0x00FF)))
				Expect(emulator.FlagC()).To(Equal(true))
				Expect(emulator.FlagH()).To(Equal(true))
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
