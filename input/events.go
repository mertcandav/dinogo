// Copyright 2021.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package input

import "github.com/mertcandav/dinogo/terminal"

// Event is input event.
type Event func(*Input)

// AppendedRune is default AppendedRune event.
func AppendedRune(i *Input) {
	terminal.ClearLineCE()
	part := i.Runes[i.Position.Column:]
	print(string(part))
	i.Position.Column++
	terminal.MoveLeft(i.Position.Column + len(part) - 1)
	terminal.MoveRight(i.Position.Column)
}
