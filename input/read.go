package input

import (
	"github.com/mertcandav/dinogo/keyboard"
)

// ReadRune reads one char from command line
// without press enter.
func ReadRune() (rune, error) {
	for {
		r, key, err := keyboard.GetSingleKey()
		if err != nil {
			return 0, err
		}
		switch key {
		case keyboard.KeyTab:
			r = '\t'
		case keyboard.KeySpace:
			r = ' '
		}
		if r >= 0 {
			return r, nil
		}
	}
}
