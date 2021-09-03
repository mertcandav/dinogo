// Copyright 2021.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package input

import (
	"github.com/mertcandav/dinogo/terminal"
)

// Event is input event.
type Event func(*Input, interface{}) ActionResult

// Actioning is default Actioning event.
func Actioning(i *Input, tag interface{}) ActionResult {
	tags := tag.([]interface{})
	return tags[0].(Action)(tags[1].(ActionInfo), nil)
}

// AppendedRune is default AppendedRune event.
func AppendedRune(i *Input, _ interface{}) ActionResult {
	terminal.ClearLineCE()
	part := i.Runes[i.Position.Column:]
	runes := i.PrintText(i, part).Tag.([]rune)
	if len(runes) > 0 {
		i.Position.Column++
		terminal.MoveLeft(i.Position.Column + len(runes) - 1)
		terminal.MoveRight(i.Position.Column)
	}
	return ActionResult{}
}

// PrintText is default PrintText event.
// Tag is should be []rune and event returns printed []rune.
func PrintText(_ *Input, tag interface{}) ActionResult {
	text := string(tag.([]rune))
	print(text)
	return ActionResult{Tag: tag}
}

// PrintTextPassword is default PrintText event for password config.
// Tag is should be []rune and event returns printed []rune.
func PrintTextPassword(_ *Input, tag interface{}) ActionResult {
	runes := tag.([]rune)
	for range runes {
		print("●")
	}
	return ActionResult{Tag: runes}
}
