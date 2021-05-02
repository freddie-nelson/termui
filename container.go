package termui

// Container base container
type Container struct {
	x        int
	y        int
	width    int
	height   int
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
	return &Container{x, y, width, height, color, bgColor, make([]*Container, 0)}
}
