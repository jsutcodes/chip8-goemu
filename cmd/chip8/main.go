package main

import (
	"fmt"

	"github.com/jsutcodes/chip8-goemu/internal/emulator"
)

func main() {
	fmt.Println("Starting CHIP-8 Emulator")
	emu := emulator.NewEmulator()
	emu.Start()
}
