package main

import "fmt"

//Exercise 4.2: Generalize the function from above and create a plusX(x) which returns functions that add x to an integer.
func main() {
	p := plusX(3)
	fmt.Println(p(2))
}

func plusX(x int) func(int) int {
	return func(i int) int { return i + x }
}
