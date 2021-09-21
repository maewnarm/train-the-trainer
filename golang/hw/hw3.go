package hw

import (
	"fmt"
)

type Kg int
type Cm int

type Person struct {
	Name       string
	Age        int
	Department string
	Weight     Kg
	Height     Cm
}

func Mainhw3() {
	var fib func(n int, m string) int
	fib = func(n int, m string) int {
		fmt.Println(n)
		fmt.Println(m)
		if n < 2 {
			return n
		}
		return fib(n-1, "-1") + fib(n-2, "-2")
	}
	fmt.Println(fib(5, "-0"))
}
