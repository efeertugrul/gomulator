package chip8

import (
	"github.com/efeertugrul/gomulator/pkg/emulator"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	DisplayWidth  = 64
	DisplayHeight = 32
)

type Display struct {
	Pixels [DisplayHeight][DisplayWidth]bool
	sdl    *SDLRenderer
}

var _ emulator.Display = (*Display)(nil)

func NewDisplay() *Display {
	return &Display{}
}

func (d *Display) Clear() {
	for y := range DisplayHeight {
		for x := range DisplayWidth {
			d.Pixels[y][x] = false
		}
	}
}

func (d *Display) Render() {
	d.sdl.renderer.SetDrawColor(0, 0, 0, 255) // Clear: black
	d.sdl.renderer.Clear()
	d.sdl.renderer.SetDrawColor(0, 255, 0, 255) // Pixel color: green

	for y := range DisplayHeight {
		for x := range DisplayWidth {
			if d.Pixels[y][x] {
				rect := sdl.Rect{
					X: int32(x * windowScale),
					Y: int32(y * windowScale),
					W: windowScale,
					H: windowScale,
				}
				d.sdl.renderer.FillRect(&rect)
			}
		}
	}

	d.sdl.renderer.Present()
}

func (d *Display) Cleanup() {
	if d.sdl != nil {
		d.sdl.Cleanup()
	}
}

func (d *Display) SetPixel(x, y uint8, value bool) bool {
	// Wrap around if out of bounds (CHIP-8 display wraps)
	x %= DisplayWidth
	y %= DisplayHeight
	prev := d.Pixels[y][x]
	d.Pixels[y][x] = value
	return prev
}

func (d *Display) Update() error {
	d.Render()
	return nil
}
