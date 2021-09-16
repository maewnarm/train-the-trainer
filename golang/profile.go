package main

import (
	"fmt"

	"github.com/fatih/color"
)

func printProfile(name string, job string) {
	color.New(color.FgBlue).Print("Hello, " + name + " You work as : ")
	if job == "SO" {
		color.New(color.FgRed).Add(color.BgGreen).Println(job)
	} else if job == "Manager" {
		color.New(color.FgBlue).Add(color.BgGreen).Println(job)
	}
	fmt.Print("a")
}
