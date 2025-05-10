package chip8

import "testing"

func TestDisplayClear(t *testing.T) {
	d := NewDisplay()
	// Set some pixels to true
	d.Pixels[0][0] = true
	d.Pixels[DisplayHeight-1][DisplayWidth-1] = true
	d.Clear()
	for y := 0; y < DisplayHeight; y++ {
		for x := 0; x < DisplayWidth; x++ {
			if d.Pixels[y][x] {
				t.Errorf("Pixel at (%d,%d) not cleared", y, x)
			}
		}
	}
}
