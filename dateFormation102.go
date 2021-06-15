// Golang program to illustrate the usage of
// time.Parse() function
  
// Including main package
package main
  
// Importing fmt and time
import "fmt"
import "time"
  
// Calling main
func main() {
  
    // Declaring layout constant
    const layout = "2 Jan 2006 3:04 PM"
  
    // Calling Parse() method with its parameters
    tm, _ := time.Parse(layout, "2 Jun 2021 12:31 AM")
    fmt.Println(tm.Format(time.RFC3339Nano))
}

