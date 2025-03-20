package cpu

import (
	"github.com/jsutcodes/chip8-goemu/internal/memory"
	"honnef.co/go/tools/printf"
)

type CPU struct {
	V   [16]byte // General purpose registers (V0 to VF)
	I   uint16   // Index register
	PC  uint16   // Program counter
	SP  byte     // Stack pointer
	DT  byte     // Delay timer
	ST  byte     // Sound timer
	RAM *memory.Memory
}

// var opcodes := []uint16{
// 	0xNNN, // Call RCA 1802 program at address NNN
// 	0x00E0, // Clear the display
// 	0x00EE, // Return from a subroutine
// 	0x1NNN, // Jump to address NNN
// 	0x2NNN, // Call subroutine at NNN
// 	0x3XNN, // Skip next instruction if VX equals NN
// 	0x4XNN, // Skip next instruction if VX doesn't equal NN
// 	0x5XY0, // Skip next instruction if VX equals VY
// 	0x6XNN, // Set VX to NN
// 	0x7XNN, // Add NN to VX
// 	0x8XY0, // Set VX to the value of VY
// 	0x8XY1, // Set VX to VX OR VY
// 	0x8XY2, // Set VX to VX AND VY
// 	0x8XY3, // Set VX to VX XOR VY
// 	0x8XY4, // Add VY to VX
// 	0x8XY5, // Subtract VY from VX
// 	0x8XY6, // Shift VX right by one
// 	0x8XY7, // Set VX to VY minus VX
// 	0x8XYE, // Shift VX left by one
// 	0x9XY0, // Skip next instruction if VX doesn't equal VY
// 	0xANNN, // Set I to the address NNN
// 	0xBNNN, // Jump to the address NNN plus V0
// 	0xCXNN, // Set VX to a random number and NN
// 	0xDXYN, // Draw a sprite at coordinate (VX, VY) that has a width of 8 pixels and a height of N pixels
// 	0xEX9E, // Skip next instruction if the key stored in VX is pressed
// 	0xEXA1, // Skip next instruction if the key stored in VX isn't pressed
// 	0xFX07, // Set VX to the value of the delay timer
// 	0xFX0A, // Wait for a key press and store the result in VX
// 	0xFX15, // Set the delay timer to VX
// 	0xFX18, // Set the sound timer to VX
// 	0xFX1E, // Add VX to I
// 	0xFX29, // Set I to the location of the sprite for the character in VX
// 	0xFX33, // Store the binary-coded decimal representation of VX at the addresses I, I+1, and I+2
// 	0xFX55, // Store V0 to VX in memory starting at address I
// 	0xFX65, // Fill V0 to VX with values from memory starting at address I
// }

func NewCPU(RAM *memory.Memory) *CPU {
	return &CPU{
		PC:  0x200, // Program counter starts at 0x200
		RAM: RAM,
	}
}

func (c *CPU) fetch(RAM *memory.Memory) uint16 {
	byte1, _ := RAM.ReadByte(c.PC)
	// if err1 != nil {
	// 	// handle error
	// 	return 0
	// }
	byte2, _ := RAM.ReadByte(c.PC + 1)
	// if err2 != nil {
	// 	// handle error
	// 	return 0
	// }
	return uint16(byte1)<<8 | uint16(byte2)
}

func (c *CPU) decode(opcode uint16) {
	// Decode opcode
}

func (c *CPU) execute() {
	// Execute opcode
}

func printState(c *CPU) {
	printf.Printf("PC: 0x%X\n", c.PC)
	printf.Printf("I: 0x%X\n", c.I)
	printf.Printf("V: %v\n", c.V)
	printf.Printf("DT: 0x%X\n", c.DT)
	printf.Printf("ST: 0x%X\n", c.ST)
}

func (c *CPU) Step(verbose bool, RAM *memory.Memory) {

	opcode := c.fetch(RAM)

	if verbose {
		printState(c)
		printf.Printf("Opcode: 0x%X\n", c.fetch())
	}

	c.decode(opcode)
	c.execute()
}
