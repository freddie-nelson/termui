package main

import (
	"github.com/freddie-nelson/termui"
)

func main() {
	// colors
	white := termui.NewColor(255, 255, 255)
	black := termui.NewColor(0, 0, 0)

	ui := termui.Create(black, black)

	c := termui.NewContainer(1, 5, 60, 20, white, white)
	ui.AddChild(c)

	ui.Start()
}
