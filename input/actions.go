// Copyright 2021.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package input

import (
	"os"

	"github.com/mertcandav/dinogo/terminal"
)

// Action is Input action.
type Action func(ActionInfo, interface{}) ActionResult

//ActionInfo is information of input actions.
type ActionInfo struct {
	Input *Input
	Rune  *rune
}

// ActionResult is represents result of input action.
type ActionResult struct {
	Stop bool // Stop get.
	Skip bool // Skip rune append.
}

// ActionEnter is default key Enter action.
func ActionEnter(info ActionInfo, _ interface{}) ActionResult {
	println()
	return ActionResult{Stop: true}
}

// ActionSpace is default key Space action.
func ActionSpace(info ActionInfo, _ interface{}) ActionResult {
	*info.Rune = ' '
	return ActionResult{}
}

// ActionTab is default key Tab action.
func ActionTab(info ActionInfo, _ interface{}) ActionResult {
	*info.Rune = '\t'
	return ActionResult{}
}

// ActionHome is default key Home action.
func ActionHome(info ActionInfo, _ interface{}) ActionResult {
	terminal.MoveLeft(info.Input.Position.Column)
	info.Input.Position.Column = 0
	return ActionResult{Skip: true}
}

// ActionEnd is default key End action.
func ActionEnd(info ActionInfo, _ interface{}) ActionResult {
	length := len(info.Input.Runes)
	if info.Input.Position.Column < length {
		terminal.MoveRight(length - info.Input.Position.Column)
		info.Input.Position.Column = length
	}
	return ActionResult{Skip: true}
}

// ActionArrowLeft is default key ArrowLeft action.
func ActionArrowLeft(info ActionInfo, _ interface{}) ActionResult {
	if info.Input.Position.Column > 0 {
		info.Input.Position.Column--
		terminal.MoveLeft(1)
	}
	return ActionResult{Skip: true}
}

// ActionArrowRight is default key ArrowRight action.
func ActionArrowRight(info ActionInfo, _ interface{}) ActionResult {
	if info.Input.Position.Column < len(info.Input.Runes) {
		info.Input.Position.Column++
		terminal.MoveRight(1)
	}
	return ActionResult{Skip: true}
}

// ActionCtrlC is default key Ctrl+C action.
func ActionCtrlC(_ ActionInfo, _ interface{}) ActionResult {
	println("\nterminated...")
	os.Exit(0)
	return ActionResult{}
}

// ActionBackspace is default key Backspace action.
func ActionBackspace(info ActionInfo, _ interface{}) ActionResult {
	if info.Input.Position.Column > 0 {
		part := info.Input.Runes[info.Input.Position.Column:]
		info.Input.Position.Column--
		info.Input.Runes = append(
			info.Input.Runes[:info.Input.Position.Column], part...)
		part = info.Input.Runes[info.Input.Position.Column:]
		terminal.MoveLeft(1)
		terminal.ClearLineCE()
		if len(part) > 0 {
			print(string(part))
			terminal.MoveLeft(len(part))
		}
		if info.Input.UpdatedRunes != nil {
			info.Input.UpdatedRunes(info.Input, nil)
		}
	}
	return ActionResult{Skip: true}
}
