package termui

import (
	"fmt"
	"sort"
)

// Screen handles rendering of ui
type Screen struct {
	frontBuffer     [][]Color
	backBuffer      [][]Color
	scrollY         int
	scrollX         int
	elements        []*Container
	visibleElements []*Container
}

// SwapBuffers swaps the pointers of the front and back buffer
func (s *Screen) SwapBuffers() {
	temp := s.frontBuffer
	s.frontBuffer = s.backBuffer
	s.backBuffer = temp
}

// FindVisibleElements finds all elements that are currently on the screen
func (s *Screen) FindVisibleElements() {
	for _, element := range s.elements {
		visibleVertical := (element.y >= s.scrollY || element.y+element.height >= s.scrollY) && element.y < s.scrollY+ENV.height
		visibleHorizontal := (element.x >= s.scrollX || element.x+element.width >= s.scrollX) && element.x < s.scrollX+ENV.width
		if visibleVertical && visibleHorizontal {
			s.visibleElements = append(s.visibleElements, element)
		}
	}
}

// SortByYPos sorts elements by y coord ascending order
func (s *Screen) SortByYPos(elements []*Container) {
	sort.Slice(elements, func(i, j int) bool {
		return elements[i].y < elements[j].y
	})
}

// BufferVisibleElements renders all visible elements into the back buffer
func (s *Screen) BufferVisibleElements() {
	s.FindVisibleElements()
	s.SortByYPos(s.visibleElements)

	for _, element := range s.visibleElements {

		for i := 0; i < element.width; i++ {
			for j := 0; j < element.height; j++ {
				s.backBuffer[element.x+i][element.y+j] = element.bgColor
			}
		}
	}
}

func SetCursorPos(col, row int) {
	fmt.Printf("\033[%v;%vH", row, col)
}

// DrawFrame swaps buffers and draws front buffer to terminal
func (s *Screen) DrawFrame() {
	s.SwapBuffers()

	output := ""
	for x := 0; x < len(s.frontBuffer); x++ {
		last := NewColor(-1, -1, -1)

		for y := 0; y < len(s.frontBuffer[x]); y++ {
			c := s.frontBuffer[x][y]
			if c.r != last.r || c.g != last.g || c.b != last.b {
				output += fmt.Sprintf("%s ", c.ToANSII(false))
			} else {
				output += " "
			}
		}
	}

	SetCursorPos(0, 0)
	fmt.Print(output)
}

// NewScreen returns a new screen
func NewScreen() *Screen {
	return &Screen{}
}
