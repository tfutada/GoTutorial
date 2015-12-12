package main

import . "fmt"

func main() {

	var p *int

	x, y := 5, 10

	p = &x
	Println("*p = ", *p, "  x = ", x, "    p = ", p)

	*p = 6
	Println("*p = ", *p, "  x = ", x, "    p = ", p)

	p = &y
	Println("*p = ", *p, " y = ", y, "   p = ", p)

	t := new(int)
	Println("*t = ", *t, "             t = ", t)

	*t = 20
	Println("*t = ", *t)

	t = p
	Println("*t = ", *t, "  t = ", t, "   p = ", p)
}
