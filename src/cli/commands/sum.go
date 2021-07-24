package commands

import (
	"fmt"

	"reflect"

	colors "github.com/mrLuisFer/go-basic-cli/src/utils/colors"
)

func Sum(){
	colors.Warn("Insert two numbers!", true)

	var numb1 int
	var numb2 int

	colors.Info("Insert a number: ", false)
	fmt.Scanln(&numb1)

	colors.Info("Insert second number: ", false)
	fmt.Scanln(&numb2)

	var typeNumber1 = reflect.TypeOf(numb1).Kind()
	var typeNumber2 = reflect.TypeOf(numb2).Kind()

	if (typeNumber1 == reflect.Int && typeNumber2 == reflect.Int) {
		result := numb1 + numb2
		colors.Succes("The result is:", false)
		fmt.Printf(" %v \n", result)
	}
}
