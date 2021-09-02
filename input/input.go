// Copyright 2021.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package input

import (
	"github.com/mertcandav/dinogo/keyboard"
	"github.com/mertcandav/dinogo/terminal"
)

// Input interface.
type Input struct {
	Actions      map[keyboard.Key]Action
	Runes        []rune            // Getted runes.
	Position     terminal.Position // Position.
	Actioning    Event
	AppendedRune Event
	UpdatedRunes Event
	PrintText    Event
}

// Init new Input instance.
func Init(config uint8) *Input {
	input := &Input{
		Actions: map[keyboard.Key]Action{
			keyboard.KeyEnter:      ActionEnter,
			keyboard.KeyArrowLeft:  ActionArrowLeft,
			keyboard.KeyArrowRight: ActionArrowRight,
			keyboard.KeyBackspace:  ActionBackspace,
			keyboard.KeyBackspace2: ActionBackspace,
			keyboard.KeyHome:       ActionHome,
			keyboard.KeyEnd:        ActionEnd,
			keyboard.KeyCtrlC:      ActionCtrlC,
			keyboard.KeySpace:      ActionSpace,
			keyboard.KeyTab:        ActionTab,
			keyboard.KeyDelete:     ActionDelete,
		},
		AppendedRune: AppendedRune,
		UpdatedRunes: nil,
		Actioning:    Actioning,
		Runes:        make([]rune, 0),
	}
	switch config {
	case Password:
		input.PrintText = func(_ *Input, tag interface{}) ActionResult {
			runes := tag.([]rune)
			for range runes {
				print("●")
			}
			return ActionResult{Tag: string(runes)}
		}
	default:
		input.PrintText = PrintText
	}
	return input
}

// Get reads line from command line.
func (i *Input) Get() error {
	i.Reset()
	for {
		r, key, err := keyboard.GetSingleKey()
		if err != nil {
			return err
		}
		action, ok := i.Actions[key]
		if ok {
			result := i.Actioning(i, []interface{}{action, ActionInfo{
				Input: i,
				Rune:  &r,
			}})
			if result.Stop {
				break
			} else if result.Skip {
				continue
			}
		}
		if r == 0 {
			continue
		}
		i.Runes = append(i.Runes[:i.Position.Column],
			append([]rune{r}, i.Runes[i.Position.Column:]...)...)
		if i.AppendedRune != nil {
			i.AppendedRune(i, nil)
		}
		if i.UpdatedRunes != nil {
			i.UpdatedRunes(i, nil)
		}
	}
	return nil
}

// Reset runes, positions and other saved positions.
func (i *Input) Reset() {
	i.Runes = make([]rune, 0)
	i.Position.Line = 0
	i.Position.Column = 0
}
