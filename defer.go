package main 

import "fmt"

func first(){
	fmt.Printf("This is the first one\n")
}

func second(){
	fmt.Printf("This is second one\n")
}

func third(){
	fmt.Printf("This is third one\n")
}

func fourth(){
	fmt.Printf("This is the fourth one\n")
}

func main(){
	defer second()
	third()
	defer first()
	defer fourth()
}