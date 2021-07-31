package commands

import (
	"fmt"

	"os"

	colors "github.com/mrLuisFer/go-basic-cli/src/utils/colors"

	utils "github.com/mrLuisFer/go-basic-cli/src/utils"
)

func RenameFile() {
	var oldFilename string
	colors.Info("Insert old filename", true)
	fmt.Scanln(&oldFilename)

	isFile := utils.FileExists(oldFilename)

	if isFile {
		var newFilename string
		colors.Info("Insert new filename", true)
		
    fmt.Scanln(&newFilename)
		fmt.Printf("Renaming: %s -> %s", oldFilename, newFilename)
		
    err := os.Rename(oldFilename, newFilename)

    if err != nil {
      colors.Error(err, true)
    }
	} else {
    colors.Error("Insert a valid fileName or filePath", true)
  }
}
