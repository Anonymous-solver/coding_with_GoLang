package main

import "fmt"

func main() {
	var fname string
	fmt.Print("Enter your first name :")
	fmt.Scan(&fname)
	var lname string
	fmt.Print("Enter your last name :")
	fmt.Scan(&lname)
	fmt.Print("I am " + fname + " " + lname)
}
