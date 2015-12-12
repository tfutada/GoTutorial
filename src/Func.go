package main

import . "fmt"

func main() {

	x := 3
	ans := Inc(x)
	Printf("%d + 1 = %d \n", x, ans)

	z := 5
	IncPtr(&z) // pass the pointer to z
	Printf("z = %d \n", z)
}

func Inc(val int) int {
	val +=1 // val is a local variable
	return val
}

func IncPtr(val *int) {
	*val = *val + 1 // increment z
}
