package main

import "fmt"

/*
Standart Arrays
*/

func main() {
	//Array deklarieren
	var arr [5]int
	arr[0] = 1

	for x := 0; x < len(arr); x++ {
		arr[x] = x
	}
	fmt.Println(1)
	fmt.Println(string(96))
	fmt.Println(arr[0])
}
