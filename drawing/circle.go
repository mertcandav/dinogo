package drawing

import (
	"math"
	"strings"

	"github.com/mertcandav/dinogo/ansiescape"
	"github.com/mertcandav/dinogo/terminal"
)

// Circle for 2D circle area.
type Circle struct {
	Position Point
	Width    uint
	Height   uint
}

// Draw to screen.
//
// Resets terminal font after drawing.
func (c Circle) Draw(pen Pen) {
	if c.Width < 1 || c.Height < 1 {
		return
	}
	var sb strings.Builder
	height := int(c.Height)
	width := int(c.Width)
	fudge := .8
	for y := 1; y <= height*2+1; y++ {
		for x := 1; x <= width*2+1; x++ {
			cy := math.Abs(float64(height - y + 1))
			cx := math.Abs(float64(width - x + 1))
			cr := int(math.Floor(math.Sqrt((cx*cx)+(cy*cy)) + fudge))
			if cr > width {
				sb.WriteString("  ")
				continue
			}
			sb.WriteRune(pen.Mark)
			sb.WriteByte(' ')
		}
		sb.WriteByte('\n')
	}
	terminal.SetPosition(c.Position.Y, c.Position.X)
	terminal.ForegroundByRGB(pen.Color.Red, pen.Color.Green, pen.Color.Blue)
	print(sb.String()[:sb.Len()-1] /* Remove unnecessary line from end */)
	terminal.Reset()
}

// Fill to screen.
//
// Resets terminal font after drawing.
func (c Circle) Fill(color RGB) {
	var sb strings.Builder
	height := int(c.Height)
	width := int(c.Width)
	fudge := .8
	moveRight := ansiescape.MoveRight(2)
	for y := 1; y <= height*2+1; y++ {
		for x := 1; x <= width*2+1; x++ {
			cy := math.Abs(float64(height - y + 1))
			cx := math.Abs(float64(width - x + 1))
			cr := int(math.Floor(math.Sqrt((cx*cx)+(cy*cy)) + fudge))
			if cr > width {
				sb.WriteString(moveRight)
				continue
			}
			sb.WriteString("  ")
		}
		sb.WriteByte('\n')
	}
	terminal.SetPosition(c.Position.Y, c.Position.X)
	terminal.BackgroundByRGB(color.Red, color.Green, color.Blue)
	print(sb.String()[:sb.Len()-1] /* Remove unnecessary line from end */)
	terminal.Reset()
}
