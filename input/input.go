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
	AppendedRune Event
}

// Init new Input instance.
func Init() *Input {
	return &Input{
		Actions: map[keyboard.Key]Action{
			keyboard.KeyEnter:      ActionEnter,
			keyboard.KeyArrowLeft:  ActionArrowLeft,
			keyboard.KeyArrowRight: ActionArrowRight,
			keyboard.KeyBackspace2: ActionBackspace,
			keyboard.KeyHome:       ActionHome,
			keyboard.KeyEnd:        ActionEnd,
			keyboard.KeyCtrlC:      ActionCtrlC,
			keyboard.KeySpace:      ActionSpace,
			keyboard.KeyTab:        ActionTab,
		},
		AppendedRune: AppendedRune,
	}
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
			result := action(ActionInfo{
				Input: i,
				Rune:  &r,
			})
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
			i.AppendedRune(i)
		}
	}
	return nil
}

// Reset runes, positions and other saved positions.
func (i *Input) Reset() {
	i.Runes = nil
	i.Position.Line = 0
	i.Position.Column = 0
}