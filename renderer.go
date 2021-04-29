// Renderer handles rendering of all elements on the screen
type Renderer struct {
	frontBuffer *[][]Color
	backBuffer  *[][]Color
	elements    *[][]Container
}