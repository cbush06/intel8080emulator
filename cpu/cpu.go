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
	Write              bool
	DataBus            memory.Register
	AddressBus         memory.RegisterPair
	A                  memory.Register
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
func (cpu *CPU) Init() {
	cpu.ALU.A = &cpu.A
	cpu.B = &cpu.BC.High
	cpu.C = &cpu.BC.Low
	cpu.D = &cpu.DE.High
	cpu.E = &cpu.DE.Low
	cpu.H = &cpu.HL.High
	cpu.L = &cpu.HL.Low
	cpu.W = &cpu.WZ.High
	cpu.Z = &cpu.WZ.Low
	cpu.Memory = make([]uint8, 16384) // 16KB

	cpu.RegisterLookup[0] = cpu.B
	cpu.RegisterLookup[1] = cpu.C
	cpu.RegisterLookup[2] = cpu.D
	cpu.RegisterLookup[3] = cpu.E
	cpu.RegisterLookup[4] = cpu.H
	cpu.RegisterLookup[5] = cpu.L
	cpu.RegisterLookup[6] = nil
	cpu.RegisterLookup[7] = &cpu.A

	cpu.RegisterPairLookup[0] = &cpu.BC
	cpu.RegisterPairLookup[1] = &cpu.DE
	cpu.RegisterPairLookup[2] = &cpu.HL
	cpu.RegisterPairLookup[3] = &cpu.SP
}

// StandardInstructionCycle increments the Program Counter and executes the next instruction
func (cpu *CPU) StandardInstructionCycle() {
	cpu.ProgramCounter++
	cpu.exec(OpCode(cpu.Memory[cpu.ProgramCounter]))
}

// InterruptInstructionCycle disables the EnableInterrupts flag, reads an OpCode off the DataBus
// and executes that OpCode. The ProgramCounter is not incremented prior to executing the OpCode.
func (cpu *CPU) InterruptInstructionCycle() {
	var opcode uint8
	cpu.DataBus.Read8(&opcode)
	cpu.EnableInterrupts = false
	cpu.exec(OpCode(opcode))
}

// exec executes the provided opcode
func (cpu *CPU) exec(opcode OpCode) {
	switch opcode {
	case NOP:
		break

	case CALL:
		cpu.Call()

	case RET:
		cpu.Return()

	case JMP:
		// Subtract 1 because 1 will be added on the next execution
		cpu.ProgramCounter = cpu.getJumpAddress() - 1
	case JNZ:
		cpu.executeJumpIfTrue(!cpu.ALU.Zero)
	case JZ:
		cpu.executeJumpIfTrue(cpu.ALU.Zero)
	case JNC:
		cpu.executeJumpIfTrue(!cpu.ALU.Carry)
	case JC:
		cpu.executeJumpIfTrue(cpu.ALU.Carry)
	case JPO:
		cpu.executeJumpIfTrue(!cpu.ALU.Parity)
	case JPE:
		cpu.executeJumpIfTrue(cpu.ALU.Parity)
	case JP:
		cpu.executeJumpIfTrue(!cpu.ALU.Sign)
	case JM:
		cpu.executeJumpIfTrue(cpu.ALU.Sign)

	case PUSHB:
	case PUSHD:
	case PUSHH:
		rp := cpu.getOpCodeRegisterPair(opcode)
		cpu.Push(rp)

	case PUSHPSW:
		cpu.PushProcessorStatusWord()

	case POPB:
	case POPD:
	case POPH:
		rp := cpu.getOpCodeRegisterPair(opcode)
		cpu.Pop(rp)

	case POPPSW:
		cpu.PopProcessorStatusWord()

	case LDA:
		cpu.LoadAccumulatorDirect(cpu.Memory[cpu.ProgramCounter+1], cpu.Memory[cpu.ProgramCounter+2])
		cpu.ProgramCounter += 2

	case LDAXB:
	case LDAXD:
		rp := cpu.getOpCodeRegisterPair(opcode)
		cpu.LoadAccumulatorIndirect(rp)

	case LXIB:
	case LXID:
	case LXIH:
	case LXISP:
		rp := cpu.getOpCodeRegisterPair(opcode)
		cpu.LoadRegisterPairImmediate(rp, cpu.Memory[cpu.ProgramCounter+1], cpu.Memory[cpu.ProgramCounter+2])
		cpu.ProgramCounter += 2

	case STA:
		cpu.StoreAccumulatorDirect(cpu.Memory[cpu.ProgramCounter+1], cpu.Memory[cpu.ProgramCounter+2])
		cpu.ProgramCounter += 2

	case STAXB:
	case STAXD:
		rp := cpu.getOpCodeRegisterPair(opcode)
		cpu.StoreAccumulatorIndirect(rp)

	case DCRA:
	case DCRB:
	case DCRC:
	case DCRD:
	case DCRE:
	case DCRH:
	case DCRL:
		r := cpu.getOpCodeRegister(opcode)
		cpu.DecrementRegister(r)
	case DCRM:
		cpu.DecrementMemory()

	case DCXB:
	case DCXD:
	case DCXH:
	case DCXSP:
		rp := cpu.getOpCodeRegisterPair(opcode)
		cpu.DecrementRegisterPair(rp)

	case INRA:
	case INRB:
	case INRC:
	case INRD:
	case INRE:
	case INRH:
	case INRL:
		r := cpu.getOpCodeRegister(opcode)
		cpu.IncrementRegister(r)
	case INRM:
		cpu.DecrementMemory()

	case INXB:
	case INXD:
	case INXH:
	case INXSP:
		rp := cpu.getOpCodeRegisterPair(opcode)
		cpu.IncrementRegisterPair(rp)

	case ADI:
		cpu.AddImmediate()
		cpu.ProgramCounter++

	case DADB:
	case DADD:
	case DADH:
	case DADSP:
		rp := cpu.getOpCodeRegisterPair(opcode)
		cpu.DoubleAdd(rp)

	case MOVAA:
	case MOVAB:
	case MOVAC:
	case MOVAD:
	case MOVAE:
	case MOVAH:
	case MOVAL:
	case MOVBA:
	case MOVBB:
	case MOVBC:
	case MOVBD:
	case MOVBE:
	case MOVBH:
	case MOVBL:
	case MOVCA:
	case MOVCB:
	case MOVCC:
	case MOVCD:
	case MOVCE:
	case MOVCH:
	case MOVCL:
	case MOVDA:
	case MOVDB:
	case MOVDC:
	case MOVDD:
	case MOVDE:
	case MOVDH:
	case MOVDL:
	case MOVEA:
	case MOVEB:
	case MOVEC:
	case MOVED:
	case MOVEE:
	case MOVEH:
	case MOVEL:
	case MOVHA:
	case MOVHB:
	case MOVHC:
	case MOVHD:
	case MOVHE:
	case MOVHH:
	case MOVHL:
	case MOVLA:
	case MOVLB:
	case MOVLC:
	case MOVLD:
	case MOVLE:
	case MOVLH:
	case MOVLL:
		r1 := cpu.getOpCodeRegister(opcode)
		r2 := cpu.getOpCodeRegister2(opcode)
		cpu.MoveRegister(r1, r2)

	case MVIA:
	case MVIB:
	case MVIC:
	case MVID:
	case MVIE:
	case MVIH:
	case MVIL:
		r := cpu.getOpCodeRegister(opcode)
		cpu.MoveImmediate(r, cpu.Memory[cpu.ProgramCounter+1])
		cpu.ProgramCounter++
	case MVIM:
		cpu.MoveToMemoryImmediate(cpu.Memory[cpu.ProgramCounter+1])
		cpu.ProgramCounter++

	case RRC:
		cpu.RotateRight()

	case ANAA:
	case ANAB:
	case ANAC:
	case ANAD:
	case ANAE:
	case ANAH:
	case ANAL:
		r := cpu.getOpCodeRegister(opcode)
		cpu.AndRegister(r)
	case ANAM:
		cpu.AndMemory()

	case XRAA:
	case XRAB:
	case XRAC:
	case XRAD:
	case XRAE:
	case XRAH:
	case XRAL:
		r := cpu.getOpCodeRegister(opcode)
		cpu.XOrRegister(r)
	case XRAM:
		cpu.XOrMemory()
	}
}

func (cpu *CPU) getOpCodeRegisterPair(opcode OpCode) *memory.RegisterPair {
	rpIndex := (opcode & 0x30) >> 4
	return cpu.RegisterPairLookup[rpIndex]
}

func (cpu *CPU) getOpCodeRegister(opcode OpCode) *memory.Register {
	rIndex := (opcode & 0x38) >> 3
	return cpu.RegisterLookup[rIndex]
}

func (cpu *CPU) getOpCodeRegister2(opcode OpCode) *memory.Register {
	rIndex := (opcode & 0x07)
	return cpu.RegisterLookup[rIndex]
}

func (cpu *CPU) getJumpAddress() uint16 {
	return (uint16(cpu.Memory[cpu.ProgramCounter+2]) << 8) | uint16(cpu.Memory[cpu.ProgramCounter+1])
}

func (cpu *CPU) executeJumpIfTrue(condition bool) {
	if condition {
		// Subtract 1 because 1 will be added on the next execution
		cpu.ProgramCounter = cpu.getJumpAddress() - 1
	} else {
		cpu.ProgramCounter += 2
	}
}
