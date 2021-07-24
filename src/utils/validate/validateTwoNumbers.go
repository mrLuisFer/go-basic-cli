package validate

import (
	"fmt"
	"reflect"

	colors "github.com/mrLuisFer/go-basic-cli/src/utils/colors"
)

func ValidateTwoNumbers(num1 int, num2 int, operation string) {
	var error bool = false
	
	if (num1 <= 0 && num2 <= 0) {
		error = true
	}
	
	typeNumber1 := reflect.TypeOf(num1).Kind()
	typeNumber2 := reflect.TypeOf(num2).Kind()
	
	if (typeNumber1 == reflect.Int && typeNumber2 == reflect.Int && !error) {
		var result int

		if (operation == "sum") {
			result = num1 + num2
		}
		colors.Succes("The result is:", false)
		fmt.Printf(" %v \n", result)
	} else {
		colors.Error("Insert a valid number!", false)
		error = true
	}	
}
