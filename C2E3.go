package main

import "fmt"

/*Exercise 3: Solve this problem, called the Fizz-Buzz problem:
Write a program that prints the numbers from 1 to 100.
But for multiples of three print, "Fizz" instead of the number, and for multiples of five, print "Buzz".
For numbers which are multiples of both three and five, print "FizzBuzz".*/

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
