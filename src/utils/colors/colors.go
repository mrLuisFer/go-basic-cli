package colors

import (
	colors "github.com/fatih/color"
)

func Warn(msg string) {
	yellow := colors.New(colors.FgYellow).Add(colors.Bold).Set().PrintlnFunc()
	yellow(msg)
}
