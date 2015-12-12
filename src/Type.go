package main

import (
	. "fmt"
	"time"
)

type Person struct {
	First string
	Last string
	Age int
}

func main() {

	david := Person{"David", "Jobs", 34}
	tom := Person{Age: 28, Last: "Gates", First: "Tom"}

	Printf("%s %s is %d as of %s\n",
		david.First, david.Last, david.Age, time.Now().Format("2006-01-02 3:04:05 PM"))

	Println(tom)
}
