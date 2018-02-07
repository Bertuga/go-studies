package main

import "fmt"
import "stack"

/*Exercise 1.1: See the Stack exercise. In this exercise we want to create a separate package for that code.
Create a proper package for your stack implementation, Push, Pop and the Stack type need to be exported.*/
func main() {
	stack := new(stack.Stack)
	stack.Push(2)
	stack.Push(3)
	stack.Push(5)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	stack.Push(55)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
