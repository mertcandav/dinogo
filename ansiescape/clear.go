// Copyright 2021.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ansiescape

const (
	// ClearScreen clears entire screen.
	ClearScreen = "\u001b[2J"
	// ClearScreenCB clears from cursor to beginning of screen.
	ClearScreenCB = "\u001b[1J"
	// ClearScreenCE clears from cursor until end of screen.
	ClearScreenCE = "\u001b[0J"

	// ClearLine clears entire line.
	ClearLine = "\u001b[2K"
	// ClearLineCB clears from cursor to start of line.
	ClearLineCB = "\u001b[1K"
	// ClearLineCE clears from cursor to end of line.
	ClearLineCE = "\u001b[0K"
)
