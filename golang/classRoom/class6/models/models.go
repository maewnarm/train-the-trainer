package models

/*type Machine struct {
	ID     string
	Name   string
	Status string
}

var machines = []Machine{
	{ID: "M1", Name: "Machine1", Status: "Green"},
	{ID: "M2", Name: "Machine2", Status: "Yellow"},
}*/

type EmpList struct {
	Lists []Employee
}

type Employee struct {
	Name       string
	Department Dept
	Grade      int
}

type Dept struct {
	Name         string
	Section      string
	Section_code string
}

var Lists = []Employee{
	{
		Name: "Voramet",
		Department: Dept{
			Name:         "PE",
			Section:      "PE2",
			Section_code: "4674",
		},
		Grade: 8,
	},
	{
		Name: "Arraya",
		Department: Dept{
			Name:         "QA",
			Section:      "QA-New model",
			Section_code: "4811",
		},
		Grade: 8,
	},
}

func GetAllEmployee() ([]Employee, error) {
	return Lists, nil
}

func AddNewEmployee(e Employee) error {
	Lists = append(Lists, e)
	return nil
}

/*func AddNewMachine(m Machine) error {
	machines = append(machines, m)
	return nil
}*/
