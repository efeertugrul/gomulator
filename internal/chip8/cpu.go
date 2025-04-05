package chip8

type CPU struct {
	V          [16]byte // General-purpose registers V0-VF
	I          uint16   // Index register
	PC         uint16   // Program counter
	SP         byte     // Stack pointer
	Stack      [32]uint16
	DelayTimer byte
	SoundTimer byte
}

func NewCPU() *CPU {
	return &CPU{
		PC: 0x200, // CHIP-8 programs start at 0x200
	}
}

func (c *CPU) Reset() {
	c.PC = 0x200
	c.SP = 0
	c.I = 0
	for i := range c.V {
		c.V[i] = 0
	}
	for i := range c.Stack {
		c.Stack[i] = 0
	}
	c.DelayTimer = 0
	c.SoundTimer = 0
}
