package termui

// Page represents a single route in the ui
type Page struct {
	Container
	resizeToWindow   bool
	isElementFocused bool
	focusedElement   *Element
}

// Tick executes every tick
// keeps page the same size as the window if resizeToWindow is true
func (p *Page) Tick() {
	if !p.resizeToWindow {
		return
	}

	width, height := WINDOW.width, WINDOW.height

	if width != p.width {
		p.width = width
	}
	if height != p.height {
		p.height = height
	}
}

// ResizeToWindow sets if the page will resize to fill the window or not
func (p *Page) ResizeToWindow(resize bool) {
	p.resizeToWindow = resize
}

// SetFocusedElement sets the currently focused element
func (p *Page) SetFocusedElement(e *Element) { // TODO rethink/rework focusing elements (needs research)
	if e == nil {
		p.isElementFocused = false
		p.focusedElement = nil
	} else {
		p.isElementFocused = true
		p.focusedElement = e
	}
}

// NewPage returns a pointer to a new page
func NewPage(color, bgColor Color) *Page {
	return &Page{Container: *NewContainer(0, 0, WINDOW.width, WINDOW.height, 0, bgColor), resizeToWindow: true}
}
