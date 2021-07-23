package colors

import (
	colors "github.com/fatih/color"
)

// Paint yellow the text for the console
func Warn(msg string, newLine bool) {
	yellow := colors.New(colors.FgYellow).Add(colors.Bold).Set()
	if(newLine){
		yellow.Printf("%v \n", msg)
	} else {
		yellow.Print(msg)
	}
}

// Paint blue the text for the console
func Info(msg string, newLine bool) {
	blue := colors.New(colors.FgBlue).Add(colors.Bold).Set()
	blue.Print(msg)
	if(newLine){
		blue.Printf("%v \n", msg)
	} else {
		blue.Print(msg)
	}
}

// Paint green the text for the console
func Succes(msg string, newLine bool) {
	green := colors.New(colors.FgGreen).Set()
	green.Print(msg)
	if(newLine){
		green.Printf("%v \n", msg)
	} else {
		green.Print(msg)
	}
}
