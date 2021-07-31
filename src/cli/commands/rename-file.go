package commands

import (
	"fmt"

	colors "github.com/mrLuisFer/go-basic-cli/src/utils/colors"
)

func RenameFile() {
	var oldFilename string
	colors.Info("Insert old filename", true)
	fmt.Scanln(&oldFilename)

	var newFilename string
	colors.Info("Insert new filename", true)
	fmt.Scanln(&newFilename)

	fmt.Printf("Renaming: %s -> %s", oldFilename, newFilename)
}
