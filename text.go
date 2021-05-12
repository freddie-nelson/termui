package termui

type Text struct {
	Container
	text       string
	color      Color
	charsDrawn int
}

// Colors returns the color and bgColor of the container
func (t *Text) Colors() (Color, Color) {
	return t.color, t.bgColor
}

// CharAt returns the character at the relative x and y of the container
func (t *Text) CharAt(x, y int) rune {
	// if x and y are 0 then element is at start of drawing so reset charsDrawn
	if x == 0 && y == 0 {
		t.charsDrawn = 0
	}

	// return blank char if done writing text or position is inside padding area
	if t.charsDrawn == len(t.text) || x < t.padding || x >= t.width-t.padding || y < t.padding || y >= t.height-t.padding {
		return ' '
	}

	char := rune(t.text[t.charsDrawn])
	t.charsDrawn++

	return char
}

// NewText returns a pointer to a new text container
func NewText(x, y, width, height, padding int, color, bgColor Color, text string) *Text {
	return &Text{*NewContainer(x, y, width, height, padding, bgColor).SetFocusable(true), text, color, 0}
}
