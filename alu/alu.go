package alu

import (
	"github.com/cbush06/intel8080emulator/memory"
)

// ALU represents the collection of components that make up the Intel 8080's Arithmetic Logic Unit,
// specifically, it contains the ALU's condition flags.
type ALU struct {
	A memory.Register
	ConditionFlags
}

// UpdateFlags updates all ALU flags according to the value provided.
func (alu *ALU) UpdateFlags(value uint16) {
	alu.ClearFlags()
	alu.UpdateZero(value)
	alu.UpdateSign(value)
	alu.UpdateParity(uint8(value))
	alu.UpdateCarry(value)
}

// UpdateFlagsExceptCarry updates all ALU flags except the Carry flag according to the value provided.
func (alu *ALU) UpdateFlagsExceptCarry(value uint16) {
	alu.ClearFlags()
	alu.UpdateZero(value)
	alu.UpdateSign(value)
	alu.UpdateParity(uint8(value))
}

// AddImmediate implements the ADI instruction. Specifically, the addend is added to the content of the
// accumulator. The result is placed in the accumulator. The AluFlags will be updated based on this
// operation's result.
func (alu *ALU) AddImmediate(addend uint8) {
	var accum uint8
	alu.A.Read8(&accum)

	var result = uint16(addend) + uint16(accum)
	alu.UpdateFlags(result)
	alu.A.Write8(uint8(result))
}

// AddImmediateWithCarry implements the ACI instruction. The addend and the content of the carry flag are
// added to the contents of the accumulator. The result is placed in the accumulator. The AluFlags
// will be updated based on this operation's result.
func (alu *ALU) AddImmediateWithCarry(addend uint8) {
	var accum uint8
	alu.A.Read8(&accum)

	var carry uint8
	if alu.Carry {
		carry = 1
	}

	var result = uint16(accum + carry + addend)
	alu.UpdateFlags(result)
	alu.A.Write8(uint8(result))
}

// SubImmediate implements the SUI instruction. Specifically, the addend is subtracted from the content of the
// accumulator. The result is placed in the accumulator. The AluFlags will be updated based on this
// operation's result.
func (alu *ALU) SubImmediate(addend uint8) {
	// Negate the input value
	var input = uint8(-1 * int8(addend))
	alu.AddImmediate(input)
}

// SubImmediateWithBorrow implements the SUI instruction. Specifically, the addend and the carry flag are both
// subtracted from the accumulator. The result is placed in the accumulator. The AluFlags will be updated based
// on this operation's result.
func (alu *ALU) SubImmediateWithBorrow(addend uint8) {
	var borrow uint8
	if alu.Carry {
		borrow = 1
	}
	alu.SubImmediate(addend + borrow)
}

// Increment increments a given value and updates all flags except the Carry flag, accordingly.
func (alu *ALU) Increment(addend uint8) uint8 {
	result := uint16(addend + 1)
	alu.UpdateFlagsExceptCarry(result)
	return uint8(result)
}

// Decrement decrements a given value and updates all flags except the Carry flag, accordingly.
func (alu *ALU) Decrement(addend uint8) uint8 {
	result := uint16(addend - 1)
	alu.UpdateFlagsExceptCarry(result)
	return uint8(result)
}
