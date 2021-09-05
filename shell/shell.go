package shell

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mertcandav/dinogo/input"
	"github.com/mertcandav/dinogo/keyboard"
	"github.com/mertcandav/dinogo/message"
	"github.com/mertcandav/dinogo/shell/history"
	"github.com/mertcandav/dinogo/shell/utils"
	"github.com/mertcandav/dinogo/terminal"
)

// Shell is an interface for CLI command application.
type Shell struct {
	historyOrigin bool
	completeIndex int

	Commands      []Command
	PromptMessage string
	// Command separator.
	Sep   string
	Input *input.Input
	// Set null if you want not log command history.
	History *history.History
	// Complete input by command history.
	InputComplete bool
	Prefix        string
}

//  Init returns new instance of Shell.
func Init() *Shell {
	shell := &Shell{
		completeIndex: -1,

		Sep:           " ",
		Prefix:        "> ",
		Input:         input.Init(input.Classic),
		History:       &history.History{},
		InputComplete: true,
	}
	shell.Input.UpdatedRunes = func(*input.Input, interface{}) input.ActionResult {
		if shell.History != nil {
			shell.History.End()
			shell.historyOrigin = true
			shell.complete()
		}
		return input.ActionResult{}
	}
	shell.Input.Actioning = func(_ *input.Input, tag interface{}) input.ActionResult {
		tags := tag.([]interface{})
		return tags[0].(input.Action)(tags[1].(input.ActionInfo), shell)
	}
	shell.Input.Actions[keyboard.KeyArrowUp] = historyUp
	shell.Input.Actions[keyboard.KeyArrowDown] = historyDown
	shell.Input.Actions[keyboard.KeyCtrlD] = completeCommand
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
	if s.History != nil {
		s.History.Add(string(s.Input.Runes))
		s.History.End()
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

// Complete complates input.
func (s *Shell) complete() {
	text := string(s.Input.Runes)
	s.completeIndex = -1
	if text == "" {
		terminal.ClearLineCE()
		return
	}
	step := len(s.Input.Runes) - s.Input.Position.Column
	terminal.HideCursor()
	defer terminal.ShowCursor()
	for index, command := range s.History.History() {
		if command != text && strings.HasPrefix(command, text) {
			terminal.ForegroundByRGB(100, 100, 100)
			if step > 0 {
				terminal.MoveRight(step)
			}
			text = command[len(text):]
			print(text)
			terminal.Reset()
			terminal.MoveLeft(len(text))
			if step > 0 {
				terminal.MoveLeft(step)
			}
			s.completeIndex = index
			return
		}
	}
	terminal.MoveRight(step)
	terminal.ClearLineCE()
	terminal.MoveLeft(step)
}
