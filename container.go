package termui

import "github.com/google/uuid"

// Container base container
type Container struct {
	x              int
	y              int
	width          int
	height         int
	padding        int
	zIndex         int
	bgColor        Color
	children       []Element
	id             string
	eventListeners map[string]func(e Event)
}

func (c *Container) Id() string {
	return c.id
}

// Size returns width and height of container
func (c *Container) Size() (int, int) {
	return c.width, c.height
}

// SetPosition sets the x and y position of the container
func (c *Container) SetSize(width, height int) *Container {
	c.width = width
	c.height = height
	return c
}

// Position returns x and y position and zIndex of the container
func (c *Container) Position() (int, int, int) {
	return c.x, c.y, c.zIndex
}

// SetPosition sets the x and y position of the container
func (c *Container) SetPosition(x, y int) *Container {
	c.x = x
	c.y = y
	return c
}

// SetZIndex sets the z index of the container
func (c *Container) SetZIndex(zIndex int) *Container {
	c.zIndex = zIndex
	return c
}

// Colors returns the color and bgColor of the container
func (c *Container) Colors() (Color, Color) {
	return NewColor(-1, -1, -1), c.bgColor
}

// CharAt returns the character at the relative x and y of the container
func (c *Container) CharAt(x, y int) rune {
	return ' '
}

// Children returns a slice of this element's children
func (c *Container) Children() []Element {
	return c.children
}

func (c *Container) AddChild(children ...Element) *Container {
	for _, child := range children {
		if c.Id() == child.Id() {
			continue
		}

		c.children = append(c.children, child)
	}

	return c
}

// AddEventListener adds an event listener to the element which
// will execute callback every time the element recieves the given event
// returns a boolean representing if the event was registered successfully or not
//
// valid events: mousedown, mouseup, click, rightclick, mousewheeldown, mousewheeldown, scroll, keydown, keyup, keypress
func (c *Container) AddEventListener(event string, callback func(e Event)) bool {
	validEvents := []string{"mousedown", "mouseup", "click", "rightclick", "mousewheeldown", "mousewheeldown", "scroll", "keydown", "keyup", "keypress"}
	valid := false

	for _, ve := range validEvents {
		if ve == event {
			valid = true
			break
		}
	}

	if !valid || c.eventListeners[event] != nil {
		return false
	}

	c.eventListeners[event] = callback
	return true
}

// NewContainer returns a pointer to a new base container
func NewContainer(x, y, width, height, padding int, bgColor Color) *Container {
	return &Container{x, y, width, height, padding, 0, bgColor, make([]Element, 0), uuid.New().String(), make(map[string]func(e Event))}
}
