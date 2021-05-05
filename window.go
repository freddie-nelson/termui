package termui

import (
	"runtime"
)

// Window top level application container
type Window struct {
	width  int
	height int
	page   *Page
}

// SetPage sets the current page
func (w *Window) SetPage(p *Page) {
	w.page = p
}

// Tick executes every tick
// keeps window the same size as the terminal
func (w *Window) Tick() {
	width, height := ENV.GetSize()

	if width != w.width {
		w.width = width
	}
	if height != w.height {
		w.height = height
	}
}

// Start starts draw, tick and environment loops
func (w *Window) Start() {
	ClearTerminal()

	ENV.StartWatcher()
	SCREEN.StartDrawLoop()
	runtime.Goexit()
}
