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
		emulator = Emulator{}
		emulator.AF.SetLow(0x00)
	})

	Describe("setting flags", func() {
		Context("when setting Z", func() {
			BeforeEach(func() {
				emulator.SetFlagZ()
			})

			It("should enable the proper flag", func() {
				Expect(emulator.FlagZ()).To(Equal(true))
				Expect(emulator.FlagN()).To(Equal(false))
				Expect(emulator.FlagH()).To(Equal(false))
				Expect(emulator.FlagC()).To(Equal(false))
			})

			Context("when clearing Z", func() {
				BeforeEach(func() {
					emulator.ClearFlagZ()
				})

				It("should clear the proper flag", func() {
					Expect(emulator.FlagZ()).To(Equal(false))
					Expect(emulator.FlagN()).To(Equal(false))
					Expect(emulator.FlagH()).To(Equal(false))
					Expect(emulator.FlagC()).To(Equal(false))
				})
			})
		})

		Context("when setting N", func() {
			BeforeEach(func() {
				emulator.SetFlagN()
			})

			It("should enable the proper flag", func() {
				Expect(emulator.FlagZ()).To(Equal(false))
				Expect(emulator.FlagN()).To(Equal(true))
				Expect(emulator.FlagH()).To(Equal(false))
				Expect(emulator.FlagC()).To(Equal(false))
			})

			Context("when clearing N", func() {
				BeforeEach(func() {
					emulator.ClearFlagN()
				})

				It("should clear the proper flag", func() {
					Expect(emulator.FlagZ()).To(Equal(false))
					Expect(emulator.FlagN()).To(Equal(false))
					Expect(emulator.FlagH()).To(Equal(false))
					Expect(emulator.FlagC()).To(Equal(false))
				})
			})
		})

		Context("when setting H", func() {
			BeforeEach(func() {
				emulator.SetFlagH()
			})

			It("should enable the proper flag", func() {
				Expect(emulator.FlagZ()).To(Equal(false))
				Expect(emulator.FlagN()).To(Equal(false))
				Expect(emulator.FlagH()).To(Equal(true))
				Expect(emulator.FlagC()).To(Equal(false))
			})

			Context("when clearing H", func() {
				BeforeEach(func() {
					emulator.ClearFlagH()
				})

				It("should clear the proper flag", func() {
					Expect(emulator.FlagZ()).To(Equal(false))
					Expect(emulator.FlagN()).To(Equal(false))
					Expect(emulator.FlagH()).To(Equal(false))
					Expect(emulator.FlagC()).To(Equal(false))
				})
			})
		})

		Context("when setting C", func() {
			BeforeEach(func() {
				emulator.SetFlagC()
			})

			It("should enable the proper flag", func() {
				Expect(emulator.FlagZ()).To(Equal(false))
				Expect(emulator.FlagN()).To(Equal(false))
				Expect(emulator.FlagH()).To(Equal(false))
				Expect(emulator.FlagC()).To(Equal(true))
			})

			Context("when clearing C", func() {
				BeforeEach(func() {
					emulator.ClearFlagC()
				})

				It("should clear the proper flag", func() {
					Expect(emulator.FlagZ()).To(Equal(false))
					Expect(emulator.FlagN()).To(Equal(false))
					Expect(emulator.FlagH()).To(Equal(false))
					Expect(emulator.FlagC()).To(Equal(false))
				})
			})
		})
	})
})
