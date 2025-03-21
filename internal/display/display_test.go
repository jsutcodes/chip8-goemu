package display

import (
	"bytes"
	"fmt"
	"log"
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

// Test failing: need to fix render
func TestRender(t *testing.T) {
	display := NewDisplay()
	display.SetPixel(0, 0, true)
	display.SetPixel(1, 0, true)
	display.SetPixel(2, 0, true)

	expectedOutput := "███" + fmt.Sprintf("%*s", width-3, "") + "\n" + fmt.Sprintf("%*s\n", width, "")
	for y := 1; y < height; y++ {
		expectedOutput += fmt.Sprintf("%*s\n", width, "")
	}

	output := captureOutput(display.Render)

	if output != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, output)
	}
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	old := log.Default()
	log.SetOutput(&buf)
	defer log.SetOutput(old.Writer())
	f()
	return buf.String()
}
