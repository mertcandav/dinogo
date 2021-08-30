package message

import (
	"fmt"

	"github.com/mertcandav/dinogo/terminal"
)

// Error prints error and reset font.
func Error(v ...interface{}) {
	print(ErrorFont)
	fmt.Print(v...)
	terminal.Reset()
}

// Errorln prints error with new line and reset font.
func Errorln(v ...interface{}) {
	Error(v...)
	println()
}

// Warn prints warn and reset font.
func Warn(v ...interface{}) {
	print(WarnFont)
	fmt.Print(v...)
	terminal.Reset()
}

// Warnln prints warn with new line and reset font.
func Warnln(v ...interface{}) {
	Warn(v...)
	println()
}

// Info prints info and reset font.
func Info(v ...interface{}) {
	print(InfoFont)
	fmt.Print(v...)
	terminal.Reset()
}

// Infoln prints info with new line and reset font.
func Infoln(v ...interface{}) {
	Info(v...)
	println()
}

// Tip prints tip and reset font.
func Tip(v ...interface{}) {
	print(TipFont)
	fmt.Print(v...)
	terminal.Reset()
}

// Tipln prints tip with new line and reset font.
func Tipln(v ...interface{}) {
	Tip(v...)
	println()
}

// Success prints success and reset font.
func Success(v ...interface{}) {
	print(SuccessFont)
	fmt.Print(v...)
	terminal.Reset()
}

// Successln prints success with new line and reset font.
func Successln(v ...interface{}) {
	Success(v...)
	println()
}
