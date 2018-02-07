package main

import "fmt"

//Exercise 4: Write a function that takes an int value and gives that many terms of the Fibonacci sequence.
func main() {
	i := 10
	fmt.Println(fibonacci(i))
}

func fibonacci(i int) []int {
	terms := make([]int, i)
	for k, _ := range terms {
		switch k {
		case 0, 1:
			terms[k] = 1
		default:
			terms[k] = terms[k-1] + terms[k-2]
		}
	}
	return terms
}
