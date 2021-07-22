package cli

import (
	"fmt"
)

func print (param interface{}) {
	fmt.Print(param)
	fmt.Printf("\n")
}


func Init() {
	fmt.Print("Working the CLI file")
print("Insert any of the following commands:")

	var cmds = []string{"help", "add"}
	for _, cmd := range cmds {
		fmt.Printf("%v ", cmd)
	}
	print(":")

	var command string
	fmt.Scan(&command)
	command = stringutils.toLowerCase(command)

	if (command == "help") {
		print("Prints this help message")
	}
}
