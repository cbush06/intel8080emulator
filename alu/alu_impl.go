package alu

import (
	"github.com/cbush06/intel8080emulator/memory"
)

// ALUImpl represents the collection of components that make up the Intel 8080's Arithmetic Logic Unit,
// specifically, it contains the ALUImpl's condition flags. Note that the Auxillary Carry flag is not
// implemented.
type ALUImpl struct {
	A *memory.Register
	ConditionFlags
}

// NewALU creates and returns a new ALUImpl struct
func NewALU(a *memory.Register) *ALUImpl {
	return &ALUImpl{
		ConditionFlags: &ConditionFlagsImpl{},
	}
}

// GetA returns a reference to the accumulator
func (alu *ALUImpl) GetA() *memory.Register {
	return alu.A
}

// SetA sets the ALU's reference to an accumulator
func (alu *ALUImpl) SetA(a *memory.Register) {
	alu.A = a
}

// UpdateFlags updates all ALU flags according to the value provided.
func (alu *ALUImpl) UpdateFlags(original uint8, new uint8) {
	alu.ClearFlags()
	alu.UpdateZero(new)
	alu.UpdateSign(new)
	alu.UpdateParity(new)
	alu.UpdateCarry(original, new)
}

// UpdateFlagsExceptCarry updates all ALU flags except the Carry flag according to the value provided.
func (alu *ALUImpl) UpdateFlagsExceptCarry(value uint8) {
	alu.ClearFlags()
	alu.UpdateZero(value)
	alu.UpdateSign(value)
	alu.UpdateParity(value)
}

// AddImmediate implements the ADI instruction. Specifically, the addend is added to the content of the
// accumulator. The result is placed in the accumulator. The AluFlags will be updated based on this
// operation's result.
func (alu *ALUImpl) AddImmediate(addend uint8) {
	var accum uint8
	alu.A.Read8(&accum)

	var result = addend + accum
	alu.UpdateFlags(accum, result)
	alu.A.Write8(uint8(result))
}

// AddImmediateWithCarry implements the ACI instruction. The addend and the content of the carry flag are
// added to the contents of the accumulator. The result is placed in the accumulator. The AluFlags
// will be updated based on this operation's result.
func (alu *ALUImpl) AddImmediateWithCarry(addend uint8) {
	var accum uint8
	alu.A.Read8(&accum)

	var carry uint8
	if alu.IsCarry() {
		carry = 1
	}

	var result = accum + carry + addend
	alu.UpdateFlags(accum, result)
	alu.A.Write8(result)
}

// DoubleAdd adds together two 16-bit words, updates the carry flag,
// and returns the result.
func (alu *ALUImpl) DoubleAdd(addend1 uint16, addend2 uint16) uint16 {
	sum := addend1 + addend2
	alu.UpdateCarryDoublePrecision(addend1, sum)
	return sum
}

// SubImmediate implements the SUI instruction. Specifically, the addend is subtracted from the content of the
// accumulator. The result is placed in the accumulator. The AluFlags will be updated based on this
// operation's result.
func (alu *ALUImpl) SubImmediate(subtrahend uint8) {
	var minuend uint8
	alu.A.Read8(&minuend)

	difference := minuend - subtrahend
	alu.UpdateFlags(minuend, difference)
	alu.A.Write8(difference)
}

// SubImmediateWithBorrow implements the SUI instruction. Specifically, the addend and the carry flag are both
// subtracted from the accumulator. The result is placed in the accumulator. The AluFlags will be updated based
// on this operation's result.
func (alu *ALUImpl) SubImmediateWithBorrow(subtrahend uint8) {
	var borrow uint8
	if alu.IsCarry() {
		borrow = 1
	}
	alu.SubImmediate(subtrahend + borrow)
}

// Increment increments a given value and updates all flags except the Carry flag, accordingly.
func (alu *ALUImpl) Increment(value uint8) uint8 {
	result := value + 1
	alu.UpdateFlagsExceptCarry(result)
	return result
}

// IncrementDouble increments a double-precision (16-bit) integer. No flags are updated.
func (alu *ALUImpl) IncrementDouble(value uint16) uint16 {
	return value + 1
}

// DecrementDouble decrements a double-precision (16-bit) integer. No flags are updated.
func (alu *ALUImpl) DecrementDouble(value uint16) uint16 {
	return value - 1
}

// Decrement decrements a given value and updates all flags except the Carry flag, accordingly.
func (alu *ALUImpl) Decrement(value uint8) uint8 {
	result := value - 1
	alu.UpdateFlagsExceptCarry(result)
	return result
}

// RotateRight rotates the 8-bit accumulator's value to the right by 1 bit such that (An) <- (An-1); (A7) <- (A0); (Cy) <- (A0). The high-order
// bit and Carry Flag are both set to the value of the low order bit.
func (alu *ALUImpl) RotateRight() {
	var accum uint8
	alu.A.Read8(&accum)

	bit0 := accum & 0x01
	if bit0 > 0 {
		alu.SetCarry()
	}

	alu.A.Write8((accum >> 1) | (bit0 << 7))
}

// AndAccumulator performs a bitwise AND operation on the contents of the accumulator and the operand.
// Flags Z, S, P, and AC are updated. The CY flag is cleared. The result is stored in the accumulator.
func (alu *ALUImpl) AndAccumulator(operand uint8) {
	var accum uint8
	var result uint8

	alu.A.Read8(&accum)
	result = accum & operand
	alu.A.Write8(result)
	alu.UpdateFlagsExceptCarry(result)
	alu.ClearCarry()
	alu.ClearAuxillaryCarry()
}

// XOrAccumulator performs a bitwise XOR operation on the contents of the accumulator and the operand.
// Flags Z, S, and P are updated. The CY and AC flags are cleared. The result is stored in the accumulator.
func (alu *ALUImpl) XOrAccumulator(operand uint8) {
	var accum uint8
	var result uint8

	alu.A.Read8(&accum)
	result = accum ^ operand
	alu.A.Write8(result)
	alu.UpdateFlagsExceptCarry(result)
	alu.ClearCarry()
	alu.ClearAuxillaryCarry()
}
