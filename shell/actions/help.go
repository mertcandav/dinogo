package actions

import (
	"strings"

	"github.com/mertcandav/dinogo/shell"
)

// Help is default action for your shell instance.
// List command and help texts.
func Help(info shell.CommandActionInfo) {
	// Margin for between command names and help texts.
	const margin = 5
	max := 0
	for _, command := range info.Shell.Commands {
		length := len(command.Name)
		if max < length {
			max = length
		}
	}
	// Ready help text.
	var sb strings.Builder
	for _, command := range info.Shell.Commands {
		sb.WriteString(command.Name)
		for counter := len(command.Name); counter <= max+margin; counter++ {
			sb.WriteByte(' ')
		}
		sb.WriteString(command.Help)
		sb.WriteByte('\n')
	}
	print(sb.String())
}
