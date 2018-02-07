package main

import "fmt"

//Exercise 6.1: Write a function that returns a function that performs a +2 on integers. Name the function plusTwo.
func main() {
	p := plusTwo()
	fmt.Println(p(2))
}

func plusTwo() func(int) int {
	return func(i int) int { return i + 2 }
}
