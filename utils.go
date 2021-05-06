package termui

import (
	"fmt"

	"golang.org/x/term"
)

// SetCursorPos sest cursor position to col, row
func SetCursorPos(col, row int) {
	SCREEN.out.WriteString(fmt.Sprintf("\033[%v;%vH", row, col))
}

// HideCursor hides the cursor in the terminal
func HideCursor() {
	SCREEN.out.WriteString("\033[?25l")
}

// ShowCursor shows the cursor in the terminal
func ShowCursor() {
	SCREEN.out.WriteString("\033[?25h")
}

// TrackMouse enables xterm VT200 mouse tracking if the current terminal supports it
func TrackMouse() {
	SCREEN.out.WriteString("\033[?1000h")
}

// UntrackMouse disables xterm VT200 mouse tracking if the current terminal supports it
func UntrackMouse() {
	SCREEN.out.WriteString("\033[?1000l")
}

// ResetAttributes resets all attributes in the terminal
func ResetAttributes() {
	SCREEN.out.WriteString("\033[0m")
}

// ClearTerminal uses platform specific commands to clear terminal
// TODO look for better implementation of this
func ClearTerminal() {
	SCREEN.out.WriteString("\033[2J")
}

// SetTermRawMode puts the terminal into raw mode, returns old state of terminal
func SetTermRawMode() *term.State {
	oldState, err := term.MakeRaw(int(SCREEN.out.Fd()))
	if err != nil {
		panic(err)
	}

	return oldState
}
