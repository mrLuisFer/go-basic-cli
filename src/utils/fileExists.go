package utils

import "os"

func FileExists (fileName string) bool {
  fileinfo, error := os.Stat(fileName)

  if os.IsNotExist(error) {
    return false
  }

  return !fileinfo.IsDir()
}
