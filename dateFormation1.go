package main

import (
	"fmt"
)

func main() {

	// this function returns the present time
	current_time := "4 May 2021 12:31"

	fmt.Println(current_time)

	// individual elements of time can
	// also be called to print accordingly
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		current_time.Year(), current_time.Month(), current_time.Day(),
		current_time.Hour(), current_time.Minute(), current_time.Second())

	// formatting time using
	// custom formats
	fmt.Println(current_time.Format("2006-01-02 15:04:05"))
	fmt.Println(current_time.Format("2006-January-02"))
	fmt.Println(current_time.Format("2006-01-02 3:4:5 pm"))
}
