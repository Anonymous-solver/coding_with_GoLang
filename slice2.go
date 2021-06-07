package main

import "fmt"

func main(){
	var numbers []int
	var numbers1 []int
	numbers = append(numbers, 100)
	numbers = append(numbers, 101)
	numbers = append(numbers, 102)

	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	copy(numbers1, numbers)
	fmt.Println(numbers)
	fmt.Println(numbers1)	
}