package termui

import (
	"os"
	"sort"
	"time"
)

var (
	drawTicker     *time.Ticker
	stopDrawTicker chan bool
)

// holds color and character of individual cell in buffer
type Cell struct {
	char    rune
	color   Color
	bgColor Color
	element *Element
}

// Screen handles rendering of ui
type Screen struct {
	frontBuffer [][]Cell
	backBuffer  [][]Cell
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

	color, bgColor := WINDOW.page.Colors()

	buffer := make([][]Cell, width)
	for x := range buffer {
		buffer[x] = make([]Cell, height)

		for y := range buffer[x] {
			buffer[x][y] = Cell{' ', color, bgColor, nil}
		}
	}

	if front {
		s.frontBuffer = buffer
	} else {
		s.backBuffer = buffer
	}
}

// IsElementOnScreen returns true if the given container is currently on screen
func (s *Screen) IsElementOnScreen(c Element) bool {
	x, y, _ := c.Position()
	width, height := c.Size()

	visibleVertical := (y >= s.scrollY || y+height >= s.scrollY) && y < s.scrollY+ENV.height
	visibleHorizontal := (x >= s.scrollX || x+width >= s.scrollX) && x < s.scrollX+ENV.width
	return visibleVertical && visibleHorizontal
}

// SortElements sorts elements by zIndex ascending, y ascending, x ascending
func (s *Screen) SortElements(es []Element) {
	sort.SliceStable(es, func(i, j int) bool {
		x1, y1, z1 := es[i].Position()
		x2, y2, z2 := es[j].Position()

		if z1 < z2 {
			return true
		} else if z1 > z2 {
			return false
		}

		if y1 < y2 {
			return true
		} else if y1 > y2 {
			return false
		}

		return x1 < x2
	})
}

// BufferElement renders the element to the back buffer
func (s *Screen) BufferElement(e Element) {
	width, height := e.Size()
	x, y, _ := e.Position()

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if s.backBuffer[x+j] != nil && len(s.backBuffer[x+j]) > y+i {
				color, bgColor := e.Colors()
				s.backBuffer[x+j][y+i] = Cell{e.CharAt(j, i), color, bgColor, &e}
			}
		}
	}
}

// BufferChildren renders all visible children of parent onto the back buffer
func (s *Screen) BufferChildren(children []Element) {
	if len(children) == 0 {
		return
	}

	s.SortElements(children)

	for _, c := range children {
		if s.IsElementOnScreen(c) {
			s.BufferElement(c)
			s.BufferChildren(c.Children())
		}
	}
}

// DrawFrame swaps buffers and draws front buffer to terminal
func (s *Screen) DrawFrame() {
	s.SwapBuffers()

	output := ""
	for y := 0; y < ENV.height; y++ {
		lastC := NewColor(-1, -1, -1)
		lastBg := NewColor(-1, -1, -1)

		for x := 0; x < ENV.width; x++ {
			c := s.frontBuffer[x][y]

			// color cell
			if c.bgColor != lastBg {
				lastBg = c.bgColor
				output += lastBg.ToANSII(false)
			}
			if c.color != lastC {
				lastC = c.color
				output += lastC.ToANSII(true)
			}

			output += string(c.char)
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
