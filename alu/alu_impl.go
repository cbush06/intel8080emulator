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
		A:              a,
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
func (alu *ALUImpl) UpdateFlags(original uint8, result uint16) {
	var resultMasked = uint8(result & 0xFF)

	alu.ClearFlags()
	alu.UpdateZero(resultMasked)
	alu.UpdateSign(resultMasked)
	alu.UpdateParity(resultMasked)
	alu.UpdateCarry(result)
	alu.UpdateAuxiliaryCarry(original, resultMasked)
}

// UpdateFlagsExceptCarry updates all ALU flags except the Carry flag according to the value provided.
func (alu *ALUImpl) UpdateFlagsExceptCarry(value uint8) {
	alu.ClearFlags()
	alu.UpdateZero(value)
	alu.UpdateSign(value)
	alu.UpdateParity(value)
}

// AddImmediate adds the addend added to the content of the accumulator. The result is placed in the accumulator. The
// AluFlags will be updated based on this operation's result.
func (alu *ALUImpl) AddImmediate(addend uint8) {
	var accum uint8
	alu.A.Read8(&accum)

	result := uint16(addend) + uint16(accum)
	alu.UpdateFlags(accum, result)
	alu.A.Write8(uint8(result & 0xFF))
}

// AddImmediateWithCarry adds the addend and the content of the carry flag to the contents of the accumulator. The
// result is placed in the accumulator. The AluFlags will be updated based on this operation's result.
func (alu *ALUImpl) AddImmediateWithCarry(addend uint8) {
	var accum uint8
	alu.A.Read8(&accum)

	var carry uint8
	if alu.IsCarry() {
		carry = 1
	}

	result := uint16(accum) + uint16(carry) + uint16(addend)
	alu.UpdateFlags(accum, result)
	alu.A.Write8(uint8(result & 0xFF))
}

// DoubleAdd adds together two 16-bit words, updates the carry flag,
// and returns the result.
func (alu *ALUImpl) DoubleAdd(addend1 uint16, addend2 uint16) uint16 {
	sum := uint32(addend1) + uint32(addend2)
	alu.UpdateCarryDoublePrecision(sum)
	return uint16(sum & 0xFFFF)
}

// SubImmediate subtracts the subtrahend from the content of the accumulator. The result is placed in the accumulator.
// The AluFlags will be updated based on this operation's result.
func (alu *ALUImpl) SubImmediate(subtrahend uint8) {
	var minuend uint8
	alu.A.Read8(&minuend)

	difference := uint16(minuend) - uint16(subtrahend)
	alu.UpdateFlags(minuend, difference)
	alu.A.Write8(uint8(difference & 0xFF))
}

// SubImmediateWithBorrow subtracts the subtrahend and the carry flag from the accumulator. The result is placed in
// the accumulator. The AluFlags will be updated based on this operation's result.
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
	alu.UpdateAuxiliaryCarry(value, result)
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
	alu.UpdateAuxiliaryCarry(value, result)
	return result
}

// RotateRight rotates the 8-bit accumulator's value to the right by 1 bit such that (An) -> (An-1); (A7) <- (A0); (Cy) <- (A0). The high-order
// bit and Carry Flag are both set to the value of the low-order bit.
func (alu *ALUImpl) RotateRight() {
	var accum uint8
	alu.A.Read8(&accum)

	bit0 := accum & 0x01
	if bit0 > 0 {
		alu.SetCarry()
	} else {
		alu.ClearCarry()
	}

	alu.A.Write8((accum >> 1) | (bit0 << 7))
}

// RotateRightThroughCarry rotates the 8-bit accumulators value to the right by 1 bit such that (An) <- (An+1); (CY) <- (A0); (A7) <- (CY). The
// high-order bit is set to the Carry Flag and the Carry Flag is set to the value of the low-order bit.
func (alu *ALUImpl) RotateRightThroughCarry() {
	var accum uint8
	alu.A.Read8(&accum)

	bit0 := accum & 0x01
	accum = accum >> 1

	if alu.IsCarry() {
		accum |= 0x80
	}

	if bit0 > 0 {
		alu.SetCarry()
	} else {
		alu.ClearCarry()
	}

	alu.A.Write8(accum)
}

// RotateLeft rotates the 8-bit accumulator's value to the left by 1 bit such that (An+1) <- (An); (A7) -> (A0); (CY) <- (A7). The low-order
// bit and Carry Flag are both set to the value of the high-order bit.
func (alu *ALUImpl) RotateLeft() {
	var accum uint8
	alu.A.Read8(&accum)

	bit7 := accum & 0x80
	if bit7 > 0 {
		alu.SetCarry()
	} else {
		alu.ClearCarry()
	}

	alu.A.Write8((accum << 1) | (bit7 >> 7))
}

// RotateLeftThroughCarry rotates the 8-bit accumulator's value to the left by 1 bit such that (An+1) <- (An); (CY) <- (A7); (A0) <- (CY). The low-order
// bit is set to the Carry Flag and the Carry Flag is set to the high-order bit.
func (alu *ALUImpl) RotateLeftThroughCarry() {
	var accum uint8
	alu.A.Read8(&accum)

	bit7 := accum & 0x80
	accum = accum << 1

	if alu.IsCarry() {
		accum |= 0x01
	}

	if bit7 > 0 {
		alu.SetCarry()
	} else {
		alu.ClearCarry()
	}

	alu.A.Write8(accum)
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
	alu.UpdateAuxiliaryCarry(accum, result)
	alu.ClearCarry()
}

// OrAccumulator performs a bitwise OR operation on the contents of the accumulator and the operand.
// Flags Z, S, P, CY, and AC are updated. The CY and AC flags are cleared.
func (alu *ALUImpl) OrAccumulator(operand uint8) {
	var accum uint8
	var result uint8

	alu.A.Read8(&accum)
	result = accum | operand
	alu.A.Write8(result)
	alu.UpdateFlagsExceptCarry(result)
	alu.ClearCarry()
	alu.ClearAuxiliaryCarry()
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
	alu.ClearAuxiliaryCarry()
}

// DecimalAdjustAccumulator converts the value of the accumulator to Binary-Coded-Decimal digits to form two 4-bit
// BCD digits. All flags are affected.
func (alu *ALUImpl) DecimalAdjustAccumulator() {
	/**
		The eight-bit number in the accumulator is adjusted
		to form two four-bit Binary-Coded-Decimal digits by
		the following process:

		1. If the value of the least significant 4 bits of the
		   accumulator is greater than 9 or if the AC flag
	       is set, 6 is added to the accumulator.

		2. If the value of the most significant 4 bits of the
		   accumulator is now greater than 9, or if the CY
		   flag is set, 6 is added to the most significant 4
		   bits of the accumulator.

		NOTE: All flags are affected.
	*/

	var orig uint8
	alu.A.Read8(&orig)

	a := uint16(orig)

	lsb := a & 0x0F

	if lsb > 9 || alu.IsAuxiliaryCarry() {
		a += 6
	}

	lsb = a & 0x0F
	msb := a >> 4

	if msb > 9 || alu.IsCarry() {
		msb += 6
		a = (msb << 4) | lsb
	}

	alu.A.Write8(uint8(a & 0xFF))
	alu.UpdateFlags(orig, a)
}

// ComplementAccumulator negates the value of the accumulator such that 1 bits become 0 bits and 0 bits become 1 bits.
func (alu *ALUImpl) ComplementAccumulator() {
	var a uint8
	alu.A.Read8(&a)
	alu.A.Write8(0xFF ^ a)
}

// CompareAccumulator subtracts the operand from the accumulator. (A) - (r). The accumulator remains unchanged.
// Condition flags are updated as a result. The flags Z, S, P, CY, and AC are affected. The Z flag is set to 1 if
// (A) = (r). The CY flag is set to 1 if (A) < (r).
func (alu *ALUImpl) CompareAccumulator(operand uint8) {
	var a uint8
	alu.A.Read8(&a)

	difference := a - operand

	alu.UpdateSign(difference)
	alu.UpdateParity(difference)
	alu.UpdateAuxiliaryCarry(a, difference)

	alu.ClearZero()
	alu.ClearCarry()

	if a == operand {
		alu.SetZero()
	} else if a < operand {
		alu.SetCarry()
	}
}
