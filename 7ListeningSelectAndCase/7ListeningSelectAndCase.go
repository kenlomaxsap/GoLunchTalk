package main

import "time"
import "fmt"

// Select and Case
var channel1Strings = make(chan string)
var channel2Ints = make(chan int)

func f1() {
	for {
		time.Sleep(1 * time.Second)
		channel1Strings <- "one"
		time.Sleep(2 * time.Second)
	}
}

func f2() {
	for {
		time.Sleep(2 * time.Second)
		channel2Ints <- 2
		time.Sleep(1 * time.Second)
	}

}

func main() {
	go f1()
	go f2()

	for {
		select {
		case msg2 := <-channel2Ints:
			fmt.Println("received on channel2Ints", msg2)
		case msg1 := <-channel1Strings:
			fmt.Println("received on channel1Strings", msg1)
		}
	}
}
