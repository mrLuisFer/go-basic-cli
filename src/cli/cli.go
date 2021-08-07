package cli

import (
	"fmt"

	colors "github.com/mrLuisFer/go-basic-cli/src/utils/colors"
	stringutils "github.com/mrLuisFer/go-basic-cli/src/utils/stringUtils"
)

// Paint in the terminal the parameter that is passed to it
func print (param interface{}) {
	fmt.Printf("%v \n", param)
}

// Paint the text blue the command that was selected
func printCmdSelected(cmd string){
	colors.Info(cmd, false)
	fmt.Printf(" command selected! \n \n")
}

// Initialize the CLI and the commands within it
func Init() {
	colors.Succes("Initializing CLI...", true)
	print("Insert any of the following commands...")

	var cmds = []string{"help", "sum", "create", "rename", "delete"}
	for _, cmd := range cmds {
		fmt.Printf("%s  ", cmd)
	}
	print(":")

	var command string
	// use Scanln to avoid multiple empty line breaks
	fmt.Scanln(&command)
	if(len(command) > 0) {
		command = stringutils.ToLowerCase(command)
	}

  if (len(command) > 0) {
    CheckCommands(command)
  }
}
