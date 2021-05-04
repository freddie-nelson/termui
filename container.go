package termui

// Container base container
type Container struct {
	x        int
	y        int
	width    int
	height   int
	zIndex   int
	color    Color
	bgColor  Color
	children []*Container
}

func (c *Container) Tick() {}

// Size returns width and height of container
func (c *Container) Size() (int, int) {
	return c.width, c.height
}

// Position returns x and y position of container
func (c *Container) Position() (int, int) {
	return c.x, c.y
}

// SetZIndex sets the z index of the container
func (c *Container) SetZIndex(zIndex int) {
	c.zIndex = zIndex
}

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
	return &Container{x, y, width, height, 0, color, bgColor, make([]*Container, 0)}
}
