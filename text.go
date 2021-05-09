package termui

type Text struct {
	Container
	text  string
	color Color
}

func (t *Text) Colors() (Color, Color) {
	return t.color, t.bgColor
}

// Colors returns the color and bgColor of the container
func (t *Text) CharAt(x, y int) rune {
	// return blank char if position is on padding
	if x < t.padding || x >= t.width-t.padding || y < t.padding || y >= t.height-t.padding {
		return ' '
	}

	charNum := (x - t.padding) * (y - t.padding + 1)
	if charNum >= len(t.text) {
		return ' '
	}

	return rune(t.text[charNum])
}

// NewText returns a pointer to a new text container
func NewText(x, y, width, height, padding int, color, bgColor Color, text string) *Text {
	return &Text{*NewContainer(x, y, width, height, padding, bgColor), text, color}
}
