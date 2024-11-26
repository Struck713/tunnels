package logger

import (
	"fmt"

	"github.com/fatih/color"
)

func Debug(message any) {
	fmt.Print(color.New(color.BgGreen).Sprint("Debug") + " ")
	fmt.Println(message)
}

func Error(message string) {
	fmt.Println(color.New(color.BgRed).Sprint("Error") + " " + message)
}

func Info(module string, message string) {
	fmt.Println(color.New(color.BgHiBlue).Sprint(module) + " " + message)
}

func Warning(module string, message string) {
	fmt.Println(color.New(color.BgHiYellow).Sprint(module) + " " + message)
}
