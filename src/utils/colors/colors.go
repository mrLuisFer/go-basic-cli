package colors

import (
	colors "github.com/fatih/color"
)

// Paint yellow the text for the console
func Warn(msg string) {
	yellow := colors.New(colors.FgYellow).Add(colors.Bold).Set()
	yellow.Print(msg)
}

// Paint blue the text for the console
func Info(msg string) {
	blue := colors.New(colors.FgBlue).Add(colors.Bold).Set().PrintlnFunc()
	blue(msg)
}

// Paint green the text for the console
func Succes(msg string) {
	green := colors.New(colors.FgGreen).Set().PrintlnFunc()
	green(msg)
}
