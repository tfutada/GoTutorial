package main

import (
	"fmt"
)

type Person struct {
	first string
	last string
}

func (p Person) Fullname() string {
	return p.first + " " + p.last
}

type Myname interface {
	Fullname() string
}

// takes an interface Myname
func WhoAreyou(my Myname) {
	fmt.Println(my.Fullname())
}

func main() {
	tom := Person{"Tom", "Jobs"}
	fmt.Println(tom.Fullname())

	WhoAreyou(tom)
}

