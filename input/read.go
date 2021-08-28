package input

import (
	"github.com/mertcandav/dinogo/keyboard"
)

// ReadRune reads one char from command line
// without press enter.
//
// Special case is:
// ReadRune() pressed enter if 'rune' is zero
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
		case keyboard.KeyEnter:
			return rune(0), nil
		}
		if r >= 0 {
			return r, nil
		}
	}
}
