package termui

import "fmt"

type Color struct {
	r int
	g int
	b int
}

// ToANSII returns the color as an ANSII escape string (either fg or bg)
func (c *Color) ToANSII(fg bool) string {
	s := "48;2"
	if fg {
		s = "38;2"
	}

	return fmt.Sprintf("\033[%s;%v;%v;%vm", s, c.r, c.g, c.b)
}

func NewColor(r, g, b int) Color {
	return Color{r, g, b}
}
