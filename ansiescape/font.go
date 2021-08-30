// Copyright 2021.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ansiescape

import "fmt"

// ANSI font escape sequences.
const (
	ResetFont = "\033[0m"

	Bold      = "\u001b[1m"
	Underline = "\u001b[4m"
	Reversed  = "\u001b[7m"

	Black         = "\033[30m"
	Red           = "\033[31m"
	Green         = "\033[32m"
	Yellow        = "\033[33m"
	Blue          = "\033[34m"
	Magenta       = "\033[35m"
	Cyan          = "\033[36m"
	White         = "\033[37m"
	BrightBlack   = "\033[90m"
	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"

	BlackBackground         = "\033[40m"
	RedBackground           = "\033[41m"
	GreenBackground         = "\033[42m"
	YellowBackground        = "\033[43m"
	BlueBackground          = "\033[44m"
	MagentaBackground       = "\033[45m"
	CyanBackground          = "\033[46m"
	WhiteBackground         = "\033[47m"
	BrightBlackBackground   = "\033[100m"
	BrightRedBackground     = "\033[101m"
	BrightGreenBackground   = "\033[102m"
	BrightYellowBackground  = "\033[103m"
	BrightBlueBackground    = "\033[104m"
	BrightMagentaBackground = "\033[105m"
	BrightCyanBackground    = "\033[106m"
	BrightWhiteBackground   = "\033[107m"
)

// ForegroundByRGB returns sequence for
// set foreground by R, G and B values.
func ForegroundByRGB(r, g, b byte) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

// BackgroundByRGB returns sequence for
// set background by R, G and B values.
func BackgroundByRGB(r, g, b byte) string {
	return fmt.Sprintf("\033[48;2;%d;%d;%dm", r, g, b)
}
