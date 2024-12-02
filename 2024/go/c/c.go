package c

import (
	"fmt"
)

type (
	foregroundColor   string
	backgroundColor   string
	commonValuesColor string
)

var (
	// Foreground colors
	FgBlack   foregroundColor = "\x1b[30m"
	FgRed     foregroundColor = "\x1b[31m"
	FgGreen   foregroundColor = "\x1b[32m"
	FgYellow  foregroundColor = "\x1b[33m"
	FgBlue    foregroundColor = "\x1b[34m"
	FgMagenta foregroundColor = "\x1b[35m"
	FgCyan    foregroundColor = "\x1b[36m"
	FgWhite   foregroundColor = "\x1b[37m"

	// Background colors
	BgBlack   backgroundColor = "\x1b[40m"
	BgRed     backgroundColor = "\x1b[41m"
	BgGreen   backgroundColor = "\x1b[42m"
	BgYellow  backgroundColor = "\x1b[43m"
	BgBlue    backgroundColor = "\x1b[44m"
	BgMagenta backgroundColor = "\x1b[45m"
	BgCyan    backgroundColor = "\x1b[46m"
	BgWhite   backgroundColor = "\x1b[47m"

	// Common consts
	resetColor     commonValuesColor = "\x1b[0m"
	faintColor     commonValuesColor = "\x1b[2m"
	underlineColor commonValuesColor = "\x1b[4m"
)

// Color string foreground
func Foreground(s string, fgColor foregroundColor) string {
	return fmt.Sprintf("%s%s%s", fgColor, s, resetColor)
}

// Color string background
func Background(s string, fgColor foregroundColor, bgColor backgroundColor) string {
	return fmt.Sprintf("%s%s%s%s", fgColor, bgColor, s, resetColor)
}
