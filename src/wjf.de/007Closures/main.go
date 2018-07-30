package main

import "fmt"

/*
Closures in Go
*/

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	counter := intSeq()
	counter2 := intSeq()

	fmt.Println(counter2())
	fmt.Println(counter())
	fmt.Println(counter())
}
