package drawing

import (
	"strings"

	"github.com/mertcandav/dinogo/ansiescape"
	"github.com/mertcandav/dinogo/terminal"
)

// Rectangle for 2D rectangle area.
type Rectangle struct {
	Position Point
	Width    uint
	Height   uint
}

// Draw to screen.
//
// Resets terminal font after drawing.
func (r Rectangle) Draw(pen Pen) {
	terminal.SetPosition(r.Position.Y, r.Position.X)
	if r.Height < 1 {
		return
	}
	var sb strings.Builder
	width := r.Width
	if width > 0 {
		for width > 0 {
			width--
			sb.WriteRune(pen.Mark)
		}
		sb.WriteRune(pen.Mark)
		sb.WriteString("\n" + ansiescape.SetColumn(r.Position.X))
	}
	for r.Height > 0 {
		r.Height--
		width = r.Width
		if width > 0 {
			sb.WriteRune(pen.Mark)
			for width > 1 {
				width--
				sb.WriteByte(' ')
			}
			sb.WriteRune(pen.Mark)
		}
		sb.WriteString("\n" + ansiescape.SetColumn(r.Position.X))
	}
	width = r.Width
	if width > 0 {
		for width > 0 {
			width--
			sb.WriteRune(pen.Mark)
		}
		sb.WriteRune(pen.Mark)
	}
	print(sb.String())
	terminal.Reset()
}

// Fill to screen.
//
// Resets terminal font after drawing.
func (r Rectangle) Fill(color RGB) {
	terminal.SetPosition(r.Position.Y, r.Position.X)
	terminal.BackgroundByRGB(color.Red, color.Green, color.Blue)
	var sb strings.Builder
	for r.Height > 0 {
		r.Height--
		width := r.Width
		for width > 0 {
			width--
			sb.WriteByte(' ')
		}
		sb.WriteString("\n" + ansiescape.SetColumn(r.Position.X))
	}
	print(sb.String())
	terminal.Reset()
}
