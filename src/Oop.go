package main

import (
	. "fmt"
)

type Person struct {
	First string
	Last string
	Age int
}

func (p Person) Fullname() string {
	return p.First + " " + p.Last
}
func (p *Person) SetAge(a int) {
	p.Age = a
}

func main() {

	tom := Person{Age: 28, Last: "Gates", First: "Tom"}
	Println(tom.Fullname())

	tom.SetAge(32)
	Println(tom)

}
