package logger

import (
	"fmt"

	"github.com/fatih/color"
)

func Error(message string) {
	fmt.Println(color.New(color.BgRed).Sprint("Error") + " " + message)
}

func Info(module string, message string) {
	fmt.Println(color.New(color.BgHiBlue).Sprint(module) + " " + message)
}

func Warning(module string, message string) {
	fmt.Println(color.New(color.BgHiYellow).Sprint(module) + " " + message)
}
