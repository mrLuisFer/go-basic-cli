package colors

import (
	colors "github.com/fatih/color"
)

func Warn(msg string) {
	yellow := colors.New(colors.FgYellow).Add(colors.Bold).Set().PrintlnFunc()
	yellow(msg)
}

func Info(msg string) {
	blue := colors.New(colors.FgBlue).Add(colors.Bold).Set().PrintlnFunc()
	blue(msg)
}

func Succes(msg string) {
	green := colors.New(colors.FgGreen).Set().PrintlnFunc()
	green(msg)
}
