package shell

import (
	"github.com/mertcandav/dinogo/input"
	"github.com/mertcandav/dinogo/terminal"
)

func completeCommand(info input.ActionInfo, tag interface{}) input.ActionResult {
	sh := tag.(*Shell)
	if sh.completeIndex >= 0 {
		terminal.HideCursor()
		defer terminal.ShowCursor()
		terminal.RestorePosition()
		terminal.ClearLineCE()
		info.Input.Runes = []rune(sh.History.History()[sh.completeIndex])
		sh.Input.Position.Column = len(info.Input.Runes)
		print(string(info.Input.Runes))
	}
	return input.ActionResult{Skip: true}
}
