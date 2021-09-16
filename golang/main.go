package main

import (
	"fmt"

	"github.com/fatih/color"
	//"test-go/shapes"
)

func main() {
	arrApp := [6]string{
		"1: FIBONACCI",
		"2: FizzBuzz",
		"3: Star inc-down-right",
		"4: Star inc-down-left",
		"5: Star diamond",
		"6: Tally - count char in string"}
	color.Green("choose application :")
	var selApp, input1, input2 int
	var inpTxt1, inpTxt2 string
	fmt.Scanln(&selApp)
	fmt.Println("you choose application : " + arrApp[selApp])
	switch selApp {
	case 1:
		fmt.Print("input 1 = ")
		fmt.Scanln(&input1)
		fmt.Print("input 2 = ")
		fmt.Scanln(&input2)
		fibo(input1, input2)
	case 2:
		fmt.Print("input 1 = ")
		fmt.Scanln(&input1)
		fmt.Print("input 2 = ")
		fmt.Scanln(&input2)
		fizzbuzz(input1, input2)
	case 3, 4, 5:
		drawstar(selApp)
	case 6:
		fmt.Print("find text = ")
		fmt.Scan(&inpTxt1)
		fmt.Print("in text = ")
		fmt.Scan(&inpTxt2)
		tallyTxt(inpTxt1, inpTxt2)
	}
}

func fibo(num1 int, num2 int) {

}

func fizzbuzz(num1 int, num2 int) {

}

func drawstar(mode int) {

}

func tallyTxt(findTxt string, fromTxt string) {

}
