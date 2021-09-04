// Copyright 2021.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ansiescape

import "fmt"

const (
	// SavePosition is save cursor position.
	SavePosition = "\033[s"
	// RestorePosition is restore cursor position.
	RestorePosition = "\033[u"
)

// MoveUp returns sequence for move the cursor up N lines.
func MoveUp(n int) string {
	return fmt.Sprintf("\033[%dA", n)
}

// MoveDown returns sequence for move the cursor down N lines.
func MoveDown(n int) string {
	return fmt.Sprintf("\033[%dB", n)
}

// MoveBack returns sequence for move the cursor backward N columns.
func MoveRight(n int) string {
	return fmt.Sprintf("\033[%dC", n)
}

// MoveBack returns sequence for move the cursor backward N columns.
func MoveLeft(n int) string {
	return fmt.Sprintf("\033[%dD", n)
}

// NextLine returns sequence for move cursor to
// beginning of line N lines down.
func NextLine(n int) string {
	return fmt.Sprintf("\u001b[%dE", n)
}

// PrevLine returns sequence for move cursor to
// beginning of line N lines up.
func PrevLine(n int) string {
	return fmt.Sprintf("\u001b[%dF", n)
}

// SetColumn returns sequence for move cursor to column N.
func SetColumn(n int) string {
	return fmt.Sprintf("\u001b[%dG", n)
}

// SetPosition returns sequence for put the cursor at line and column.
func SetPosition(line, column int) string {
	return fmt.Sprintf("\033[%d;%dH", line, column)
}
