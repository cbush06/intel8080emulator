package alu

import (
	"github.com/cbush06/intel8080emulator/memory"
)

// ALU is the interface to the Intel 8080's ALU
type ALU interface {
	ConditionFlags
	GetA() *memory.Register
	SetA(a *memory.Register)
	UpdateFlags(original uint8, new uint8)
	UpdateFlagsExceptCarry(value uint8)
	AddImmediate(addend uint8)
	AddImmediateWithCarry(addend uint8)
	DoubleAdd(addend1 uint16, addend2 uint16) uint16
	SubImmediate(subtrahend uint8)
	SubImmediateWithBorrow(subtrahend uint8)
	Increment(value uint8) uint8
	IncrementDouble(value uint16) uint16
	DecrementDouble(value uint16) uint16
	Decrement(value uint8) uint8
	RotateRight()
	RotateLeft()
	RotateRightThroughCarry()
	RotateLeftThroughCarry()
	AndAccumulator(operand uint8)
	OrAccumulator(operand uint8)
	XOrAccumulator(operand uint8)
	DecimalAdjustAccumulator()
	ComplementAccumulator()
	CompareAccumulator(operand uint8)
}
