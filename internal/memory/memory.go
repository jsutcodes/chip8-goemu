package memory

import (
	"fmt"
	"os"
)

// 4kb of RAM
const MemorySize = 4096

type Memory struct {
	size  int
	bytes [MemorySize]byte
}

func (m *Memory) PrintMemoryToFile(filename string) error {
	file, err := os.Create(filename) // os.Create will always overwrite the file
	if err != nil {
		return err
	}
	defer file.Close()

	// Print memory contents to file in the specified format
	for i := 0; i < MemorySize; i += 16 {
		_, err := fmt.Fprintf(file, "%04x: ", i)
		if err != nil {
			return err
		}
		for j := 0; j < 16; j++ {
			if i+j < MemorySize {
				_, err := fmt.Fprintf(file, "%02x", m.bytes[i+j])
				if err != nil {
					return err
				}
				if j%2 == 1 {
					_, err := fmt.Fprint(file, " ")
					if err != nil {
						return err
					}
				}
			} else {
				_, err := fmt.Fprint(file, "   ")
				if err != nil {
					return err
				}
			}
		}
		_, err = fmt.Fprint(file, " ")
		if err != nil {
			return err
		}
		for j := 0; j < 16; j++ {
			if i+j < MemorySize {
				b := m.bytes[i+j]
				if b >= 32 && b <= 126 {
					_, err := fmt.Fprintf(file, "%c", b)
					if err != nil {
						return err
					}
				} else {
					_, err := fmt.Fprint(file, ".")
					if err != nil {
						return err
					}
				}
			}
		}
		_, err = fmt.Fprintln(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Memory) ReadByte(address uint16) (byte, error) {
	if address >= MemorySize {
		panic("address out of bounds")
	}
	return m.bytes[address], nil
}

func (m *Memory) WriteByte(address uint16, value byte) {
	if address >= MemorySize {
		panic("address out of bounds")
	}
	m.bytes[address] = value
}
func (m *Memory) LoadROM(data []byte) {
	fmt.Printf("ROM size: %d\n", len(data))
	if len(data) > MemorySize-0x200 {
		panic("ROM too large")
	}
	for i, b := range data {
		m.bytes[i+0x200] = b
	}
}

func NewMemory() *Memory {
	return &Memory{
		size:  MemorySize,
		bytes: [MemorySize]byte{},
	}
}
