package commands

import (
	"fmt"

	"os"

	colors "github.com/mrLuisFer/go-basic-cli/src/utils/colors"
)

// CreateFile is a command that creates a file
// with the name you pass in the terminal
func CreateFile() {
	var fileName string
	
	colors.Info("Insert filename...", true)
	fmt.Scanln(&fileName)

	file, err := os.Create(fileName)
	if err != nil {
		colors.Error("Error creating file: ", false)
		fmt.Print(err)
		return
	}
	defer file.Close()

}
