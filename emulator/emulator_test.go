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

	Describe("verifying the initial values", func() {
		Context("when using new instance", func() {
			It("should contain the right initial values", func() {
				// Registers
				Expect(emulator.AF.Value()).To(Equal(uint16(0x01B0)))
				Expect(emulator.BC.Value()).To(Equal(uint16(0x0013)))
				Expect(emulator.DE.Value()).To(Equal(uint16(0x00D8)))
				Expect(emulator.HL.Value()).To(Equal(uint16(0x014D)))
				// SP & PC
				Expect(emulator.StackPointer.Value()).To(Equal(uint16(0xFFFE)))
				Expect(emulator.ProgramCounter.Value()).To(Equal(uint16(0x100)))
				// ROM (only one address)
				Expect(emulator.ROM[0xFF19]).To(Equal(uint8(0xBF)))
				// ROM/RAM bank counter
				Expect(emulator.CurrentROMBank).To(Equal(uint16(1)))
				Expect(emulator.CurrentRAMBank).To(Equal(uint16(0)))
			})
		})
	})
})
