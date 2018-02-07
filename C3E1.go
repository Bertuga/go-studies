package main

import "fmt"

//Exercise 1: Write a function that calculates the average of a float64 slice.
func main() {
	list := []float64{1.5, 0.5, 3.5, 7.5, 9.5}
	piece := make([]float64, 4)
	copy(piece, list[1:])
	calc64(piece)
}

func calc64(list []float64) {
	var total float64 = 0
	var qty float64 = 0
	for _, v := range list {
		total += v
		qty++
	}
	fmt.Println(total / qty)
}
