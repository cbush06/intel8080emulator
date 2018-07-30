package cpu

import (
	"github.com/cbush06/intel8080emulator/alu"
	"github.com/cbush06/intel8080emulator/memory"
)

// CPU represents the collection of components that comprise the 8080's central processing unit. In short,
// it encapsulates the ALU, registers, and interpreter.
type CPU struct {
	ProgramCounter     uint16
	SP                 memory.RegisterPair
	EnableInterrupts   bool
	A                  *memory.Register
	BC                 memory.RegisterPair
	B                  *memory.Register
	C                  *memory.Register
	DE                 memory.RegisterPair
	D                  *memory.Register
	E                  *memory.Register
	HL                 memory.RegisterPair
	H                  *memory.Register
	L                  *memory.Register
	WZ                 memory.RegisterPair
	W                  *memory.Register
	Z                  *memory.Register
	ALU                alu.ALU
	Memory             []uint8
	RegisterLookup     [8]*memory.Register
	RegisterPairLookup [4]*memory.RegisterPair
}

// Init must be called before using the CPU. This method initializes pointers and other elements necessary for the CPU to function correctly.
func (c *CPU) Init() {
	c.A = &c.ALU.A
	c.B = &c.BC.High
	c.C = &c.BC.Low
	c.D = &c.DE.High
	c.E = &c.DE.Low
	c.H = &c.HL.High
	c.L = &c.HL.Low
	c.W = &c.WZ.High
	c.Z = &c.WZ.Low
	c.Memory = make([]uint8, 16384) // 16KB

	c.RegisterLookup[0] = c.B
	c.RegisterLookup[1] = c.C
	c.RegisterLookup[2] = c.D
	c.RegisterLookup[3] = c.E
	c.RegisterLookup[4] = c.H
	c.RegisterLookup[5] = c.L
	c.RegisterLookup[6] = nil
	c.RegisterLookup[7] = c.A

	c.RegisterPairLookup[0] = &c.BC
	c.RegisterPairLookup[1] = &c.DE
	c.RegisterPairLookup[2] = &c.HL
	c.RegisterPairLookup[3] = &c.SP
}

// Exec increments the Program Counter and executes the next opcode.
func (c *CPU) Exec() {
	c.ProgramCounter++

	opcode := OpCode(c.Memory[c.ProgramCounter])

	switch opcode {
	case NOP:
		break

	case LXIB:
	case LXID:
	case LXIH:
	case LXISP:
		rp := c.getOpCodeRegisterPair(opcode)
		c.LoadRegisterPairImmediate(rp, c.Memory[c.ProgramCounter+1], c.Memory[c.ProgramCounter+2])
		c.ProgramCounter += 2

	case STAXB:
	case STAXD:
		rp := c.getOpCodeRegisterPair(opcode)
		c.StoreAccumulatorIndirect(rp)
		c.ProgramCounter++

	case DCRA:
	case DCRB:
	case DCRC:
	case DCRD:
	case DCRE:
	case DCRH:
	case DCRL:
		r := c.getOpCodeRegister(opcode)
		c.DecrementRegister(r)

	case DCRM:
		c.DecrementMemory()
		c.ProgramCounter += 2

	case DADB:
	case DADD:
	case DADH:
	case DADSP:
		rp := c.getOpCodeRegisterPair(opcode)
		c.DoubleAdd(rp)
		c.ProgramCounter += 2

	case MVIA:
	case MVIB:
	case MVIC:
	case MVID:
	case MVIE:
	case MVIH:
	case MVIL:
		r := c.getOpCodeRegister(opcode)
		c.MoveImmediate(r, c.Memory[c.ProgramCounter+1])
		c.ProgramCounter++

	case MVIM:
		c.MoveToMemoryImmediate(c.Memory[c.ProgramCounter+1])
		c.ProgramCounter += 2
	}
}

func (c *CPU) getOpCodeRegisterPair(opcode OpCode) *memory.RegisterPair {
	rpIndex := (opcode & 0x30) >> 4
	return c.RegisterPairLookup[rpIndex]
}

func (c *CPU) getOpCodeRegister(opcode OpCode) *memory.Register {
	rIndex := (opcode & 0x38) >> 3
	return c.RegisterLookup[rIndex]
}
