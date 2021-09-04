package terminal

import "github.com/mertcandav/dinogo/ansiescape"

// SetTitle sets title of window.
func SetTitle(title string) { print(ansiescape.SetTitle(title)) }
