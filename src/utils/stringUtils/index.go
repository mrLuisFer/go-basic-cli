package stringutils

import "strings"

func ToLowerCase (text string) string {
	var textFormatted = strings.ToLower(text)
	return textFormatted
}
