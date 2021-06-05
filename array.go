package main

import "fmt"

func main() {
	var num int
	// fmt.Scan(&n)
	var arr [10]int
	for i := 0; i <= 9; i++ {
		fmt.Scan(&num)
		arr[i] = num
	}
	fmt.Print(arr)
}
