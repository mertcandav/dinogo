package shell

// CommandAction is type alias for command actions.
type CommandAction func(CommandActionInfo)

// CommandActionInfo is information for command actions.
type CommandActionInfo struct {
	Shell   *Shell
	Command *Command
	Input   string // Inputted command.
}

// Command is shell command.
type Command struct {
	Name   string
	Help   string
	Action CommandAction
}
