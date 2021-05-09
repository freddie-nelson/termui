package termui

import "github.com/google/uuid"

// Container base container
type Container struct {
	x         int
	y         int
	width     int
	height    int
	padding   int
	zIndex    int
	bgColor   Color
	children  []Element
	focusable bool
	id        string
}

func (c *Container) Id() string {
	return c.id
}

// Size returns width and height of container
func (c *Container) Size() (int, int) {
	return c.width, c.height
}

// SetPosition sets the x and y position of the container
func (c *Container) SetSize(width, height int) {
	c.width = width
	c.height = height
}

// Position returns x and y position and zIndex of the container
func (c *Container) Position() (int, int, int) {
	return c.x, c.y, c.zIndex
}

// SetPosition sets the x and y position of the container
func (c *Container) SetPosition(x, y int) {
	c.x = x
	c.y = y
}

// SetZIndex sets the z index of the container
func (c *Container) SetZIndex(zIndex int) {
	c.zIndex = zIndex
}

// Colors returns the color and bgColor of the container
func (c *Container) Colors() (Color, Color) {
	return NewColor(-1, -1, -1), c.bgColor
}

// CharAt returns the character at the position in the container
func (c *Container) CharAt(x, y int) rune {
	return ' '
}

// Children returns a slice of this element's children
func (c *Container) Children() []Element {
	return c.children
}

func (c *Container) AddChild(children ...Element) {
	for _, child := range children {
		if c.Id() == child.Id() {
			// TODO add logger
			return
		}

		c.children = append(c.children, child)
	}
}

// NewContainer returns a pointer to a new base container
func NewContainer(x, y, width, height, padding int, bgColor Color) *Container {
	return &Container{x, y, width, height, padding, 0, bgColor, make([]Element, 0), false, uuid.New().String()}
}
