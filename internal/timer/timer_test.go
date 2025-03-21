package timer

import (
	"testing"
	"time"
)

func TestTimerUpdate(t *testing.T) {
	timer := NewTimer()
	timer.delay = 10
	timer.sound = 10

	go timer.Update()

	time.Sleep(time.Second / 60 * 11)

	if timer.delay != 0 {
		t.Errorf("Expected delay to be 0, got %d", timer.delay)
	}

	if timer.sound != 0 {
		t.Errorf("Expected sound to be 0, got %d", timer.sound)
	}
}

func TestNewTimer(t *testing.T) {
	timer := NewTimer()

	if timer.delay != 0 {
		t.Errorf("Expected initial delay to be 0, got %d", timer.delay)
	}

	if timer.sound != 0 {
		t.Errorf("Expected initial sound to be 0, got %d", timer.sound)
	}
}
