package class3

import "fmt"

type Kg float32
type Cm float32

type Person struct {
	name       string
	age        int
	department string
	weight     Kg
	height     Cm
}

func (p *Person) CheckAge(age int) {
	if age < 18 || age > 99 {
		fmt.Println("INVALID AGE")
	}
	p.age = age
}

func (p Person) PritnProfile() {
	fmt.Printf("My name is %s of dept. %s\n", p.name, p.department)
	fmt.Println("Age is ", p.age)
	fmt.Printf("Weight is %v and Height is %v\n", p.weight, p.height)
}
