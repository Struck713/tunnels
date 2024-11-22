package main

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintError(message string) {
	fmt.Println(color.RedString("Error:") + " " + message)
}

func PrintInfo(message string) {
	fmt.Println(color.CyanString("Tunnel:") + " " + message)
}
