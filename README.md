# intel8080emulator

This emulator has a full implementation of the Intel8080 instruction set. It has been tested against the Kelly Smith test.

## Roadmap

I plan to use Go's RPC capabilities to make this extensible for use with various harnesses. Specifically, I intend to write 
a Space Invaders hardware emulator (buttons, joy stick, and display) that will integrate with this emulator via RPC. Other plans
include adding the ability to external disks, running CP/M, and, hopefully, I can play Zork on it.
