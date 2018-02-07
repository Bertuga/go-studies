package main

import "fmt"

/*Exercise 9.1: Create a simple stack which can hold a fixed number of ints. It does not have to grow beyond this limit.
Define push -- put something on the stack -- and pop -- retrieve something from the stack -- functions.
The stack should be a LIFO (last in, first out) stack.*/

var stack [3]int
var index int = 0

func main() {
	fmt.Println(stack)
	push(2)
	push(6)
	push(16)
	fmt.Println(stack)
	push(25)
	fmt.Println(stack)
	fmt.Println(pop())
	fmt.Println(stack)
	push(25)
	fmt.Println(pop())
	fmt.Println(pop())
	fmt.Println(pop())
	fmt.Println(stack)
}

func push(v int) bool {
	if index == len(stack) {
		return false
	} else {
		stack[index] = v
		index += 1
		return true
	}
}

func pop() int {
	index -= 1
	i := stack[index]
	stack[index] = 0
	return i
}
