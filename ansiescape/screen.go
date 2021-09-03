package ansiescape

// Escape sequences for create/close alternate screen buffer.
const (
	AlternateScreenBufferOpen  = "\033[?1049h"
	AlternateScreenBufferClose = "\033[?1049l"
)
