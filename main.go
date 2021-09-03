package main

import (
	"github.com/mertcandav/dinogo/input"
	"github.com/mertcandav/dinogo/shell"
	"github.com/mertcandav/dinogo/shell/actions"
)

func about(shell.CommandActionInfo) {
	println("This shell is example for Dinogo framework.")
}

func main() {
	sh := shell.Init()
	sh.Input.PrintText = input.PrintTextPassword
	sh.Commands = []shell.Command{
		{Name: "clear", Help: "Clear entire screen.", Action: actions.Clear},
		{Name: "about", Help: "Show about this shell.", Action: about},
		{Name: "history", Help: "List command history.", Action: actions.History},
		{Name: "help", Help: "Show help.", Action: actions.Help},
		{Name: "exit", Help: "Exit shell.", Action: actions.Exit},
	}
	sh.Loop()
}
