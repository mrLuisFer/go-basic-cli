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
	var error bool

	colors.Info("Insert a number: ", false)
	fmt.Scanln(&numb1)

	colors.Info("Insert second number: ", false)
	fmt.Scanln(&numb2)

	if (numb1 <= 0 && numb2 <= 0) {
		error = true
	}
	
	typeNumber1 := reflect.TypeOf(numb1).Kind()
	typeNumber2 := reflect.TypeOf(numb2).Kind()

	if (typeNumber1 == reflect.Int && typeNumber2 == reflect.Int && !error) {
		result := numb1 + numb2
		colors.Succes("The result is:", false)
		fmt.Printf(" %v \n", result)
	} else {
		colors.Error("Insert a valid number!", false)
		error = true
	}
}
