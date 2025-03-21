package timer

import "time"

type Timer struct {
	delay uint8
	sound uint8
}

func (t *Timer) Update() {
	ticker := time.NewTicker(time.Second / 60)
	defer ticker.Stop()

	for range ticker.C {
		if t.delay > 0 {
			t.delay--
		}
		if t.sound > 0 {
			t.sound--
		}
	}
}

func NewTimer() *Timer {
	return &Timer{
		delay: 0, //?
		sound: 0,
	}
}
