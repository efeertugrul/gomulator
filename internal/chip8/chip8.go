package chip8

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type Chip8 struct {
	CPU     *CPU
	Memory  *Memory
	Display *Display
	Input   *Input
}

func New() *Chip8 {
	chip8 := &Chip8{
		CPU:     NewCPU(),
		Memory:  NewMemory(),
		Display: NewDisplay(),
		Input:   &Input{},
	}

	return chip8
}

func (c *Chip8) InitDisplay() {
	var err error
	sdl.Init(sdl.INIT_VIDEO)

	c.Display.sdl, err = newSDLRenderer()
	if err != nil {
		log.Fatalf("SDL Init failed: %v", err)
	}
}

func (c *Chip8) Initialize() error {
	c.InitDisplay()

	return nil
}

func (c *Chip8) UpdateTimers() {
	if c.CPU.DelayTimer > 0 {
		c.CPU.DelayTimer--
	}
	if c.CPU.SoundTimer > 0 {
		c.CPU.SoundTimer--
		c.PlaySound()
	}
}

func (c *Chip8) Cleanup() {
}

func (c *Chip8) Render() error {
	c.Display.Render()

	return nil
}
