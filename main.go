package main

import (
	cli "github.com/mrLuisFer/go-basic-cli/src/cli"
)

func main() {
	err := cli.Init()

  if err != nil {
    panic(err)
  }
}
