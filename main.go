package main

//go:generate mockgen -destination=alu/mocks/condition_flags_mock.go -package=alu github.com/cbush06/intel8080emulator/alu ConditionFlags
//go:generate mockgen -destination=alu/mocks/alu_mock.go -package=alu github.com/cbush06/intel8080emulator/alu ALU

import (
	"fmt"
	"github.com/cbush06/intel8080emulator/cpu"
	"io/ioutil"
	"log"
)

// CPUInterface is a struct comprising multiple go channels to allow IPC between
// this CPU, the machine using it, and any peripherals.
type CPUInterface struct {
	Interrupt chan uint8
}

func main() {
	if bytes, e := ioutil.ReadFile("/Users/cbush/projects/intel8080emulator/cpudiag.bin"); e != nil {
		log.Fatal(e)
	} else {
		bytes[368] = 0x7 // Skip to
		startCPU(bytes, 0x100)
	}
}

// StartCPU begins the CPU fetch-execute cycle and loads the specified program.
func startCPU(program []byte, memShift uint16) *CPUInterface {
	cpuInt := &CPUInterface{
		Interrupt: make(chan uint8),
	}

	var mainCpu = new(cpu.CPU)
	mainCpu.Init()

	// Set ProgramCounter to execution starting point
	mainCpu.ProgramCounter = memShift

	// Copy program into working memory
	copy(mainCpu.Memory[memShift:], program)

	// Where the RUBBER MEETS THE ROAD
	for {
		// Handle Interrupts (button presses, vertical blank interrupts from screen, etc.)
		if mainCpu.InterruptsEnabled {
			// Check for Interrupt; if set, execute interrupt instruction cycle
			select {
			case interruptCommand := <-cpuInt.Interrupt:
				fmt.Printf("INTERRUPT 0x%2X\n", interruptCommand)
				mainCpu.InterruptInstructionCycle(interruptCommand)
			default:
				mainCpu.StandardInstructionCycle()
			}
		}
	}

	return cpuInt
}
