package cpu

type CPU struct {
	V  [16]byte // General purpose registers (V0 to VF)
	I  uint16   // Index register
	PC uint16   // Program counter
	SP byte     // Stack pointer
	DT byte     // Delay timer
	ST byte     // Sound timer
}

func NewCPU() *CPU {
	return &CPU{
		PC: 0x200, // Program counter starts at 0x200
	}
}

// func (CPU) InitCPU() CPU {
// 	return CPU{}
// }
