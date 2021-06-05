package main

import "fmt"

func main() {
	var day int
	fmt.Scan(&day)
	switch day {
	case 1:
		fmt.Print("hey! Today is Sunday")
	case 2:
		fmt.Print("Hi! Today is Monday")
	case 3:
		fmt.Print("Yup! Today is Tuesday")
	case 4:
		fmt.Print("Uh! Today is Wednesday")
	case 5:
		fmt.Print("Yes! Today is Thursday")
	case 6:
		fmt.Print("Yeh! Today is Friday")
	case 7:
		fmt.Print("Whats up! Today is Saturday")
	default:
		fmt.Print("Oops! It doesn't mean any day")
	}
}
