package display

import (
	"fmt"
)

const (
	width  = 64
	height = 32
	fps    = 60
)

type Display struct {
	pixels *[width * height]bool
}

func (d *Display) Clear() {
	for i := range d.pixels {
		d.pixels[i] = false
	}
}

func (d *Display) SetPixel(x, y int, on bool) {
	if x >= 0 && x < width && y >= 0 && y < height {
		d.pixels[y*width+x] = on
	}
}

func (d *Display) Render() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if d.pixels[y*width+x] {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func NewDisplay() *Display {
	return &Display{
		pixels: &[width * height]bool{},
	}
}

// func main() {
// 	pixels := [width * height]bool{}
// 	display := &Display{pixels: &pixels}
// 	ticker := time.NewTicker(time.Second / fps)
// 	defer ticker.Stop()

// 	for range ticker.C {
// 		display.Clear()
// 		// Example: Turn on some pixels
// 		// display.SetPixel(10, 10, true)
// 		// display.SetPixel(20, 20, true)
// 		// display.SetPixel(30, 30, true)

// 		display.Render()
// 		fmt.Print("\033[H\033[2J") // Clear the console
// 	}
// }
