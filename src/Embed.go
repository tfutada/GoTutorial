package main

import (
	. "fmt"
	"encoding/json"
	"os"
)

type Person struct {
	First string
	Last string
	Age int
}

type Developer struct {
	Person
	Skill []string
}
type Manager struct {
	Person
	Staff []Developer
}

func main() {

	tom := Developer{
		Person{"Tom", "Sandel", 21},
		[]string{"Java", "Scala"},
	}

	david := Manager{
		Person{"David", "Sandel", 36},
		[]Developer{tom},
	}

	Println("Tom   = ", tom)
	Println("David = ", david)


	enc := json.NewEncoder(os.Stdout)
	enc.Encode(david)
}








