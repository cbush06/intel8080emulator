package main

//go:generate mockgen -destination=alu/mocks/condition_flags_mock.go -package=alu github.com/cbush06/intel8080emulator/alu ConditionFlags
//go:generate mockgen -destination=alu/mocks/alu_mock.go -package=alu github.com/cbush06/intel8080emulator/alu ALU

import (
	"fmt"
	"github.com/cbush06/intel8080emulator/cpu"
	"os"
)

// CPUInterface is a struct comprising multiple go channels to allow IPC between
// this CPU, the machine using it, and any peripherals.
type CPUInterface struct {
	Interrupt <-chan uint8
	PowerOff  <-chan bool
	DataBus   chan uint8
	Memory    []uint8
	cpu       *cpu.CPU
}

// StartCPU begins the CPU fetch-execute cycle and loads the specified program.
func StartCPU(program []byte, memShift uint16) *CPUInterface {
	var mainCpu = new(cpu.CPU)
	mainCpu.Init()

	cpuInt := &CPUInterface{
		Interrupt: make(chan uint8),
		PowerOff:  make(chan bool),
		DataBus:   make(chan uint8),
		Memory:    mainCpu.Memory,
		cpu:       mainCpu,
	}

	// Set ProgramCounter to execution starting point
	mainCpu.ProgramCounter = memShift

	// Copy program into working memory
	copy(mainCpu.Memory[memShift:], program)

	return cpuInt
}

func (cpuInt *CPUInterface) TickCPU() {
	// Check for PowerOff command
	if powerOff := <-cpuInt.PowerOff; powerOff {
		os.Exit(0)
	}

	// Read DataBus in
	if data := <-cpuInt.DataBus; data > 0 {
		cpuInt.cpu.DataBus.Write8(data)
	}

	// Check for Interrupt; if set, execute interrupt instruction cycle
	select {
	case interruptCommand := <-cpuInt.Interrupt:
		fmt.Printf("INTERRUPT 0x%2X\n", interruptCommand)
		cpuInt.cpu.DataBus.Write8(interruptCommand)
		cpuInt.cpu.InterruptInstructionCycle()
	default:
		cpuInt.cpu.StandardInstructionCycle()
	}

	// Write DataBus out
	var data uint8
	cpuInt.cpu.DataBus.Read8(&data)
	cpuInt.DataBus <- data
}
