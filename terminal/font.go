package terminal

import (
	"github.com/mertcandav/dinogo/ansiescape"
)

// Reset resets font.
func Reset() {
	print(ansiescape.ResetFont)
}

// ForegroundByRGB sets foreground by R, G and B values.
func ForegroundByRGB(r, g, b byte) {
	print(ansiescape.ForegroundByRGB(r, g, b))
}

// BackgroundByRGB sets background by R, G and B values.
func BackgroundByRGB(r, g, b byte) {
	print(ansiescape.BackgroundByRGB(r, g, b))
}
