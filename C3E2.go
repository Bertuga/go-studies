package main

import "fmt"

//Exercise 2: Write a function that performs a bubble sort on a slice of ints.
func main() {
	list := []int{5, 12, 9, 6, 87, 2, 2, 4}
	piece := make([]int, 8)
	copy(piece, list)
	fmt.Println(bubblesort(list))
}

func bubblesort(list []int) (sorted []int) {
	sorted = list
	for x, _ := range sorted {
		for y := x + 1; y < len(sorted); y++ {
			if sorted[x] > sorted[y] {
				swap := sorted[x]
				sorted[x] = sorted[y]
				sorted[y] = swap
			}
		}
	}
	return sorted
}
