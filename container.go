package termui

// Container base container
type Container struct {
	x         int
	y         int
	width     int
	height    int
	zIndex    int
	color     Color
	bgColor   Color
	children  []*Container
	text      string
	focusable bool
}

func (c *Container) Tick() {}

// Size returns width and height of container
func (c *Container) Size() (int, int) {
	return c.width, c.height
}

// SetPosition sets the x and y position of the container
func (c *Container) SetSize(width, height int) {
	c.width = width
	c.height = height
}

// Position returns x and y position of container
func (c *Container) Position() (int, int) {
	return c.x, c.y
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

func (c *Container) SetText() {} // TODO implement text in containers

func (c *Container) AddChild(children ...*Container) {
	for _, child := range children {
		if c == child {
			// TODO add logger
			return
		}

		c.children = append(c.children, child)
	}
}

// NewContainer returns a pointer to a new base container
func NewContainer(x, y, width, height int, color, bgColor Color) *Container {
	return &Container{x, y, width, height, 0, color, bgColor, make([]*Container, 0), "", false}
}
