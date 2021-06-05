package main

import "fmt"

type Students struct {
	name   string
	dept   string
	std_id int32
}

func main() {
	var student1 Students
	var student2 Students

	student1.name = "Anik Barua"
	student1.dept = "CSE"
	student1.std_id = 1001

	student2.name = "Masud Rana"
	student2.dept = "EEE"
	student2.std_id = 2001

	fmt.Printf("%s\n", student1.name)
	fmt.Printf("%s\n", student1.dept)
	fmt.Printf("%d\n", student1.std_id)

	fmt.Print("\n\n")

	fmt.Printf("%s\n", student2.name)
	fmt.Printf("%s\n", student2.dept)
	fmt.Printf("%d\n", student2.std_id)
}
