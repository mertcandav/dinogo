package terminal

import (
	"github.com/mertcandav/dinogo/ansiescape"
)

// SavePosition saves cursor position.
func SavePosition() {
	print(ansiescape.SavePosition)
}

// RestorePosition is restore cursor position.
func RestorePosition() {
	print(ansiescape.RestorePosition)
}

// MoveUp moves the cursor up N lines.
func MoveUp(n int) {
	print(ansiescape.MoveUp(n))
}

// MoveDown moves the cursor down N lines.
func MoveDown(n int) {
	print(ansiescape.MoveDown(n))
}

// MoveBack moves the cursor backward N columns.
func MoveLeft(n int) {
	print(ansiescape.MoveLeft(n))
}

// MoveBack moves the cursor backward N columns.
func MoveRight(n int) {
	print(ansiescape.MoveRight(n))
}

// NextLine moves cursor to beginning of line N lines down.
func NextLine(n int) {
	print(ansiescape.NextLine(n))
}

// PrevLine moves cursor to beginning of line N lines up.
func PrevLine(n int) {
	print(ansiescape.PrevLine(n))
}

// SetColumn moves cursor to column N.
func SetColumn(n int) {
	print(ansiescape.SetColumn(n))
}

// SetPosition puts the cursor at line and column.
func SetPosition(line, column int) {
	print(ansiescape.SetPosition(line, column))
}

// SetTitle sets title of window.
func SetTitle(title string) {
	print(ansiescape.SetTitle(title))
}
