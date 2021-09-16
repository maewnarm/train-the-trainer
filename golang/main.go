package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/fatih/color"
	//"test-go/shapes"
)

func main() {
	rand.Seed(time.Now().UnixMicro())
	n := rand.Intn(100)
	input := ""
	maxTry := 5
	fmt.Println("Generated random number")
	for try := 0; try < maxTry; try += 1 {
		fmt.Printf("Input guessed number (%v/%v):", try+1, maxTry)
		fmt.Scanln(&input)
		inp, err := strconv.Atoi(input)
		if err != nil {
			panic("convert error")
		}
		if inp < n {
			color.Yellow("TOO LOW")
		} else if inp > n {
			color.Red("TOO HIGH")
		} else {
			color.Green("CORRECT !!")
			return
		}
	}
}
