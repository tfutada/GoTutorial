package main

import . "fmt"

func main() {

	defer func(val string) {
		Println(val)
	}("Hello World!")

	Println("Start!")
}