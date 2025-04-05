package chip8

import "github.com/veandco/go-sdl2/sdl"

const (
	DisplayWidth  = 64
	DisplayHeight = 32
)

type Display struct {
	Pixels [DisplayHeight][DisplayWidth]bool
	sdl    *SDLRenderer
}

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
