package termui

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
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

// ClearTerminal uses platform specific commands to clear terminal
// TODO look for better implementation of this
func ClearTerminal() {
	clear := make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
