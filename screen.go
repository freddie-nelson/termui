package termui

import (
	"fmt"
	"os"
	"sort"
	"time"
)

var (
	drawTicker     *time.Ticker
	stopDrawTicker chan bool
)

// Screen handles rendering of ui
type Screen struct {
	frontBuffer     [][]Color
	backBuffer      [][]Color
	scrollY         int
	scrollX         int
	visibleElements []*Container
	out             *os.File
}

// SwapBuffers swaps the pointers of the front and back buffer
func (s *Screen) SwapBuffers() {
	temp := s.frontBuffer
	s.frontBuffer = s.backBuffer
	s.backBuffer = temp
}

// ClearBuffer clears either the front or back buffer
func (s *Screen) ClearBuffer(front bool) {
	width, height := ENV.GetSize()

	buffer := make([][]Color, width)
	for x, _ := range buffer {
		buffer[x] = make([]Color, height)
	}

	if front {
		s.frontBuffer = buffer
	} else {
		s.backBuffer = buffer
	}
}

// FindVisibleElements finds all elements that are currently on the screen
func (s *Screen) FindVisibleElements(parent *Container) {
	visibleVertical := (parent.y >= s.scrollY || parent.y+parent.height >= s.scrollY) && parent.y < s.scrollY+ENV.height
	visibleHorizontal := (parent.x >= s.scrollX || parent.x+parent.width >= s.scrollX) && parent.x < s.scrollX+ENV.width
	if visibleVertical && visibleHorizontal {
		s.visibleElements = append(s.visibleElements, parent)
	}

	if len(parent.children) > 0 {
		for _, c := range parent.children {
			s.FindVisibleElements(c)
		}
	}

	// fmt.Printf("\r %v %v", ENV.width, ENV.height)
}

// SortByYPos sorts elements by y coord ascending order
func (s *Screen) SortByYPos(elements []*Container) {
	sort.SliceStable(elements, func(i, j int) bool {
		return elements[i].y < elements[j].y
	})
}

// BufferVisibleElements renders all visible elements into the back buffer
func (s *Screen) BufferVisibleElements() {
	s.visibleElements = make([]*Container, 0)
	for _, c := range WINDOW.children {
		s.FindVisibleElements(c)
	}

	s.SortByYPos(s.visibleElements)

	for _, element := range s.visibleElements {

		for i := 0; i < element.width; i++ {
			for j := 0; j < element.height; j++ {
				if s.backBuffer[element.x+i] != nil && len(s.backBuffer[element.x+i]) > element.y+j {
					s.backBuffer[element.x+i][element.y+j] = element.bgColor
				}
			}
		}
	}
}

// DrawFrame swaps buffers and draws front buffer to terminal
func (s *Screen) DrawFrame() {
	s.SwapBuffers()

	output := ""
	for y := 0; y < ENV.height; y++ {
		last := NewColor(-1, -1, -1)

		for x := 0; x < ENV.width; x++ {
			c := s.frontBuffer[x][y]
			if c.r != last.r || c.g != last.g || c.b != last.b {
				output += fmt.Sprintf("%s ", c.ToANSII(false))
			} else {
				output += " "
			}
		}
	}

	SetCursorPos(0, 0)
	HideCursor()
	s.out.WriteString(output)
}

// StartWatcher starts the environment watcher ticker
func (s *Screen) StartDrawLoop() {
	drawTicker = time.NewTicker(time.Duration(DesiredDelta) * time.Millisecond)
	stopDrawTicker = make(chan bool)

	go func() {
		for {
			// calculate delta time
			now := int(time.Now().UnixNano())
			Delta = now - LastFrameTime
			LastFrameTime = now

			select {
			case <-stopDrawTicker:
				return
			case <-drawTicker.C:
				s.ClearBuffer(false)
				s.BufferVisibleElements()
				s.DrawFrame()
			}
		}
	}()
}

// StopWatcher stops the environment watcher ticker
func (s *Screen) StopDrawLoop() {
	drawTicker.Stop()
	stopDrawTicker <- true
}

// NewScreen returns a pointer to a new screen
func NewScreen() *Screen {
	return &Screen{out: os.Stdout}
}
