package main

import (
	"fmt"
	"test-golang/hw"

	"github.com/fatih/color"
)

func main() {
	for {
		app := ""
		color.Green("Welcome to main apps")
		fmt.Print("input app's name :")
		fmt.Scanln(&app)
		switch app {
		case "hw2":
			hw.Mainhw2()
		default:
			color.Red("not found this app")
		}
	}
}
