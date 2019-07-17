// Sharing more complex data across Go Routines
// The State we care about is in one routine only
package main

import (
	"fmt"
)

type msg struct {
	name  string
	value int
}

var ch = make(chan msg)

func routineWithState() {
	var state msg // HERE IS THE STATE -  in one GoRoutine only
	for {
		state = <-ch
		fmt.Printf("In f with state  %+v \n", state)
	}
}

func main() {
	go routineWithState()

	for i := 0; i < 40; i = i + 1 {
		ch <- msg{name: "SomeName", value: i}
	}
}
