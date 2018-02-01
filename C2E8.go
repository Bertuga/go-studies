package main

import "fmt"

//Exercise 8: Write a simple map()-function in Go. It is sufficient for this function only to work for ints.
func main() {
	list := []int{2, 4, 6, 8, 10}
	minus := func(i int) int {
		return i - 1
	}
	fmt.Println(Map(minus, list))
}

func Map(f func(int) int, list []int) []int {
	for k, v := range list {
		list[k] = f(v)
	}
	return list
}
