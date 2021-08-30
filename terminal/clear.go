// Copyright 2021.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package terminal

import "github.com/mertcandav/dinogo/ansiescape"

// Clear clears entire screen.
func Clear() {
	print(ansiescape.ClearScreen)
}

// ClearCB clears from cursor to beginning of screen.
func ClearCB() {
	print(ansiescape.ClearLineCB)
}

// ClearCE clears from cursor until end of screen.
func ClearCE() {
	print(ansiescape.ClearLineCE)
}

// ClearLine clears entire line.
func ClearLine() {
	print(ansiescape.ClearLine)
}

// ClearLineCB clears from cursor to start of line.
func ClearLineCB() {
	print(ansiescape.ClearLineCB)
}

// ClearLineCE clears from cursor to end of line.
func ClearLineCE() {
	print(ansiescape.ClearLineCE)
}
