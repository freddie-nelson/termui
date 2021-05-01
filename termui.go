package termui

import (
	"os"
	"os/exec"
	"runtime"
	"time"
)

// Framerate the desired fps of the ui
const Framerate int = 60

// DesiredDelta the desired time between frames in milliseconds
const DesiredDelta int = 1000 / 60

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

// ClearTerminal uses platform specific commands to clear terminal
// TODO look for better implementation of this
func ClearTerminal() {
	clear := make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
