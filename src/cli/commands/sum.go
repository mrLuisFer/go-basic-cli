package commands

import (
	"fmt"

	colors "github.com/mrLuisFer/go-basic-cli/src/utils/colors"

	validate "github.com/mrLuisFer/go-basic-cli/src/utils/validate"
)

func Sum(){
	colors.Warn("Insert two numbers!", true)

	var numb1 int
	var numb2 int

	colors.Info("Insert a number: ", false)
	fmt.Scanln(&numb1)

	colors.Info("Insert second number: ", false)
	fmt.Scanln(&numb2)

	validate.ValidateTwoNumbers(numb1, numb2, "sum")
}
