package input

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

func TestKeypad_IsKeyPressed(t *testing.T) {
	keypad := NewKeypad()

	if keypad.IsKeyPressed(0x2) {
		t.Errorf("Expected key 0x2 to be not pressed")
	}

	keypad.SetKeyPressed(0x2, true)

	if !keypad.IsKeyPressed(0x2) {
		t.Errorf("Expected key 0x2 to be pressed")
	}
}

func TestKeypad_SetKeyPressed(t *testing.T) {
	keypad := NewKeypad()

	keypad.SetKeyPressed(0x2, true)
	if !keypad.IsKeyPressed(0x2) {
		t.Errorf("Expected key 0x2 to be pressed")
	}

	keypad.SetKeyPressed(0x2, false)
	if keypad.IsKeyPressed(0x2) {
		t.Errorf("Expected key 0x2 to be not pressed")
	}
}

func TestKeypad_HandleEvent(t *testing.T) {
	keypad := NewKeypad()

	event := sdl.KeyboardEvent{
		Type: sdl.KEYDOWN,
		Keysym: sdl.Keysym{
			Sym: sdl.K_2,
		},
	}
	keypad.HandleEvent(&event)

	if !keypad.IsKeyPressed(0x2) {
		t.Errorf("Expected key 0x2 to be pressed after KEYDOWN event")
	}

	event.Type = sdl.KEYUP
	keypad.HandleEvent(&event)

	if keypad.IsKeyPressed(0x2) {
		t.Errorf("Expected key 0x2 to be not pressed after KEYUP event")
	}
}
