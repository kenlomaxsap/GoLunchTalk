// Sharing data across Go Routines..

// race conditions and atomic uint64..
// go  run  -race 4SharingDataBadly/4SharingDataBadly.go

package main

import (
	"fmt"
)

var ch = make(chan bool)

var data string //  <- Some data that is referenced in both f() and main()

func f() {
	for i := 0; i < 40; i = i + 1 {
		data = data + "f"
		fmt.Print("In f()", data)
	}
	ch <- true
}

func main() {
	go f() // <- Go Keyword to create a goroutine

	for i := 0; i < 40; i = i + 1 {
		data = data + "g"
		fmt.Print("In f()", data)

	}

	<-ch
	fmt.Println("\n\nFinal result: ", data)
	fmt.Println("Conclusion: Sharing Data is Hard/Bad/Tricky/Dangerous/<add your own word>  ")
}
