package cli

import (
	commands "github.com/mrLuisFer/go-basic-cli/src/cli/commands"
	colors "github.com/mrLuisFer/go-basic-cli/src/utils/colors"
)

// Validate each command with the string that is passed to it as a parameter
func CheckCommands(command string) {
	printCmdSelected(command)
	if command == "help" {
		commands.Help()
	} else if command == "sum" {
		commands.Sum()
	} else if command == "create" {
		commands.CreateFile()
	} else if command == "rename" {
		commands.RenameFile()
	} else if command == "delete" {
		commands.RemoveFile()
	} else {
		colors.Error("✘ Choose an available command!", false)
	}
}
