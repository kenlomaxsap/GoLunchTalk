// Timeouts
package main

import "time"
import "fmt"

var ch1 = make(chan string)
var ch2 = make(chan string)

func oneSecondPause() {
	time.Sleep(1 * time.Second)
	ch1 <- "result 1"
}

func twoSecondPause() {
	time.Sleep(2 * time.Second)
	ch2 <- "result 2"
}
func main() {

	go oneSecondPause()
	select {
	case res := <-ch1:
		fmt.Println(res)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("timeout 1")
	}

	go twoSecondPause()

	select {
	case res := <-ch2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
