package display

import (
	"testing"
)

func TestClear(t *testing.T) {
	display := NewDisplay()
	display.SetPixel(10, 10, true)
	display.Clear()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if display.IsPixelOn(x, y) {
				t.Errorf("Expected pixel at (%d, %d) to be off", x, y)
			}
		}
	}
}

func TestSetPixel(t *testing.T) {
	display := NewDisplay()
	display.SetPixel(10, 10, true)

	if !display.IsPixelOn(10, 10) {
		t.Errorf("Expected pixel at (10, 10) to be on")
	}

	display.SetPixel(10, 10, false)

	if display.IsPixelOn(10, 10) {
		t.Errorf("Expected pixel at (10, 10) to be off")
	}
}

func TestIsPixelOn(t *testing.T) {
	display := NewDisplay()

	if display.IsPixelOn(10, 10) {
		t.Errorf("Expected pixel at (10, 10) to be off")
	}

	display.SetPixel(10, 10, true)

	if !display.IsPixelOn(10, 10) {
		t.Errorf("Expected pixel at (10, 10) to be on")
	}
}

func TestRender(t *testing.T) {
	display := NewDisplay()
	display.SetPixel(0, 0, true)
	display.SetPixel(1, 0, true)
	display.SetPixel(2, 0, true)

	expectedPixels := [2048]bool{}
	expectedPixels[0] = true
	expectedPixels[1] = true
	expectedPixels[2] = true

	actualPixels := display.Render()

	if *actualPixels != expectedPixels {
		t.Errorf("Expected array:\n%v\nGot array:\n%v", expectedPixels, *actualPixels)
	}
}
