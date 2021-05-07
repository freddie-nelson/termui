package termui

import (
	"bufio"
	"os"
)

// Mouse buttons
const (
	LEFT_CLICK = iota
	MOUSE_WHEEL
	RIGHT_CLICK
	RELEASE_MOUSE
)

// Input handles all user input and associated events e.g. focused element
type Input struct {
	in *os.File
}

// IsMouseEvent returns true if a mouse event has been sent to stdin
func (input *Input) IsMouseEvent(r *bufio.Reader, key rune) bool {
	val := false

	if r.Buffered() == 5 {
		if char, _, _ := r.ReadRune(); key == 27 && char == 91 {
			val = true
		}

		r.UnreadRune()
	}

	return val
}

// readMouseEvent reads mouse event from r and returns it as button, modifier, x, y
func (input *Input) readMouseEvent(r *bufio.Reader) (byte, byte, byte, byte) {
	// discard M in ansii escape
	r.Discard(2)

	Cb, _ := r.ReadByte()

	// button pressed or release
	button := Cb & byte(3)
	modifier := Cb & byte(28)

	// position
	mouseX, _ := r.ReadByte()
	mouseY, _ := r.ReadByte()

	// correct for weird X10 scheme
	mouseX -= 33
	mouseY -= 33

	return button, modifier, mouseX, mouseY
}

// handleMouse handles mouse events
func (input *Input) handleMouse(r *bufio.Reader) {
	button, modifier, mouseX, mouseY := input.readMouseEvent(r)
}

// handleKeyboard handles keyboard events
func (input *Input) handleKeyboard(key rune) {
}

func (input *Input) StartInputListener() {
	go func() {
		reader := bufio.NewReader(input.in)

		for {
			key, _, _ := reader.ReadRune()

			if input.IsMouseEvent(reader, key) {
				input.handleMouse(reader)
			} else {
				// if ESC (27) or Ctrl-c (3) are pressed, exit
				if key == 27 || key == 3 {
					Close()
				}

				input.handleKeyboard(key)
			}
		}
	}()
}

// NewInput returns a pointer to a new Input
func NewInput(in *os.File) *Input {
	input := Input{in: in}
	input.StartInputListener()

	return &input
}
