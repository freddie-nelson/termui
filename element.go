package termui

type Element interface {
	Id() string
	Size() (int, int)
	SetSize(width, height int)
	Position() (int, int, int)
	SetPosition(x, y int)
	SetZIndex(zIndex int)
	IsFocusable() bool
	SetFocusable(focusable bool)
	Colors() (Color, Color)
	CharAt(x, y int) rune
	Children() []Element
	AddChild(child ...Element)
}
