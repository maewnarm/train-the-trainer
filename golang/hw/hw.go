package hw

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/fatih/color"
	//"test-go/shapes"
)

func Mainhw2() {
	contFlag := "y"
	arrApp := [7]string{
		"0: exit",
		"1: FIBONACCI",
		"2: FizzBuzz",
		"3: Star inc-down-right",
		"4: Star inc-down-left",
		"5: Star diamond",
		"6: Tally - count char in string"}
	for {
		color.Blue("Applications list")
		color.Blue(strings.Join(arrApp[:], "\n"))
		var selApp, input1, input2 int
		var inpTxt1, inpTxt2 string
		color.Green("choose application :")
		fmt.Scan(&selApp)
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
			inpLine := 0
			fmt.Print("number of lines :")
			fmt.Scanln(&inpLine)
			drawstar(selApp, inpLine)
		case 6:
			fmt.Print("find text = ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				inpTxt1 = scanner.Text()
			}
			fmt.Print("in text = ")
			if scanner.Scan() {
				inpTxt2 = scanner.Text()
			}
			tallyTxt(inpTxt1, inpTxt2)
		case 0:
			return
		}
		fmt.Print("want to continue ? (y/n) :")
		fmt.Scanln(&contFlag)
		if contFlag != "y" {
			break
		}
	}
}

func fibo(num1 int, num2 int) {
	fmt.Printf("%v %v", num1, num2)
	fiboNum := 0
	newNum1 := num1
	newNum2 := num2
	for i := 0; i < 10; i += 1 {
		fiboNum = newNum1 + newNum2
		fmt.Printf(" %v", fiboNum)
		newNum1 = newNum2
		newNum2 = fiboNum
	}
	fmt.Println("")
	fmt.Println("FIBONACCI end")
}

func fizzbuzz(num1 int, num2 int) {
	for i := num1; i <= num2; i++ {
		if i%3 == 0 {
			fmt.Printf("%v Fizz\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%v Buzz\n", i)
		} else {
			fmt.Printf("%v\n", i)
		}
	}
}

func drawstar(mode int, line int) {
	switch mode {
	case 3:
		for i := 0; i < line; i += 1 {
			for j := 0; j < i+1; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	case 4:
		for i := 0; i < line; i += 1 {
			for j := 0; j < line-i-1; j++ {
				fmt.Print(" ")
			}
			for k := 0; k < i+1; k++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	case 5:
		if line%2 == 0 { //even number
			aboveLine := math.Floor(float64(line)/2) - 1
			for i := 0; i < int(aboveLine); i += 1 {
				for j := 0; j < int(aboveLine)-i; j += 1 {
					fmt.Print(" ")
				}
				for k := 0; k < (2*i)+1; k += 1 {
					fmt.Print("*")
				}
				fmt.Println()
			}
			//print center
			for i := 0; i < 2; i++ {
				for j := 0; j < line-1; j += 1 {
					fmt.Print("*")
				}
				fmt.Println()
			}
			for i := int(aboveLine) - 1; i >= 0; i -= 1 {
				for j := 0; j < int(aboveLine)-i; j += 1 {
					fmt.Print(" ")
				}
				for k := 0; k < (2*i)+1; k += 1 {
					fmt.Print("*")
				}
				fmt.Println()
			}
		} else { //even number
			aboveLine := math.Floor(float64(line) / 2)
			for i := 0; i < int(aboveLine); i += 1 {
				for j := 0; j < int(aboveLine)-i; j += 1 {
					fmt.Print(" ")
				}
				for k := 0; k < (2*i)+1; k += 1 {
					fmt.Print("*")
				}
				fmt.Println()
			}
			//print center
			for i := 0; i < line; i += 1 {
				fmt.Print("*")
			}
			fmt.Println()
			for i := int(aboveLine) - 1; i >= 0; i -= 1 {
				for j := 0; j < int(aboveLine)-i; j += 1 {
					fmt.Print(" ")
				}
				for k := 0; k < (2*i)+1; k += 1 {
					fmt.Print("*")
				}
				fmt.Println()
			}
		}
	}

}

func tallyTxt(findTxt string, fromTxt string) {
	findLen := len(findTxt)
	txtLen := len(fromTxt)
	cnt := 0
	if findLen < txtLen {
		for i := 0; i < txtLen-findLen+1; i += 1 {
			subTxt := fromTxt[i : i+findLen]
			if subTxt == findTxt {
				cnt += 1
			}
		}
		fmt.Printf("Found '%s' = %v times\n", findTxt, cnt)
	}
}
