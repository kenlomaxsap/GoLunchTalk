// Running logic in a seperate Go Routine, with the go key word
// How to wait for completion?
package main

import (
	"fmt"
)

func aNormalFunction() { //.. but executed as a GoRoutine,.
	for i := 0; i < 1000; i = i + 1 {
		fmt.Printf("in f()..   %v \n", i)
	}
}

func main() {
	fmt.Println("Before go f()")
	go aFunction()              // <- "Go" Keyword will start the function as a goroutine, to run at the *same time* as the rest of the main function
	fmt.Println("After go f()") // We might need to sleep for some time..
}
