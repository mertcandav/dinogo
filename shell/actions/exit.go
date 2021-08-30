package actions

import (
	"os"

	"github.com/mertcandav/dinogo/shell"
)

// Exit is default action for your shell instance.
// Exit program.
func Exit(info shell.CommandActionInfo) {
	os.Exit(0)
}
