package alu

import (
	"github.com/cbush06/intel8080emulator/memory"
)

// ALU is the interface to the Intel 8080's ALU
type ALU interface {
	ConditionFlags
	GetA() *memory.Register
	SetA(a *memory.Register)
	UpdateFlags(value uint16)
	UpdateFlagsExceptCarry(value uint16)
	AddImmediate(addend uint8)
	AddImmediateWithCarry(addend uint8)
	DoubleAdd(addend1 uint16, addend2 uint16) uint16
	SubImmediate(addend uint8)
	SubImmediateWithBorrow(addend uint8)
	Increment(addend uint8) uint8
	IncrementDouble(addend uint16) uint16
	DecrementDouble(addend uint16) uint16
	Decrement(addend uint8) uint8
	RotateRight()
	AndAccumulator(operand uint8)
	XOrAccumulator(operand uint8)
}
