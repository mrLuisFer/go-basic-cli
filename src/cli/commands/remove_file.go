package commands

import (
	"fmt"
	"os"
)

func RemoveFile() error {
  fmt.Println("Insert file path")
  var path string
  fmt.Scanln(&path)
  err := os.Remove(path)
  
  if(err != nil) {
    fmt.Println(err)
    return err
  }
  return nil
}