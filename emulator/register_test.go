package emulator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/Ruenzuo/nana/emulator"
)

var _ = Describe("Register", func() {
	var (
		register Register
	)

	BeforeEach(func() {
		register = Register{}
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