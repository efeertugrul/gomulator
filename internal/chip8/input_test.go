package chip8

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

func TestInputMapKey(t *testing.T) {
	input := &Input{}
	tests := []struct {
		scancode sdl.Scancode
		expect   int
	}{
		{sdl.SCANCODE_1, 0x0},
		{sdl.SCANCODE_2, 0x1},
		{sdl.SCANCODE_3, 0x2},
		{sdl.SCANCODE_4, 0x3},
		{sdl.SCANCODE_Q, 0x4},
		{sdl.SCANCODE_W, 0x5},
		{sdl.SCANCODE_E, 0x6},
		{sdl.SCANCODE_R, 0x7},
		{sdl.SCANCODE_A, 0x8},
		{sdl.SCANCODE_S, 0x9},
		{sdl.SCANCODE_D, 0xA},
		{sdl.SCANCODE_F, 0xB},
		{sdl.SCANCODE_Z, 0xC},
		{sdl.SCANCODE_X, 0xD},
		{sdl.SCANCODE_C, 0xE},
		{sdl.SCANCODE_V, 0xF},
		{sdl.SCANCODE_B, -1}, // unmapped
	}
	for _, tt := range tests {
		if got := input.mapKey(tt.scancode); got != tt.expect {
			t.Errorf("mapKey(%v) = %d; want %d", tt.scancode, got, tt.expect)
		}
	}
}
