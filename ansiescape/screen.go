package ansiescape

const (
	// ClearScreen clears entire screen.
	ClearScreen = "\u001b[2J"
	// ClearScreenCB clears from cursor to beginning of screen.
	ClearScreenCB = "\u001b[1J"
	// ClearScreenCE clears from cursor until end of screen.
	ClearScreenCE = "\u001b[0J"
)
