package cli

import (
	"fmt"

	commands "github.com/mrLuisFer/go-basic-cli/src/cli/commands"
	colors "github.com/mrLuisFer/go-basic-cli/src/utils/colors"
	stringutils "github.com/mrLuisFer/go-basic-cli/src/utils/stringUtils"
)

// Paint in the terminal the parameter that is passed to it
func print (param interface{}) {
	fmt.Print(param)
	fmt.Printf("\n")
}

// Initialize the CLI and the commands within it
func Init() {
	colors.Succes("Initializing CLI...")
	print("Insert any of the following commands:")

	var cmds = []string{"help", "add"}
	for _, cmd := range cmds {
		fmt.Printf("%v, ", cmd)
	}
	print(":")

	var command string
	fmt.Scan(&command)
	
	if(len(command) > 0) {
		command = stringutils.ToLowerCase(command)
		colors.Warn(command)
		fmt.Println(" command selected!")
	}
	if (command == "help") {
		commands.Help()
	}
}
