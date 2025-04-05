package chip8

import "log"

func (c *Chip8) PlaySound() {
	if c.CPU.SoundTimer > 0 {
		log.Print("Beep!")
	}
}
