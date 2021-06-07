package main

import (
	"fmt"
	"time"
)

func main() {
	const (
		layoutEN = "2006-01-02T15:04:05-0700"
		layoutBD = "2 January 2006 3:04 PM"
	)
	date := "4 June 2021"
	t, _ := time.Parse(layoutBD, date)
	// fmt.Println(t)
	fmt.Println(t.Format(layoutEN))
	fmt.Println(t.Format(time.RFC3339))
	// fmt.Println(date)
}
