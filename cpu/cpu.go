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
	cpu.EnableInterrupts = true
	cpu.A = *memory.NewRegister(0)
	cpu.BC = *memory.NewRegisterPair(0, 0)
	cpu.DE = *memory.NewRegisterPair(0, 0)
	cpu.HL = *memory.NewRegisterPair(0, 0)
	cpu.WZ = *memory.NewRegisterPair(0, 0)
	cpu.ALU = alu.NewALU(&cpu.A)
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
	cpu.exec(OpCode(cpu.Memory[cpu.ProgramCounter]))
}

// InterruptInstructionCycle disables the EnableInterrupts flag, reads an OpCode off the DataBus
// and executes that OpCode and re-enables the EnableInterrupts flag. The ProgramCounter is not
// incremented prior to executing the OpCode.
func (cpu *CPU) InterruptInstructionCycle(interruptCmd uint8) {
	cpu.EnableInterrupts = false
	cpu.exec(OpCode(interruptCmd))
	cpu.EnableInterrupts = true
}

// exec executes the provided opcode
func (cpu *CPU) exec(opcode OpCode) {
	switch opcode {
	case NOP:
		break
	case CALL:
		cpu.Call()
	case RST0, RST1, RST2, RST3, RST4, RST5, RST6, RST7:
		cpu.Restart(opcode)
	case RET:
		cpu.Return()
	case JMP:
		cpu.ProgramCounter = cpu.getJumpAddress()
	case JNZ:
		cpu.executeJumpIfTrue(!cpu.ALU.IsZero())
	case JZ:
		cpu.executeJumpIfTrue(cpu.ALU.IsZero())
	case JNC:
		cpu.executeJumpIfTrue(!cpu.ALU.IsCarry())
	case JC:
		cpu.executeJumpIfTrue(cpu.ALU.IsCarry())
	case JPO:
		cpu.executeJumpIfTrue(!cpu.ALU.IsParity())
	case JPE:
		cpu.executeJumpIfTrue(cpu.ALU.IsParity())
	case JP:
		cpu.executeJumpIfTrue(!cpu.ALU.IsSign())
	case JM:
		cpu.executeJumpIfTrue(cpu.ALU.IsSign())
	case PUSHB, PUSHD, PUSHH:
		rp := cpu.getOpCodeRegisterPair(opcode)
		cpu.Push(rp)
	case PUSHPSW:
		cpu.PushProcessorStatusWord()
	case POPB, POPD, POPH:
		rp := cpu.getOpCodeRegisterPair(opcode)
		cpu.Pop(rp)
	case POPPSW:
		cpu.PopProcessorStatusWord()
	case LDA:
		cpu.LoadAccumulatorDirect(cpu.Memory[cpu.ProgramCounter+1], cpu.Memory[cpu.ProgramCounter+2])
	case LDAXB, LDAXD:
		rp := cpu.getOpCodeRegisterPair(opcode)
		cpu.LoadAccumulatorIndirect(rp)
	case LXIB, LXID, LXIH, LXISP:
		rp := cpu.getOpCodeRegisterPair(opcode)
		cpu.LoadRegisterPairImmediate(rp, cpu.Memory[cpu.ProgramCounter+1], cpu.Memory[cpu.ProgramCounter+2])
	case STA:
		cpu.StoreAccumulatorDirect(cpu.Memory[cpu.ProgramCounter+1], cpu.Memory[cpu.ProgramCounter+2])
	case STAXB, STAXD:
		rp := cpu.getOpCodeRegisterPair(opcode)
		cpu.StoreAccumulatorIndirect(rp)
	case LHLD:
		cpu.LoadHandLDirect(cpu.Memory[cpu.ProgramCounter+1], cpu.Memory[cpu.ProgramCounter+2])
	case SHLD:
		cpu.StoreHandLDirect(cpu.Memory[cpu.ProgramCounter+1], cpu.Memory[cpu.ProgramCounter+2])
	case DCRA, DCRB, DCRC, DCRD, DCRE, DCRH, DCRL:
		r := cpu.getOpCodeRegisterDestination(opcode)
		cpu.DecrementRegister(r)
	case DCRM:
		cpu.DecrementMemory()
	case DCXB, DCXD, DCXH, DCXSP:
		rp := cpu.getOpCodeRegisterPair(opcode)
		cpu.DecrementRegisterPair(rp)
	case INRA, INRB, INRC, INRD, INRE, INRH, INRL:
		r := cpu.getOpCodeRegisterDestination(opcode)
		cpu.IncrementRegister(r)
	case INRM:
		cpu.DecrementMemory()
	case INXB, INXD, INXH, INXSP:
		rp := cpu.getOpCodeRegisterPair(opcode)
		cpu.IncrementRegisterPair(rp)
	case ADI:
		cpu.AddImmediate()
	case DADB, DADD, DADH, DADSP:
		rp := cpu.getOpCodeRegisterPair(opcode)
		cpu.DoubleAdd(rp)
	case MOVAA, MOVAB, MOVAC, MOVAD, MOVAE, MOVAH, MOVAL, MOVBA, MOVBB, MOVBC, MOVBD, MOVBE, MOVBH, MOVBL, MOVCA, MOVCB,
		 MOVCC,  MOVCD, MOVCE, MOVCH, MOVCL, MOVDA, MOVDB, MOVDC, MOVDD, MOVDE, MOVDH, MOVDL, MOVEA, MOVEB, MOVEC,
		 MOVED, MOVEE, MOVEH, MOVEL, MOVHA, MOVHB, MOVHC, MOVHD, MOVHE, MOVHH, MOVHL, MOVLA, MOVLB, MOVLC, MOVLD, MOVLE,
		 MOVLH, MOVLL:
		r1 := cpu.getOpCodeRegisterDestination(opcode)
		r2 := cpu.getOpCodeRegisterSource(opcode)
		cpu.MoveRegister(r1, r2)
	case MOVMA, MOVMB, MOVMC, MOVMD, MOVME, MOVMH, MOVML:
		r := cpu.getOpCodeRegisterSource(opcode)
		cpu.MoveToMemory(r)
	case MVIA, MVIB, MVIC, MVID, MVIE, MVIH, MVIL:
		r := cpu.getOpCodeRegisterDestination(opcode)
		cpu.MoveImmediate(r, cpu.Memory[cpu.ProgramCounter+1])
	case CMA:

	case MVIM:
		cpu.MoveToMemoryImmediate(cpu.Memory[cpu.ProgramCounter+1])
	case DAA:
		cpu.DecimalAccumulatorAdjust()
	case RRC:
		cpu.RotateRight()
	case RAR:
		cpu.RotateRightThroughCarry()
	case RLC:
		cpu.RotateLeft()
	case RAL:
		cpu.RotateLeftThroughCarry()
	case ANAA, ANAB, ANAC, ANAD, ANAE, ANAH, ANAL:
		r := cpu.getOpCodeRegisterSource(opcode)
		cpu.AndRegister(r)
	case ANAM:
		cpu.AndMemory()
	case XRAA, XRAB, XRAC, XRAD, XRAE, XRAH, XRAL:
		r := cpu.getOpCodeRegisterSource(opcode)
		cpu.XOrRegister(r)
	case XRAM:
		cpu.XOrMemory()
	}
}

func (cpu *CPU) getOpCodeRegisterPair(opcode OpCode) *memory.RegisterPair {
	rpIndex := (opcode & 0x30) >> 4
	return cpu.RegisterPairLookup[rpIndex]
}

func (cpu *CPU) getOpCodeRegisterDestination(opcode OpCode) *memory.Register {
	rIndex := (opcode & 0x38) >> 3
	return cpu.RegisterLookup[rIndex]
}

func (cpu *CPU) getOpCodeRegisterSource(opcode OpCode) *memory.Register {
	rIndex := (opcode & 0x07)
	return cpu.RegisterLookup[rIndex]
}

func (cpu *CPU) getJumpAddress() uint16 {
	return (uint16(cpu.Memory[cpu.ProgramCounter+2]) << 8) | uint16(cpu.Memory[cpu.ProgramCounter+1])
}

func (cpu *CPU) executeJumpIfTrue(condition bool) {
	if condition {
		cpu.ProgramCounter = cpu.getJumpAddress()
	} else {
		cpu.ProgramCounter += 2
	}
}
