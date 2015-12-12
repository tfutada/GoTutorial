package main

import . "fmt"

func main() {

	c := make(chan int)

	go sumsUp(c, 1, 5)
	go sumsUp(c, 5, 8)
	go sumsUp(c, 8, 10)

	ret1 := <- c // receive the first sumsUp call
	ret2 := <- c // receive the first sumsUp call
	ret3 := <- c // receive the first sumsUp call

	Printf("%d + %d + %d = %d \n",
		ret1, ret2, ret3, ret1 + ret2 + ret3 )
}

func sumsUp(c chan int, from int, to int) {

	sum := 0
	for i:=from; i < to; i++ {
		sum += i
	}

	c <- sum // send the result
}
