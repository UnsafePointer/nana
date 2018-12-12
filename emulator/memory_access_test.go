package emulator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/Ruenzuo/nana/emulator"
)

var _ = Describe("Emulator", func() {
	var (
		emulator    Emulator
		memoryValue uint8
	)

	BeforeEach(func() {
		emulator = *NewEmulator(false, false, false, false, 0)
	})

	Describe("verifying memory access", func() {
		Context("when writting to ROM", func() {
			BeforeEach(func() {
				memoryValue = emulator.ReadMemory8Bit(0x0001)
				emulator.WriteMemory(0x0001, 0xFF)
			})

			It("should not change the memory at that address", func() {
				Expect(emulator.ROM[0x0001]).To(Equal(memoryValue))
			})
		})

		Context("when writting to restricted section", func() {
			BeforeEach(func() {
				memoryValue = emulator.ReadMemory8Bit(0xFEA0)
				emulator.WriteMemory(0xFEA0, 0xFF)
			})

			It("should not change the memory at that address", func() {
				Expect(emulator.ROM[0xFEA0]).To(Equal(memoryValue))
			})
		})

		Context("when writting to echo", func() {
			BeforeEach(func() {
				emulator.WriteMemory(0xE000, 0xFF)
			})

			It("should write to echo address", func() {
				Expect(emulator.ROM[0xE000]).To(Equal(uint8(0xFF)))
				Expect(emulator.ROM[0xC000]).To(Equal(uint8(0xFF)))
			})
		})

		Context("when writing to unrestricted section", func() {
			BeforeEach(func() {
				emulator.WriteMemory(0xFFFF, 0xFF)
			})

			It("should write to address", func() {
				Expect(emulator.ROM[0xFFFF]).To(Equal(uint8(0xFF)))
			})

			Context("when accessing 16 bits", func() {
				BeforeEach(func() {
					emulator.WriteMemory(0xFFFE, 0xFE)
				})

				It("should retrieve the write value", func() {
					Expect(emulator.ReadMemory16Bit(0xFFFE)).To(Equal(uint16(0xFFFE)))
				})
			})
		})
	})
})
