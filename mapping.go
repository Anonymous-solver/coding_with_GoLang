package main

import "fmt"

func main(){
	countryCapitalMap := map[string] string{"Bangladesh":"Dhaka", "India" : "Dellhi", "Japan" : "Tokya"}
	
	for country := range countryCapitalMap{
		fmt.Println("Capital of ", country, "is ", countryCapitalMap[country])
	}

	fmt.Println("\n\n")

	for country, capital := range countryCapitalMap{
		fmt.Println("Capital of ", country, "is ", capital)
	}
}