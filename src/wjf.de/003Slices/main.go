package main

import "fmt"

/*
Wie Arrays bloß flexibler. In kombination mit maps wie ein Dictionary in C# oder ein Array in PHP nutzbar
*/

func main() {
	//Slice
	myslice := make([]int, 7)
	for i := 0; i < 7; i++ {
		myslice[i] = i
	}
	myslice = append(myslice, 90)
	fmt.Println(myslice[7])
}
