// Channels (that block) are used to communicate between GoRoutines..
// You can (and need to) share channels across GoRoutines.

package main

import (
	"fmt"
)

var aGlabalChannelCarryingInts = make(chan int)

func aFunction() {
	for i := 0; i < 1000; i = i + 1 {
		fmt.Printf("in f()..   %v \n", i)
	}
	//Once we are finished, send a message (answer) on aGlabalChannelCarryingInts
	aGlabalChannelCarryingInts <- 123
}

func main() {
	fmt.Println("Before go f()")
	go aFunction() // <- Go Keyword to create a goroutine

	fmt.Println("After go f()")

	// Blocking call.. will wait for message on aBlockingChannel
	i := <-aGlabalChannelCarryingInts
	fmt.Println("Value received ", i)
}
