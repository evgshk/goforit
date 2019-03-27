package main

import "fmt"

func main() {
	// Define map
	countries := make(map[string]string)

	// Assign key-values
	countries["USA"] = "Washington"
	countries["Russia"] = "Moscow"
	countries["France"] = "Paris"

	fmt.Println(countries)
	fmt.Println(countries["France"])
	fmt.Println(len(countries))

	// Delete from map
	delete(countries, "France")
	fmt.Println(countries)

	//Declare map and add key-values
	employees := map[int]string{1: "John", 2: "Brad"}
	fmt.Println(employees)
}
