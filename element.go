package termui

type Element interface {
	Id() string
	Size() (int, int)
	SetSize(width, height int) *Container
	Position() (int, int, int)
	SetPosition(x, y int) *Container
	SetZIndex(zIndex int) *Container
	IsFocusable() bool
	SetFocusable(focusable bool) *Container
	Colors() (Color, Color)
	CharAt(x, y int) rune
	Children() []Element
	AddChild(child ...Element) *Container
}
