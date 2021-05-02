package termui

import (
	"time"
)

// Framerate the desired fps of the ui
const Framerate int = 60

// DesiredDelta the desired time between frames in milliseconds
const DesiredDelta int = 1000 / 120

// Delta is the time since the last frame
var Delta int = 0

// LastFrameTime is the unix time of the last frame in nanoseconds
var LastFrameTime int = int(time.Now().UnixNano())

var ENV *Environment
var SCREEN *Screen
var WINDOW *Window

// Create creates a new termui instance
// color and bgColor are colors of main window
// returns a pointer to the termui window
func Create(color Color, bgColor Color) *Window {
	ENV = NewEnvironment()
	SCREEN = NewScreen()
	WINDOW = &Window{Container: *NewContainer(0, 0, 0, 0, color, bgColor)}

	return WINDOW
}
