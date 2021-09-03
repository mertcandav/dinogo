package terminal

import "github.com/mertcandav/dinogo/ansiescape"

// SetAlternateScreenBuffer create/close alternate screen buffer.
func SetAlternateScreenBuffer(state bool) {
	if state {
		print(ansiescape.AlternateScreenBufferOpen)
		return
	}
	print(ansiescape.AlternateScreenBufferClose)
}
