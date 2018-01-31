package main

import "fmt"

//Exercise 5: Write a function that takes a variable number of ints and print each integer on a separate line.
func main() {
	print_ints(1, 3, 6, 8)
	print_ints(2, 44, 66, 43, 0, 22)
}

func print_ints(ints ...int) {
	for _, v := range ints {
		fmt.Println(v)
	}
}
