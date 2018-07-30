package main

import "fmt"

/*
Just a dictionary
*/

func main() {
	m := make(map[string]int)
	m["t1"] = 1
	fmt.Println(m["t1"])
}
