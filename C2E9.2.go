package main

import "fmt"
import "strconv"

/*Exercise 9.2: Write a String method which converts the stack to a string representation. The stack in the figure could be represented as: [0:m] [1:l] [2:k] .*/

var stack [3]int
var index int = 0

func main() {
	fmt.Println(stack)
	push(2)
	push(6)
	push(16)
	fmt.Println(stack)
	push(25)
	fmt.Printf("%v", String())
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

func String() string {
	var text string
	for _, v := range stack {
		text += strconv.Itoa(v)
	}
	return text
}
