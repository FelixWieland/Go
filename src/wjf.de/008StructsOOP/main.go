package main

import "fmt"

/*
Structs sind Typisierte Gruppierungen von Functionen, Variablen und Feldern --> OOP
*/

type auto struct {
	reader int
	tueren int
	drive  bool
}

func constAuto() auto {
	var a auto
	a.reader = 4
	a.tueren = 5
	a.drive = false

	return a
}

func (a *auto) start() {
	a.drive = true
}

func main() {
	mycar := constAuto()
	mycar.start()
	fmt.Println(mycar.drive)
}
