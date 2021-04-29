package termui

import "runtime"

// Framerate the desired fps of the ui
const Framerate int = 60

// DesiredDelta the desired time between frames in milliseconds
const DesiredDelta int = 1000 / 60

var ENV *Environment
var WINDOW *Window

func Create() {
	// init environment
	ENV = NewEnvironment()
	ENV.StartWatcher()

	runtime.Goexit()
}
