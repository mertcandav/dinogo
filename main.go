package main

import (
	"fmt"

	"github.com/mertcandav/dinogo/input"
	"github.com/mertcandav/dinogo/shell"
	"github.com/mertcandav/dinogo/shell/actions"
	"github.com/mertcandav/dinogo/shell/history"
)

func about(info shell.CommandActionInfo) {
	println("This shell is example for Dinogo framework.")
}

func main() {
	sh := shell.Shell{
		Sep:     " ",
		Prefix:  "$ ",
		Input:   input.Init(),
		History: &history.History{},
		Commands: []shell.Command{
			{Name: "help", Help: "Show help.", Action: actions.Help},
			{Name: "history", Help: "List command history.", Action: actions.History},
			{Name: "exit", Help: "Exit shell.", Action: actions.Exit},
			{Name: "about", Help: "Show about this shell.", Action: about},
		},
	}
	for {
		err := sh.Prompt("Example Shell")
		if err != nil {
			fmt.Println("error:", err)
		}
	}
}
