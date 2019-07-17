// Sharing data across Go Routines..
// race conditions and atomic uint64..
// go  run  -race 5SharingDataOk/5SharingDataOk.go

package main

import (
	"fmt"
	"sync/atomic"
)

var ch = make(chan bool)

var data uint64 //  <- Some data that is referenced in both f() and main()

func f() {
	for i := 0; i < 40; i = i + 1 {
		atomic.AddUint64(&data, 1)
		fmt.Print("In f()", atomic.LoadUint64(&data))
	}
	ch <- true
}

func main() {
	go f() // <- Go Keyword to create a goroutine

	for i := 0; i < 40; i = i + 1 {
		atomic.AddUint64(&data, 1)
		fmt.Print("In f()", atomic.LoadUint64(&data))
	}

	<-ch
	fmt.Println("\n\nFinal result: ", atomic.LoadUint64(&data))
	fmt.Println("Conclusion: Sharing Data is Hard/Bad/Tricky/Dangerous/<add your own word>  ")
}
