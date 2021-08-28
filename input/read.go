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
		case 9: // tab
			r = '\t'
		case 32: // space
			r = ' '
		}
		if r >= 0 {
			return r, nil
		}
	}
}
