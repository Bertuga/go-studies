package main

import "fmt"

//Exercise 1.2: Rewrite the loop from 1.1 to use goto. The keyword for may not be used.
func main() {
	i := 0
Loop:
	fmt.Println(i)
	i++
	if i < 10 {
		goto Loop
	}
}
