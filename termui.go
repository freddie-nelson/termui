package termui

import (
	"os"
	"time"

	"golang.org/x/term"
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
var INPUT *Input

// PreviousState stores the previous state of the terminal before ui started
var PreviousState *term.State

// Create creates a new termui instance
// returns a pointer to the termui window
func Create() *Window {
	ENV = NewEnvironment()
	SCREEN = NewScreen()
	WINDOW = &Window{}
	INPUT = NewInput(os.Stdin)

	PreviousState = SetTermRawMode()

	return WINDOW
}

// Close cleanly exits the ui and restores the previous terminal state
func Close() {
	term.Restore(int(SCREEN.out.Fd()), PreviousState)
	ResetAttributes()
	ClearTerminal()
	ShowCursor()
	SetCursorPos(0, 0)
	os.Exit(0)
}
