package emulator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/Ruenzuo/nana/emulator"
)

var _ = Describe("Register8Bit", func() {
	var (
		register Register8Bit
	)

	BeforeEach(func() {
		register = Register8Bit{}
	})

	Describe("setting value", func() {
		Context("when setting value", func() {
			BeforeEach(func() {
				register.SetValue(0xFF)
			})

			It("should contain the right value", func() {
				Expect(register.Value()).To(Equal(uint8(0xFF)))
			})
		})
	})
})

var _ = Describe("Register16Bit", func() {
	var (
		register Register16Bit
	)

	BeforeEach(func() {
		register = Register16Bit{}
	})

	Describe("setting value", func() {
		Context("when setting value", func() {
			BeforeEach(func() {
				register.SetValue(0xFFEE)
			})

			It("should contain the right value", func() {
				Expect(register.Value()).To(Equal(uint16(0xFFEE)))
			})

			Context("when incrementing the value", func() {
				BeforeEach(func() {
					register.Increment()
				})

				It("should contain the right value", func() {
					Expect(register.Value()).To(Equal(uint16(0xFFEF)))
				})

				Context("when decrementing the value", func() {
					BeforeEach(func() {
						register.Decrement()
					})

					It("should contain the right value", func() {
						Expect(register.Value()).To(Equal(uint16(0xFFEE)))
					})
				})
			})
		})
	})

	Describe("setting bits", func() {
		Context("when setting low bits", func() {
			BeforeEach(func() {
				register.SetLow(0xFF)
			})

			It("should contain low bits value", func() {
				Expect(register.Low()).To(Equal(uint8(0xFF)))
			})

			Context("when setting high bits", func() {
				BeforeEach(func() {
					register.SetHigh(0xEE)
				})

				It("should contain high bits value", func() {
					Expect(register.High()).To(Equal(uint8(0xEE)))
				})

				Context("when getting value", func() {
					It("should append values correctly", func() {
						Expect(register.Value()).To(Equal(uint16(0xEEFF)))
					})
				})
			})
		})
	})
})
