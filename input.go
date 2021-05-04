package termui

import (
	"bufio"
	"os"
)

// Input handles all user input and associated events e.g. focused element
type Input struct {
	in *os.File
}

func (input *Input) StartKeyboardListener() {
	go func() {
		for {
			// only read single characters, the rest will be ignored!!
			consoleReader := bufio.NewReaderSize(input.in, 1)
			input, _ := consoleReader.ReadByte()

			// ESC = 27 and Ctrl-C = 3
			if input == 27 || input == 3 {
				Close()
			}
		}
	}()
}

// NewInput returns a pointer to a new Input
func NewInput(in *os.File) *Input {
	input := Input{in: in}
	input.StartKeyboardListener()

	return &input
}
