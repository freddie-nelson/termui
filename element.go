package termui

type Element interface {
	Tick()
	Size() (int, int)
	Position() (int, int)
	AddChild(child *Container)
}
