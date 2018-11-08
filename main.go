package main

import (
	"fmt"
	"os"
	"time"

	"bitbucket.org/avd/go-ipc/fifo"
	"github.com/cbush06/intel8080emulator/cpu"
)

func main() {
	var cpu cpu.CPU
	cpu.Init()

	// Open Interrupt pipe
	interruptPipe, err := fifo.New("intel8080_interrupt", os.O_CREATE|os.O_RDONLY|fifo.O_NONBLOCK, 0666)
	if err != nil {
		panic("Unable to connect to named pipe [intell8080_interrupt]")
	}

	defer interruptPipe.Close()

	interruptData := []byte{0}
	for {
		if cpu.EnableInterrupts {
			// Check for Interrupt; if set, execute interrupt instruction cycle
			if readLen, err := interruptPipe.Read(interruptData); readLen == 1 && err == nil {
				fmt.Println("INTERRUPT")
				// cpu.InterruptInstructionCycle()
				continue
			}
		}
		fmt.Println("Standard Execution")
		time.Sleep(2 * time.Second)
	}

}
