package commands

import (
	"fmt"

	"os"

	"reflect"

	colors "github.com/mrLuisFer/go-basic-cli/src/utils/colors"
)

// CreateFile is a command that creates a file
// with the name you pass in the terminal
func CreateFile() {
	var fileName string

	colors.Info("Insert filename...", true)
	fmt.Scanln(&fileName)

	if reflect.DeepEqual(fileName, "") || reflect.DeepEqual(fileName, " ") || len(fileName) <= 0 {
		colors.Error("File name cannot be empty", true)
		os.Exit(1)
	}

	typeOfFileName := reflect.TypeOf(fileName).Kind()

	if typeOfFileName != reflect.String {
		colors.Error("File name must be a string", true)
		os.Exit(1)
    
	} else {
		file, err := os.Create(fileName)
		if err != nil {
			colors.Error("Error creating file: ", false)
			fmt.Print(err)
			return
		}
		defer file.Close()
	}
}
