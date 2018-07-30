package main

import "fmt"

/*
Ermöglicht es mehrere Werte aus einer Funktion zurückzugeben
*/

func vals() (int, int) {
	return 3, 6
}

func main() {
	a, b := vals()
	fmt.Println(a, b)
}
