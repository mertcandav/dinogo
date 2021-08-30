package actions

import (
	"strings"

	"github.com/mertcandav/dinogo/shell"
)

// History is default action for your shell instance.
// List command history.
func History(info shell.CommandActionInfo) {
	if info.Shell.History == nil {
		return
	}
	var sb strings.Builder
	for _, command := range info.Shell.History.History() {
		sb.WriteString(command)
		sb.WriteByte('\n')
	}
	print(sb.String())
}
