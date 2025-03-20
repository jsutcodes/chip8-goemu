package input

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Keypad represents the CHIP-8 keypad which has a 4x4 grid of keys:
//
//	chip8  vs. sdl
//
// ========|========
// 1 2 3 C | 1 2 3 4
// 4 5 6 D | q w e r
// 7 8 9 E | a s d f
// A 0 B F | z x c v
//
// The Keypad struct contains a boolean array to keep track of the state (pressed or not pressed) of each key.
// The keys are mapped to the following SDL keys:
//
// CHIP-8 Key | SDL Key
// -----------|--------
// 1          | 1
// 2          | 2
// 3          | 3
// C          | 4
// 4          | Q
// 5          | W
// 6          | E
// D          | R
// 7          | A
// 8          | S
// 9          | D
// E          | F
// A          | Z
// 0          | X
// B          | C
// F          | V
type Keypad struct {
	keys [16]bool
}

func NewKeypad() *Keypad {
	return &Keypad{}
}

func (k *Keypad) IsKeyPressed(key uint8) bool {
	if key >= 16 {
		return false
	}
	return k.keys[key]
}

func (k *Keypad) SetKeyPressed(key uint8, pressed bool) {
	if key < 16 {
		k.keys[key] = pressed
	}
}

func (k *Keypad) HandleEvent(event sdl.Event) {
	switch e := event.(type) {
	case *sdl.KeyboardEvent:
		if e.Type == sdl.KEYDOWN || e.Type == sdl.KEYUP {
			pressed := e.Type == sdl.KEYDOWN
			switch e.Keysym.Sym {
			case sdl.K_1:
				k.SetKeyPressed(0x1, pressed)
			case sdl.K_2:
				k.SetKeyPressed(0x2, pressed)
			case sdl.K_3:
				k.SetKeyPressed(0x3, pressed)
			case sdl.K_4:
				k.SetKeyPressed(0xC, pressed)
			case sdl.K_q:
				k.SetKeyPressed(0x4, pressed)
			case sdl.K_w:
				k.SetKeyPressed(0x5, pressed)
			case sdl.K_e:
				k.SetKeyPressed(0x6, pressed)
			case sdl.K_r:
				k.SetKeyPressed(0xD, pressed)
			case sdl.K_a:
				k.SetKeyPressed(0x7, pressed)
			case sdl.K_s:
				k.SetKeyPressed(0x8, pressed)
			case sdl.K_d:
				k.SetKeyPressed(0x9, pressed)
			case sdl.K_f:
				k.SetKeyPressed(0xE, pressed)
			case sdl.K_z:
				k.SetKeyPressed(0xA, pressed)
			case sdl.K_x:
				k.SetKeyPressed(0x0, pressed)
			case sdl.K_c:
				k.SetKeyPressed(0xB, pressed)
			case sdl.K_v:
				k.SetKeyPressed(0xF, pressed)
			}
		}
	}
}
