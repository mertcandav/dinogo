package actions

import (
	"github.com/mertcandav/dinogo/shell"
	"github.com/mertcandav/dinogo/terminal"
)

// Clear is default action for your shell instance.
// Clear entire screen.
func Clear(info shell.CommandActionInfo) {
	terminal.Clear()
	terminal.SetPosition(0, 0)
}
