package main

import (
	"fmt"
	classroom "test-golang/classRoom"
	"test-golang/hw"

	"github.com/fatih/color"
)

func main() {
	color.Green("Welcome to HW train the trainer apps")
	for {
		app := ""
		fmt.Print("input app's name :")
		fmt.Scanln(&app)
		switch app {
		case "class3":
			classroom.Mainclass(3)
		case "hw2":
			hw.Mainhw2()
		case "hw3":
			hw.Mainhw3()
		default:
			color.Red("not found this app")
		}
	}
}
