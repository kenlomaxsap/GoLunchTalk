// Single threaded programming in Go.. - no surprises..

package main

import "fmt"

func aNormalFunction() {
	for i := 0; i < 1000; i = i + 1 {
		fmt.Printf("%v \n", i)
	}
}
func main() {
	aNormalFunction()
}
