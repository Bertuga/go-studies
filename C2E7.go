package main

import "fmt"

//Exercise 7: Write a function that finds the maximum value in an int slice ([]int).
func main() {
	list := []int{1, 2, 16, 8, 12, 6}
	fmt.Println(max(list))
}

func max(list []int) int {
	max := 0
	for k, v := range list {
		if list[k] > max {
			max = v
		}
	}
	return max
}
