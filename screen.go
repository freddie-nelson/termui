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
	frontBuffer [][]Color
	backBuffer  [][]Color
	scrollY     int
	scrollX     int
	out         *os.File
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

// IsElementOnScreen returns true if the given container is currently on screen
func (s *Screen) IsElementOnScreen(c *Container) bool {
	visibleVertical := (c.y >= s.scrollY || c.y+c.height >= s.scrollY) && c.y < s.scrollY+ENV.height
	visibleHorizontal := (c.x >= s.scrollX || c.x+c.width >= s.scrollX) && c.x < s.scrollX+ENV.width
	return visibleVertical && visibleHorizontal
}

// SortElements sorts elements by zIndex ascending, y ascending, x ascending
func (s *Screen) SortElements(es []*Container) {
	sort.SliceStable(es, func(i, j int) bool {
		if es[i].zIndex < es[j].zIndex {
			return true
		} else if es[i].zIndex > es[j].zIndex {
			return false
		}

		if es[i].y < es[j].y {
			return true
		} else if es[i].y > es[j].y {
			return false
		}

		return es[i].x < es[j].x
	})
}

// BufferContainer renders the container to the back buffer
func (s *Screen) BufferContainer(c *Container) {
	for i := 0; i < c.width; i++ {
		for j := 0; j < c.height; j++ {
			if s.backBuffer[c.x+i] != nil && len(s.backBuffer[c.x+i]) > c.y+j {
				s.backBuffer[c.x+i][c.y+j] = c.bgColor
			}
		}
	}
}

// BufferChildren renders all visible children of parent onto the back buffer
func (s *Screen) BufferChildren(children []*Container) {
	if len(children) == 0 {
		return
	}

	s.SortElements(children)

	for _, c := range children {
		if s.IsElementOnScreen(c) {
			s.BufferContainer(c)
			s.BufferChildren(c.children)
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
				last = c

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
				s.BufferChildren(WINDOW.page.children)
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
