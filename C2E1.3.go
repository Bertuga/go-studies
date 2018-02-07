package main

import "fmt"

//Exercise 1.3: Rewrite the loop again so that it fills an array and then prints that array to the screen.
func main() {
	var list [10]int
	for i, _ := range list {
		list[i] = i
	}
	fmt.Println(list)
}
