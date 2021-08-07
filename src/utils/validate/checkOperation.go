package validate

// Verifies the operation and returns the result or an error
func CheckOperation(num1 int, num2 int, operation string) int {
	var result int

	if operation == "sum" {
		result = num1 + num2
	}

	return result
}
