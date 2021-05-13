package termui

import "time"

type Event struct {
	time time.Time
}

type KeyEvent struct {
	Event
	key rune
}

type MouseEvent struct {
	Event
	mouseX   int
	mouseY   int
	button   int
	modifier int
}

type ScrollEvent struct {
	Event
	// direction:
	// 0 = up,
	// 1 = down,
	// 2 = left,
	// 3 = right
	direction int
}
