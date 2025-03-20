package input

import (
	"github.com/veandco/go-sdl2/sdl"
)

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
