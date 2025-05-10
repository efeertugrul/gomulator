package chip8

import "testing"

func TestPlaySoundNoPanic(t *testing.T) {
	c := New()
	c.CPU.SoundTimer = 1
	c.PlaySound()
	c.CPU.SoundTimer = 0
	c.PlaySound()
}
