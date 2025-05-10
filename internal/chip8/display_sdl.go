package chip8

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowScale = 10
	windowTitle = "CHIP-8 Emulator"
)

type SDLRenderer struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

func newSDLRenderer() (*SDLRenderer, error) {
	window, err := sdl.CreateWindow(
		windowTitle,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		DisplayWidth*windowScale,
		DisplayHeight*windowScale,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		return nil, err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		window.Destroy()
		return nil, err
	}

	return &SDLRenderer{window: window, renderer: renderer}, nil
}

func (r *SDLRenderer) Cleanup() {
	if r.renderer != nil {
		r.renderer.Destroy()
	}
	if r.window != nil {
		r.window.Destroy()
	}
}
