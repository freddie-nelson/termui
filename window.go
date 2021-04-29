package termui

// Window top level application container
type Window struct {
	*Container
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
