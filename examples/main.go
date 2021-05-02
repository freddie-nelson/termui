package main

import (
	"github.com/freddie-nelson/termui"
)

func main() {
	// colors
	green := termui.NewColor(50, 255, 100)
	blue := termui.NewColor(50, 100, 255)
	red := termui.NewColor(255, 100, 50)
	black := termui.NewColor(0, 0, 0)

	ui := termui.Create(black, black)

	c := termui.NewContainer(0, 0, 10, 5, green, green)
	c2 := termui.NewContainer(3, 3, 10, 5, blue, blue)
	c3 := termui.NewContainer(6, 6, 10, 5, red, red)
	c4 := termui.NewContainer(12, 0, 10, 5, red, red)
	ui.AddChild(c, c2, c3, c4)

	ui.Start()
}
