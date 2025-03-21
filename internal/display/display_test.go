package display

import (
	"fmt"
	"strings"
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

	outputStr := ""
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if display.IsPixelOn(x, y) {
				outputStr += "█"
			} else {
				outputStr += " "
			}
		}
		outputStr += "\n"
	}

	if outputStr != expectedOutput {
		t.Errorf("Expected array:\n%v\nGot array:\n%v", display.pixels, getPixelsArray(outputStr))
	}
}

func getPixelsArray(outputStr string) []bool {
	lines := strings.Split(outputStr, "\n")
	pixels := make([]bool, height*width)
	for y, line := range lines {
		if y >= len(lines)-1 {
			break
		}
		for x, char := range line {
			pixels[y*width+x] = char == '█'
		}
	}
	return pixels
}
