// Copyright 2021.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package input

import (
	"fmt"

	"github.com/mertcandav/dinogo/keyboard"
)

// ReadRune reads one char from command line
// without press enter.
// Not printed pressed rune to cli.
//
// Special case is:
// ReadRune() 'rune' is zero if pressed enter
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
		if r != 0 {
			return r, nil
		}
	}
}

// ReadLine reads line from command line.
// No supports remove, moving. Only get rune input.
func ReadLine() ([]rune, error) {
	var runes []rune
	for {
		r, err := ReadRune()
		if err != nil {
			return nil, err
		}
		if r == 0 {
			break
		}
		fmt.Print(string(r))
		runes = append(runes, r)
	}
	return runes, nil
}
