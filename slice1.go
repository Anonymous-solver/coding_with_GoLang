package main

import "fmt"

func main(){
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Numbers: ", numbers)
	fmt.Println("Numbers[1:4] ", numbers[1:4])
	fmt.Println("Numbers[4:] ", numbers[4:])
	fmt.Println("Numbers[:4] ", numbers[:4])
}