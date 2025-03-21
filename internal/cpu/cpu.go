package cpu

import (
	"math/rand"

	"github.com/jsutcodes/chip8-goemu/internal/display"
	"github.com/jsutcodes/chip8-goemu/internal/input"
	"github.com/jsutcodes/chip8-goemu/internal/memory"
	"honnef.co/go/tools/printf"
)

type CPU struct {
	V       [16]byte   // General purpose registers (V0 to VF)
	I       uint16     // Index register
	PC      uint16     // Program counter
	SP      byte       // Stack pointer
	Stack   [16]uint16 // Stack
	DT      byte       // Delay timer
	ST      byte       // Sound timer
	RAM     *memory.Memory
	Display *display.Display
	Input   *input.Keypad
}

// Opcode Table:
// | Opcode | Description                                                                 |
// |--------|-----------------------------------------------------------------------------|
// | 0xNNN  | Call RCA 1802 program at address NNN                                        |
// | 0x00E0 | Clear the display                                                           |
// | 0x00EE | Return from a subroutine                                                    |
// | 0x1NNN | Jump to address NNN                                                         |
// | 0x2NNN | Call subroutine at NNN                                                      |
// | 0x3XNN | Skip next instruction if VX equals NN                                       |
// | 0x4XNN | Skip next instruction if VX doesn't equal NN                                |
// | 0x5XY0 | Skip next instruction if VX equals VY                                       |
// | 0x6XNN | Set VX to NN                                                                |
// | 0x7XNN | Add NN to VX                                                                |
// | 0x8XY0 | Set VX to the value of VY                                                   |
// | 0x8XY1 | Set VX to VX OR VY                                                          |
// | 0x8XY2 | Set VX to VX AND VY                                                         |
// | 0x8XY3 | Set VX to VX XOR VY                                                         |
// | 0x8XY4 | Add VY to VX                                                                |
// | 0x8XY5 | Subtract VY from VX                                                         |
// | 0x8XY6 | Shift VX right by one                                                       |
// | 0x8XY7 | Set VX to VY minus VX                                                       |
// | 0x8XYE | Shift VX left by one                                                        |
// | 0x9XY0 | Skip next instruction if VX doesn't equal VY                                |
// | 0xANNN | Set I to the address NNN                                                    |
// | 0xBNNN | Jump to the address NNN plus V0                                             |
// | 0xCXNN | Set VX to a random number and NN                                            |
// | 0xDXYN | Draw a sprite at coordinate (VX, VY) that has a width of 8 pixels and a height of N pixels |
// | 0xEX9E | Skip next instruction if the key stored in VX is pressed                    |
// | 0xEXA1 | Skip next instruction if the key stored in VX isn't pressed                 |
// | 0xFX07 | Set VX to the value of the delay timer                                      |
// | 0xFX0A | Wait for a key press and store the result in VX                             |
// | 0xFX15 | Set the delay timer to VX                                                   |
// | 0xFX18 | Set the sound timer to VX                                                   |
// | 0xFX1E | Add VX to I                                                                 |
// | 0xFX29 | Set I to the location of the sprite for the character in VX                 |
// | 0xFX33 | Store the binary-coded decimal representation of VX at the addresses I, I+1, and I+2 |
// | 0xFX55 | Store V0 to VX in memory starting at address I                              |
// | 0xFX65 | Fill V0 to VX with values from memory starting at address I                 |

func NewCPU(RAM *memory.Memory, Display *display.Display, Input *input.Keypad) *CPU {
	return &CPU{
		PC:      0x200, // Program counter starts at 0x200
		RAM:     RAM,
		Display: Display,
		Input:   Input,
	}
}

func (c *CPU) fetch() uint16 {
	byte1, _ := c.RAM.ReadByte(c.PC)
	// if err1 != nil {
	// 	// handle error
	// 	return 0
	// }
	byte2, _ := c.RAM.ReadByte(c.PC + 1)
	// if err2 != nil {
	// 	// handle error
	// 	return 0
	// }
	return uint16(byte1)<<8 | uint16(byte2)
}

func (c *CPU) decodeAndExecute(opcode uint16) {
	// Decode opcode
	nibble := opcode & 0xF000 // get the top 4 bits

	switch nibble {
	case 0x0000:
		switch opcode {
		case 0x00E0:
			c.Display.Clear()
			break
		case 0x00EE:
			// Return from a subroutine
			c.SP--
			c.PC = c.Stack[c.SP]
			c.PC += 2
			break
		default:
			printf.Printf("Unknown opcode: 0x%X\n", opcode)
			break
		}
	case 0x1000:
		// Jump to address NNN
		c.PC = opcode & 0x0FFF
		break
	case 0x2000:
		// Call subroutine at NNN
		c.Stack[c.SP] = c.PC
		c.SP++
		c.PC = opcode & 0x0FFF
		break
	case 0x3000:
		// Skip next instruction if VX equals NN
		val := opcode & 0x00FF        // get lower 8 bits
		reg := (opcode & 0x0F00) >> 8 // get X
		c.PC += 2                     // next instruction
		if c.V[reg] == byte(val) {
			c.PC += 2 // skip this instruction
		}
		break
	case 0x4000:
		// Skip next instruction if VX doesn't equal NN
		val := opcode & 0x00FF        // get lower 8 bits
		reg := (opcode & 0x0F00) >> 8 // get X
		c.PC += 2                     // next instruction
		if c.V[reg] != byte(val) {
			c.PC += 2 // skip this instruction
		}
		break
	case 0x5000:
		// Skip next instruction if VX equals VY
		reg1 := (opcode & 0x0F00) >> 8 // get X
		reg2 := (opcode & 0x00F0) >> 4 // get Y
		c.PC += 2                      // next instruction
		if c.V[reg1] == c.V[reg2] {
			c.PC += 2 // skip this instruction
		}
		break
	case 0x6000:
		// Set VX to NN
		val := opcode & 0x00FF        // get lower 8 bits
		reg := (opcode & 0x0F00) >> 8 // get X
		c.V[reg] = byte(val)
		c.PC += 2 // next instruction
		break
	case 0x7000:
		// Add NN to VX
		val := opcode & 0x00FF        // get lower 8 bits
		reg := (opcode & 0x0F00) >> 8 // get X
		c.V[reg] += byte(val)
		c.PC += 2 // next instruction
		break
	case 0x8000:
		switch opcode & 0x000F {
		case 0x0000:
			// Set VX to the value of VY
			reg1 := (opcode & 0x0F00) >> 8 // get X
			reg2 := (opcode & 0x00F0) >> 4 // get Y
			c.V[reg1] = c.V[reg2]
			c.PC += 2 // next instruction
			break
		case 0x0001:
			// Set VX to VX OR VY
			reg1 := (opcode & 0x0F00) >> 8 // get X
			reg2 := (opcode & 0x00F0) >> 4 // get Y
			c.V[reg1] |= c.V[reg2]
			c.PC += 2 // next instruction
			break
		case 0x0002:
			// Set VX to VX AND VY
			reg1 := (opcode & 0x0F00) >> 8 // get X
			reg2 := (opcode & 0x00F0) >> 4 // get Y
			c.V[reg1] &= c.V[reg2]
			c.PC += 2 // next instruction
			break
		case 0x0003:
			// Set VX to VX XOR VY
			reg1 := (opcode & 0x0F00) >> 8 // get X
			reg2 := (opcode & 0x00F0) >> 4 // get Y
			c.V[reg1] ^= c.V[reg2]
			c.PC += 2 // next instruction
			break
		case 0x0004:
			// Add VY to VX
			reg1 := (opcode & 0x0F00) >> 8 // get X
			reg2 := (opcode & 0x00F0) >> 4 // get Y
			sum := uint16(c.V[reg1]) + uint16(c.V[reg2])
			c.V[0xF] = 0 // reset carry flag
			if sum > 0xFF {
				c.V[0xF] = 1 // set carry flag
			}
			c.V[reg1] = byte(sum)
			c.PC += 2 // next instruction
			break
		case 0x0005:
			// Subtract VY from VX
			reg1 := (opcode & 0x0F00) >> 8 // get X
			reg2 := (opcode & 0x00F0) >> 4 // get Y
			c.V[0xF] = 0                   // reset borrow flag
			if c.V[reg1] > c.V[reg2] {
				c.V[0xF] = 1 // set borrow flag
			}
			c.V[reg1] -= c.V[reg2]
			c.PC += 2 // next instruction
			break
		case 0x0006:
			// Shift VX right by one
			reg := (opcode & 0x0F00) >> 8 // get X
			c.V[0xF] = c.V[reg] & 0x1     // store least significant bit in VF
			c.V[reg] >>= 1
			c.PC += 2 // next instruction
			break
		case 0x0007:
			// Set VX to VY minus VX
			reg1 := (opcode & 0x0F00) >> 8 // get X
			reg2 := (opcode & 0x00F0) >> 4 // get Y
			c.V[0xF] = 0                   // reset borrow flag
			if c.V[reg2] > c.V[reg1] {
				c.V[0xF] = 1 // set borrow flag
			}
			c.V[reg1] = c.V[reg2] - c.V[reg1]
			c.PC += 2 // next instruction
			break
		case 0x000E:
			// Shift VX left by one
			reg := (opcode & 0x0F00) >> 8     // get X
			c.V[0xF] = (c.V[reg] & 0x80) >> 7 // store most significant bit in VF
			c.V[reg] <<= 1
			c.PC += 2 // next instruction
			break

		default:
			printf.Printf("Unknown opcode: 0x%X\n", opcode)
			break
		}
	case 0x9000:
		// Skip next instruction if VX doesn't equal VY
		reg1 := (opcode & 0x0F00) >> 8 // get X
		reg2 := (opcode & 0x00F0) >> 4 // get Y
		c.PC += 2                      // next instruction
		if c.V[reg1] != c.V[reg2] {
			c.PC += 2 // skip this instruction
		}
		break
	case 0xA000:
		// Set I to the address NNN
		c.I = opcode & 0x0FFF
		c.PC += 2 // next instruction
		break
	case 0xB000:
		// Jump to the address NNN plus V0
		c.PC = (opcode & 0x0FFF) + uint16(c.V[0])
		break
	case 0xC000:
		// Set VX to a random number and NN
		val := opcode & 0x00FF        // get lower 8 bits
		reg := (opcode & 0x0F00) >> 8 // get X
		c.V[reg] = byte(rand.Intn(256)) & byte(val)
		c.PC += 2 // next instruction
		break
	case 0xD000:
		// Draw a sprite at coordinate (VX, VY) that has a width of 8 pixels and a height of N pixels
		x := c.V[(opcode&0x0F00)>>8]
		y := c.V[(opcode&0x00F0)>>4]
		height := opcode & 0x000F
		c.V[0xF] = 0
		for yline := uint16(0); yline < height; yline++ {
			pixel, _ := c.RAM.ReadByte(c.I + yline)
			for xline := uint16(0); xline < 8; xline++ {
				if (pixel & (0x80 >> xline)) != 0 {
					if c.Display.IsPixelOn(int(uint16(x)+xline), int(uint16(y)+yline)) { // there is a check for 1 here, not sure if this is correct
						c.V[0xF] = 1
					}
				}
			}
		}
		c.PC += 2 // next instruction
		break
	case 0xE000:
		switch opcode & 0x00FF {
		case 0x009E:
			// Skip next instruction if the key stored in VX is pressed
			reg := (opcode & 0x0F00) >> 8 // get X
			if c.Input.IsKeyPressed(c.V[reg]) {
				c.PC += 4 // skip next instruction
			} else {
				c.PC += 2 // next instruction
			}
			break
		case 0x00A1:
			// Skip next instruction if the key stored in VX isn't pressed
			reg := (opcode & 0x0F00) >> 8 // get X
			if !c.Input.IsKeyPressed(c.V[reg]) {
				c.PC += 4 // skip next instruction
			} else {
				c.PC += 2 // next instruction
			}
			break

		default:
			printf.Printf("Unknown opcode: 0x%X\n", opcode)
			break
		}
		break
	case 0xF000:
		switch opcode & 0x00FF {
		case 0x0007:
			// Set VX to the value of the delay timer
			reg := (opcode & 0x0F00) >> 8 // get X
			c.V[reg] = c.DT
			c.PC += 2 // next instruction
			break
		case 0x000A:
			// Wait for a key press and store the result in VX
			reg := (opcode & 0x0F00) >> 8 // get X
			key := c.Input.WaitForKeyPress()
			c.V[reg] = key
			c.PC += 2 // next instruction
			break
		case 0x0015:
			// Set the delay timer to VX
			reg := (opcode & 0x0F00) >> 8 // get X
			c.DT = c.V[reg]
			c.PC += 2 // next instruction
			break
		case 0x0018:
			// Set the sound timer to VX
			reg := (opcode & 0x0F00) >> 8 // get X
			c.ST = c.V[reg]
			c.PC += 2 // next instruction
			break
		case 0x001E:
			// Add VX to I
			reg := (opcode & 0x0F00) >> 8 // get X
			c.I += uint16(c.V[reg])
			c.PC += 2 // next instruction
			break
		case 0x0029:
			// Set I to the location of the sprite for the character in VX
			reg := (opcode & 0x0F00) >> 8 // get X
			c.I = uint16(c.V[reg]) * 5    // each character is 5 bytes long
			c.PC += 2                     // next instruction
			break
		case 0x0033:
			// Store the binary-coded decimal representation of VX at the addresses I, I+1, and I+2
			reg := (opcode & 0x0F00) >> 8 // get X
			value := c.V[reg]
			c.RAM.WriteByte(c.I, value/100)
			c.RAM.WriteByte(c.I+1, (value/10)%10)
			c.RAM.WriteByte(c.I+2, (value%100)%10)
			c.PC += 2 // next instruction
			break
		case 0x0055:
			// Store V0 to VX in memory starting at address I
			reg := (opcode & 0x0F00) >> 8 // get X
			for i := uint16(0); i <= reg; i++ {
				c.RAM.WriteByte(c.I+i, c.V[i])
			}
			c.PC += 2 // next instruction
			break
		case 0x0065:
			// Fill V0 to VX with values from memory starting at address I
			reg := (opcode & 0x0F00) >> 8 // get X
			for i := uint16(0); i <= reg; i++ {
				value, _ := c.RAM.ReadByte(c.I + i)
				c.V[i] = value
			}
			c.I = c.I + reg + 1
			c.PC += 2 // next instruction
			break

		default:
			printf.Printf("Unknown opcode: 0x%X\n", opcode)
			break
		}

	default:
		printf.Printf("Unknown opcode: 0x%X\n", opcode)
	}

	// timer update
}

// func (c *CPU) execute() {
// 	// Execute opcode
// }

func printState(c *CPU) {
	printf.Printf("PC: 0x%X\n", c.PC)
	printf.Printf("I: 0x%X\n", c.I)
	printf.Printf("V: %v\n", c.V)
	printf.Printf("DT: 0x%X\n", c.DT)
	printf.Printf("ST: 0x%X\n", c.ST)
}

func (c *CPU) Cycle(verbose bool, RAM *memory.Memory) {

	opcode := c.fetch()
	printf.Printf("Opcode: 0x%X\n", opcode)
	if verbose {
		printState(c)
	}

	c.decodeAndExecute(opcode)
}
