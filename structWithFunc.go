package main

import "fmt"

type Students struct {
	name   string
	dept   string
	std_id int32
}

func printStdInfo(student Students) {
	fmt.Printf("%s\n", student.name)
	fmt.Printf("%s\n", student.dept)
	fmt.Printf("%d\n", student.std_id)
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

	printStdInfo(student1)

	fmt.Print("\n\n")

	printStdInfo(student2)
}
