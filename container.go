package termui

// Container base container
type Container struct {
	x      int
	y      int
	width  int
	height int
}

func (c *Container) Render() {}

func (c *Container) Tick() {}
