package stringutils

import "strings"

// Converts text to lowercase
func ToLowerCase (text string) string {
	var textFormatted string = strings.ToLower(text)
	return textFormatted
}

// Converts text to uppercase
func ToUpperCase (text string) string {
	var textFormatted string = strings.ToUpper(text)
	return textFormatted
}
