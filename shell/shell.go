package shell

import (
	"errors"
	"fmt"

	"github.com/mertcandav/dinogo/input"
	"github.com/mertcandav/dinogo/keyboard"
	"github.com/mertcandav/dinogo/message"
	"github.com/mertcandav/dinogo/shell/history"
	"github.com/mertcandav/dinogo/shell/utils"
	"github.com/mertcandav/dinogo/terminal"
)

// Shell is an interface for CLI command application.
type Shell struct {
	Commands      []Command
	PromptMessage string
	Sep           string // Command seperator.
	Input         *input.Input
	History       *history.History // Set null if you want not log command history.
	Prefix        string
}

//  Init returns new instance of Shell.
func Init() *Shell {
	shell := &Shell{
		Sep:     " ",
		Prefix:  "$ ",
		Input:   input.Init(),
		History: history.Init(),
	}
	shell.Input.UpdatedRunes = func(*input.Input, interface{}) input.ActionResult {
		shell.History.End()
		return input.ActionResult{}
	}
	shell.Input.Actioning = func(_ *input.Input, tag interface{}) input.ActionResult {
		tags := tag.([]interface{})
		return tags[0].(input.Action)(tags[1].(input.ActionInfo), shell)
	}
	shell.Input.Actions[keyboard.KeyArrowUp] = historyUp
	shell.Input.Actions[keyboard.KeyArrowDown] = historyDown
	return shell
}

// Loop entire default input loop.
func (s *Shell) Loop() {
	for {
		err := s.Prompt()
		if err != nil {
			message.Errorln(err)
		}
	}
}

// GetInput gets input via command line.
//
// Returns nil if failed.
func (s *Shell) GetInput(msg string) []rune {
	if s.Input == nil {
		return nil
	}
	fmt.Print(msg)
	terminal.SavePosition()
	err := s.Input.Get()
	if err != nil {
		return nil
	}
	return s.Input.Runes
}

// Prompt gets input with prompt message and prefix
// via command line and process input.
//
// Returns nil if failed.
func (s *Shell) Prompt() error {
	runes := s.GetInput(s.PromptMessage + s.Prefix)
	if runes == nil {
		return nil
	}
	input := string(runes)
	if input == "" {
		return nil
	}
	if s.History != nil {
		s.History.Add(input)
		s.History.End()
	}
	return s.DoCommand(input)
}

// PromptMsg sets prompt message to specified message
// and returns input from prompt.
func (s *Shell) PromptMsg(msg string) error {
	s.PromptMessage = msg
	return s.Prompt()
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
