package main

import "fmt"

func main() {
	// a string array
	var country = []string{"Japan", "Vietnam", "US",}

	for idx, val := range country {
		fmt.Println("country ", idx, ":", val)
	}

	// a hash map
	person := map[string]int {
		"David" : 32,
		"Tom" : 18,
		"Bob" : 27,
	}

	var kids, adults []string

	for k, v := range person {
		if v > 20 { // outputs over 20
			adults = append(adults, k)
		} else {
			kids = append(kids, k)
		}
	}

	fmt.Printf("Adults : %s", adults)
}

