package message

import (
	"github.com/mertcandav/dinogo/ansiescape"
)

// ErrorMessage returns error and reset font message.
func ErrorMessage(m string) string {
	return ErrorFont + m + ansiescape.ResetFont
}

// WarnMessage returns warn and reset font message.
func WarnMessage(m string) string {
	return WarnFont + m + ansiescape.ResetFont
}

// InfoMessage returns info and reset font message.
func InfoMessage(m string) string {
	return InfoFont + m + ansiescape.ResetFont
}

// TipMessage returns tip and reset font message.
func TipMessage(m string) string {
	return TipFont + m + ansiescape.ResetFont
}

// SuccessMessage returns success and reset font message.
func SuccessMessage(m string) string {
	return SuccessFont + m + ansiescape.ResetFont
}
