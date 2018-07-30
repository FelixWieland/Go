package main

import (
	"fmt"
	"time"
)

/*
Einer der Hauptgründe für Go. Threads bzw Concurrent functions
*/

func test() {
	for i := 0; i < 6; i++ {
		fmt.Println(i)
	}
}

func main() {
	go test()
	go fmt.Println("demo")
	time.Sleep(1000)
}
