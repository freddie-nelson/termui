package termui

// Container base container
type Container struct {
	x       int
	y       int
	width   int
	height  int
	color   Color
	bgColor Color
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
