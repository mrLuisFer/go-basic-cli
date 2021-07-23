package commands

import (
	colors "github.com/mrLuisFer/go-basic-cli/src/utils/colors"
)

// Execute the Help command that displays information about the CLI
func Help() {
	colors.Info("Prints this help message", true)
}