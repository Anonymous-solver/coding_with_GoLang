package main

import (
	"fmt"
	"time"
)

func main() {
	const (
		layoutEN = "2006-01-02"
		layoutBD = "2 January 2006"
	)
	date := "4 May 2021"
	t, _ := time.Parse(layoutBD, date)
	fmt.Println(t)
	fmt.Println(t.Format(layoutEN))
}
