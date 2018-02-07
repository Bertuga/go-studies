package main

import "fmt"

//Exercise 3: Take what you did in exercise to write the for loop (C1E1.1) and extend it a bit. Put the body of the for loop - the fmt.Printf - in a separate function.
func main() {
	for i := 0; i < 10; i++ {
		show(i)
	}
}

func show(i int) {
	fmt.Println(i)
}
