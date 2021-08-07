package commands

import (
	"fmt"
	"os"

	colors "github.com/mrLuisFer/go-basic-cli/src/utils/colors"
)

func RemoveFile() error {
  colors.Info("Insert file path", true)
  var path string
  fmt.Scanln(&path)
  err := os.Remove(path)
  
  if(err != nil) {
    colors.Error(fmt.Sprintf("✘ %v",err), true)
    return err
  } else {
    colors.Succes(fmt.Sprintf("File %s removed", path), true)
  }
  return nil
}
