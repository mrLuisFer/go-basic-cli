package stringutils

import "strings"

func ToLowerCase (text string) string {
	var textFormatted string = strings.ToLower(text)
	return textFormatted
}
