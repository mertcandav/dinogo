package shell

import (
	"errors"
	"fmt"

	"github.com/mertcandav/dinogo/input"
	"github.com/mertcandav/dinogo/shell/history"
	"github.com/mertcandav/dinogo/shell/utils"
)

// Shell is an interface for CLI command application.
type Shell struct {
	Commands []Command
	Sep      string // Command seperator.
	Input    *input.Input
	History  *history.History // Set null if you want not log command history.
	Prefix   string
}

// GetInput gets input via command line.
// Returns nil if failed.
func (s *Shell) GetInput(msg string) []rune {
	if s.Input == nil {
		return nil
	}
	fmt.Print(msg)
	err := s.Input.Get()
	if err != nil {
		return nil
	}
	return s.Input.Runes
}

// Prompt gets input with prefix via command line and process input.
// Returns nil if failed.
func (s *Shell) Prompt(msg string) error {
	runes := s.GetInput(msg + s.Prefix)
	if runes == nil {
		return nil
	}
	input := string(runes)
	if input == "" {
		return nil
	}
	if s.History != nil {
		s.History.Add(input)
	}
	return s.DoCommand(input)
}

// DoCommand executes command by specified command.
func (s *Shell) DoCommand(cmd string) error {
	commandName := utils.GetCommandName(cmd, s.Sep)
	for _, command := range s.Commands {
		if commandName != command.Name {
			continue
		}
		command.Action(CommandActionInfo{
			Shell:   s,
			Command: &command,
			Input:   cmd,
		})
		return nil
	}
	return errors.New("command is not found: '" + commandName + "'")
}
