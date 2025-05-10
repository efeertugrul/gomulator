package chip8

import (
	"context"
	"errors"
	"os"
)

func (c *Chip8) LoadROM(ctx context.Context, path string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
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
