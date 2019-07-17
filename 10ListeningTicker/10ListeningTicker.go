// ticker channel

package main

import "time"
import "fmt"

var ticker = time.NewTicker(500 * time.Millisecond)

func f() {
	for t := range ticker.C {
		fmt.Println("Tick at", t)
	}
}

func main() {

	go f()

	// Tickers can be stopped like timers. Once a ticker
	// is stopped it won't receive any more values on its
	// channel. We'll stop ours after 1600ms.

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
