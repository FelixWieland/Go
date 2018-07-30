package main

import (
	"fmt"
	"strconv"
	"time"
)

func demo(channel chan string) {
	arr := []int{1, 2, 3, 4, 56, 6, 78, 9}
	time.Sleep(time.Second * 5)
	for x := range arr {
		channel <- strconv.Itoa(arr[x])
	}
	close(channel) //close to avoid deadlock
}

func printThis(channel2 chan string) {
	time.Sleep(time.Second * 5)
	channel2 <- "Passd So"
}

func main() {

	channel := make(chan string)
	channel2 := make(chan string)
	fmt.Println(time.Now())
	go demo(channel)       //starts new goroutine -> 5 sec
	go printThis(channel2) //starts new goroutine -> 5 sec

	for x := range channel { //iterate over channel till closed
		fmt.Println(x)
	}

	select {
	case msg := <-channel2: //wait till received messages
		fmt.Println(msg)
	}

	//Reached end in only ~5sec and not in 10sec because both run concurrently

	defer fmt.Println(time.Now())

}
