package colors

import (
	colors "github.com/fatih/color"
)

// Paint yellow the text for the console
func Warn(msg interface {}, newLine bool) {
	yellow := colors.New(colors.FgYellow).Add(colors.Bold).Set()
	if(newLine){
		yellow.Printf("%v \n", msg)
	} else {
		yellow.Print(msg)
	}
}

// Paint blue the text for the console
func Info(msg interface {}, newLine bool) {
	blue := colors.New(colors.FgBlue).Add(colors.Bold).Set()
	if(newLine){
		blue.Printf("%v \n", msg)
	} else {
		blue.Print(msg)
	}
}

// Paint green the text for the console
func Succes(msg interface {}, newLine bool) {
	green := colors.New(colors.FgHiGreen).Add(colors.Bold).Set()
	if(newLine){
		green.Printf("%v \n", msg)
	} else {
		green.Print(msg)
	}
}

// Paint red text in case there is an error
func Error(msg interface {}, newLine bool){
	red := colors.New(colors.FgHiRed).Set()
	if(newLine){
		red.Printf("%v \n", msg)
	} else {
		red.Print(msg)
	}
}
