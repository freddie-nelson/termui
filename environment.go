package termui

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/term"
)

var (
	envTicker     *time.Ticker
	stopEnvTicker chan bool
)

// Environment stores information about the current terminal environment
type Environment struct {
	width  int
	height int
}

// SetSize sets the width and height of the terminal in characters
func (e *Environment) SetSize(w int, h int) {
	e.width = w
	e.height = h
}

// GetSize returns the width and height of the terminal in characters
func (e *Environment) GetSize() (int, int) {
	return e.width, e.height
}

func (e *Environment) Update() {
	// update rows and cols
	fd := int(os.Stdout.Fd())
	width, height, err := term.GetSize(fd)
	if err != nil {
		// TODO add logger
		fmt.Println("error:", err)
		return
	}
	e.SetSize(width, height)
	// fmt.Printf("%v %v", width, height)
}

// StartWatcher starts the environment watcher ticker
func (e *Environment) StartWatcher() {
	envTicker = time.NewTicker(300 * time.Millisecond)
	stopEnvTicker = make(chan bool)

	go func() {
		for {
			select {
			case <-stopEnvTicker:
				return
			case <-envTicker.C:
				e.Update()
			}
		}
	}()
}

// StopWatcher stops the environment watcher ticker
func (e *Environment) StopWatcher() {
	envTicker.Stop()
	stopEnvTicker <- true
}

func NewEnvironment() *Environment {
	return &Environment{}
}
