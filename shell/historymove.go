package shell

import (
	"github.com/mertcandav/dinogo/input"
	"github.com/mertcandav/dinogo/terminal"
)

func historyUp(info input.ActionInfo, tag interface{}) input.ActionResult {
	sh := tag.(*Shell)
	if sh.History != nil &&
		(sh.historyOrigin && sh.History.Any() || sh.History.Prev()) {
		command, _ := sh.History.Get()
		sh.Input.Runes = []rune(command)
		info.Input.Position.Column = len(info.Input.Runes)
		terminal.HideCursor()
		terminal.RestorePosition()
		terminal.ClearLineCE()
		print(command)
		terminal.ShowCursor()
		sh.historyOrigin = false
	}
	return input.ActionResult{Skip: true}
}

func historyDown(info input.ActionInfo, tag interface{}) input.ActionResult {
	sh := tag.(*Shell)
	if sh.History != nil && !sh.historyOrigin {
		terminal.HideCursor()
		terminal.RestorePosition()
		terminal.ClearLineCE()
		if sh.History.Next() {
			command, _ := sh.History.Get()
			sh.Input.Runes = []rune(command)
			info.Input.Position.Column = len(info.Input.Runes)
			print(command)
		} else {
			sh.Input.Runes = make([]rune, 0)
			info.Input.Position.Column = 0
			sh.historyOrigin = true
		}
		terminal.ShowCursor()
	}
	return input.ActionResult{Skip: true}
}
