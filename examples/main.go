package main

import (
	"github.com/freddie-nelson/termui"
)

func main() {
	// colors
	green := termui.NewColor(50, 255, 100)
	blue := termui.NewColor(50, 100, 255)
	red := termui.NewColor(255, 100, 50)
	bg := termui.NewColor(10, 10, 10)

	// ui
	ui := termui.Create()

	// home page
	homePage := termui.NewPage(bg, bg)

	c := termui.NewContainer(0, 0, 10, 12, green, green)
	c2 := termui.NewContainer(3, 3, 10, 5, blue, blue)
	c2.SetZIndex(1)
	c3 := termui.NewContainer(6, 6, 10, 5, red, red)
	c4 := termui.NewContainer(12, 0, 10, 5, red, red)
	c4.SetZIndex(2)
	c5 := termui.NewContainer(17, 3, 10, 5, blue, blue)
	c4.AddChild(c5)

	homePage.AddChild(c, c2, c3, c4)
	ui.SetPage(homePage)

	ui.Start()
}
