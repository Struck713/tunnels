package logger

import (
	"fmt"

	"github.com/fatih/color"
)

func Error(message string) {
	fmt.Println(color.RedString("Error:") + " " + message)
}

func Info(message string) {
	fmt.Println(color.CyanString("Tunnel:") + " " + message)
}
