package main

import "fmt"

//Exercise 2: Write code to calculate the average of a float64 slice. In a later exercise you will make it into a function.
func main() {
	list := []float64{1.5, 0.5, 3.5, 7.5, 9.5}
	var piece = make([]float64, 4)
	copy(piece, list[1:])
	var total float64 = 0
	var qty float64 = 0
	for _, v := range piece {
		total += v
		qty++
	}
	fmt.Println(total / qty)
}
