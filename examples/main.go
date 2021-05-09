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

	c := termui.NewContainer(0, 0, 10, 12, 0, green)
	c2 := termui.NewContainer(3, 3, 10, 5, 0, blue)
	c2.SetZIndex(1)
	c3 := termui.NewContainer(6, 6, 10, 5, 0, red)
	c4 := termui.NewContainer(12, 0, 10, 5, 0, red)
	c4.SetZIndex(2)
	c5 := termui.NewContainer(17, 3, 10, 5, 0, blue)
	c4.AddChild(c5)

	title := termui.NewText(11, 5, 30, 3, 1, bg, green, "Hello world!")
	title.SetZIndex(1)

	verticalText := termui.NewText(30, 3, 5, 18, 2, bg, red, "Hello world!")
	verticalText.SetZIndex(2)

	homePage.AddChild(c, c2, c3, c4, title, verticalText)
	ui.SetPage(homePage)

	ui.Start()
}
