package chip8

import (
	"log/slog"
)

func (c *Chip8) PlaySound() {
	if c.CPU.SoundTimer > 0 {
		slog.Default().Info("Beep!")
	}
}
