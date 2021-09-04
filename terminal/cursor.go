package terminal

import "github.com/mertcandav/dinogo/ansiescape"

// ShowCursor shows cursor.
func ShowCursor() { print(ansiescape.ShowCursor) }

// HideCursor hides cursor.
func HideCursor() { print(ansiescape.HideCursor) }
