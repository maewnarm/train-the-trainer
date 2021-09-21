package class3

import (
	"fmt"
	"os"
)

func Mainclass3() {
	fmt.Println("class 3")
	personel := Person{}
	fmt.Print("Enter name :")
	fmt.Scanln(&personel.name)
	fmt.Print("Enter age :")
	fmt.Scanln(&personel.age)
	fmt.Print("Enter weight :")
	fmt.Scanln(&personel.weight)
	fmt.Print("Enter height :")
	fmt.Scanln(&personel.height)
	fmt.Print("Enter department :")
	fmt.Scanln(&personel.department)
	personel.CheckAge(personel.age)
	personel.PritnProfile()
	file, _ := os.Create("output.txt")
	defer file.Close()
	fmt.Fprintln(file, "My name is ", personel.name, " of dept. ", personel.department)
	fmt.Fprintln(file, "Age is ", personel.age)
	fmt.Fprintf(file, "Weight is %v and Height is %v", personel.weight, personel.height)
}
