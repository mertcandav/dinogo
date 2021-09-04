package ansiescape

// SetTitle returns sequence for set title of window.
func SetTitle(title string) string {
	return "\033]0;" + title + "\007"
}
