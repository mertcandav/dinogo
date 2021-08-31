package shell

import (
	"github.com/mertcandav/dinogo/input"
	"github.com/mertcandav/dinogo/terminal"
)

func historyUp(info input.ActionInfo, tag interface{}) input.ActionResult {
	sh := tag.(*Shell)
	if sh.History != nil && sh.History.Position() != -1 {
		terminal.RestorePosition()
		terminal.ClearLineCE()
		command, _ := sh.History.Get()
		sh.Input.Runes = []rune(command)
		info.Input.Position.Column = len(info.Input.Runes)
		print(command)
		_ = sh.History.Prev()
	}
	return input.ActionResult{Skip: true}
}

func historyDown(info input.ActionInfo, tag interface{}) input.ActionResult {
	sh := tag.(*Shell)
	if sh.History != nil {
		if sh.History.Position() == -1 {
			_ = sh.History.Next()
		}
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
		}
	}
	return input.ActionResult{Skip: true}
}
