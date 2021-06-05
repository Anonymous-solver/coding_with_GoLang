package main

import "fmt"

func main() {

	limit := 1

	//fmt.Scan(&i, &j)

	fmt.Scan(&limit)

	for i := 1; i <= limit; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
