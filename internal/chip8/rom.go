package chip8

import (
	"errors"
	"os"
)

func (c *Chip8) LoadROM(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if len(data)+0x200 > len(c.Memory.RAM) {
		return errors.New("ROM too large to fit in memory")
	}

	copy(c.Memory.RAM[0x200:], data)

	return nil
}
