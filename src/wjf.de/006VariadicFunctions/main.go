package main

import "fmt"

/*
Variadic Functions ermöglichen es das man einer function eine unbegrenzte anzahl an Parametern mitgeben kann
*/

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	x := sum(2, 4557, 6, 767, 34, 53, 34)
	fmt.Println(x)
}
