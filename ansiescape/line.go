package ansiescape

const (
	// ClearLine clears entire line.
	ClearLine = "\u001b[2K"
	// ClearLineCB clears from cursor to start of line.
	ClearLineCB = "\u001b[1K"
	// ClearLineCE clears from cursor to end of line.
	ClearLineCE = "\u001b[0K"
)
