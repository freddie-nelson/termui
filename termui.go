package termui

// Framerate the desired fps of the ui
const Framerate int = 60

// DesiredDelta the desired time between frames in milliseconds
const DesiredDelta int = 1000 / 60

var ENV *Environment

func Create() {
	// init environment
	ENV = NewEnvironment()
	ENV.StartWatcher()
}
