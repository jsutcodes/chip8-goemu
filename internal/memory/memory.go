package memory

// 4kb of RAM
const MemorySize = 4096

type RAM struct {
	size  int
	bytes [MemorySize]byte
}

func (RAM) initMemory() RAM {
	return RAM{}
}
