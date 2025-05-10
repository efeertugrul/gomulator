package chip8

import (
	"context"

	"github.com/efeertugrul/gomulator/pkg/emulator"
	"github.com/veandco/go-sdl2/sdl"
)

// Input handles keyboard input for the CHIP-8 emulator
type Input struct {
	Keys [16]bool
}

// Update processes SDL events and updates the key states
func (i *Input) Update() error {
	for j := range i.Keys {
		i.Keys[j] = false
	}

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			return emulator.ErrQuit
		case *sdl.KeyboardEvent:
			if key := i.mapKey(e.Keysym.Scancode); key >= 0 {
				i.Keys[key] = e.Type == sdl.KEYDOWN
			}
		}
	}
	return nil
}

// mapKey maps SDL scancodes to CHIP-8 keys
func (i *Input) mapKey(code sdl.Scancode) int {
	switch code {
	case sdl.SCANCODE_1:
		return 0x0
	case sdl.SCANCODE_2:
		return 0x1
	case sdl.SCANCODE_3:
		return 0x2
	case sdl.SCANCODE_4:
		return 0x3
	case sdl.SCANCODE_Q:
		return 0x4
	case sdl.SCANCODE_W:
		return 0x5
	case sdl.SCANCODE_E:
		return 0x6
	case sdl.SCANCODE_R:
		return 0x7
	case sdl.SCANCODE_A:
		return 0x8
	case sdl.SCANCODE_S:
		return 0x9
	case sdl.SCANCODE_D:
		return 0xa
	case sdl.SCANCODE_F:
		return 0xb
	case sdl.SCANCODE_Z:
		return 0xc
	case sdl.SCANCODE_X:
		return 0xd
	case sdl.SCANCODE_C:
		return 0xe
	case sdl.SCANCODE_V:
		return 0xF
	default:
		return -1
	}
}

func (c *Chip8) HandleInput(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return c.Input.Update()
}
