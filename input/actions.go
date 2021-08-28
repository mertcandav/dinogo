package input

import (
	"os"

	"github.com/mertcandav/dinogo/terminal"
)

// Action is Input action.
type Action func(ActionInfo) ActionResult

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
func ActionEnter(i ActionInfo) ActionResult {
	println()
	return ActionResult{Stop: true}
}

// ActionSpace is default key Space action.
func ActionSpace(i ActionInfo) ActionResult {
	*i.Rune = ' '
	return ActionResult{}
}

// ActionTab is default key Tab action.
func ActionTab(i ActionInfo) ActionResult {
	*i.Rune = '\t'
	return ActionResult{}
}

// ActionHome is default key Home action.
func ActionHome(i ActionInfo) ActionResult {
	terminal.MoveLeft(i.Input.Position.Column)
	i.Input.Position.Column = 0
	return ActionResult{Skip: true}
}

// ActionEnd is default key End action.
func ActionEnd(i ActionInfo) ActionResult {
	l := len(i.Input.Runes)
	if i.Input.Position.Column < l {
		terminal.MoveRight(l - i.Input.Position.Column)
		i.Input.Position.Column = l
	}
	return ActionResult{Skip: true}
}

// ActionArrowLeft is default key ArrowLeft action.
func ActionArrowLeft(i ActionInfo) ActionResult {
	if i.Input.Position.Column > 0 {
		i.Input.Position.Column--
		terminal.MoveLeft(1)
	}
	return ActionResult{Skip: true}
}

// ActionArrowRight is default key ArrowRight action.
func ActionArrowRight(i ActionInfo) ActionResult {
	if i.Input.Position.Column < len(i.Input.Runes) {
		i.Input.Position.Column++
		terminal.MoveRight(1)
	}
	return ActionResult{Skip: true}
}

// ActionCtrlC is default key Ctrl+C action.
func ActionCtrlC(i ActionInfo) ActionResult {
	println("\nterminated...")
	os.Exit(0)
	return ActionResult{}
}

// ActionBackspace is default key Backspace action.
func ActionBackspace(i ActionInfo) ActionResult {
	if i.Input.Position.Column > 0 {
		terminal.MoveLeft(1)
		terminal.ClearLineCE()
		part := i.Input.Runes[i.Input.Position.Column:]
		i.Input.Runes = append(i.Input.Runes[:i.Input.Position.Column-1], part...)
		print(string(part))
		i.Input.Position.Column--
		terminal.MoveLeft(len(part) - 1)
	}
	return ActionResult{Skip: true}
}
