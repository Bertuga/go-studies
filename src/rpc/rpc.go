package rpc

import "stack"

/*Just a simple implementation, there is a LOT of room for improvement*/
type calculator struct {
	numbers stack.Stack
}

func (c *calculator) Add(i int) {
	c.numbers.Push(i)
}

func (c *calculator) Calc(s string) {
	if c.numbers.Size() > 1 {
		n1 := c.numbers.Pop()
		n2 := c.numbers.Pop()
		switch s {
		case "+":
			c.numbers.Push(n2 + n1)
		case "-":
			c.numbers.Push(n2 - n1)
		case "*":
			c.numbers.Push(n2 * n1)
		case "/":
			c.numbers.Push(n2 / n1)
		}
	}
}

func (c *calculator) Get() int {
	if c.numbers.Size() == 1 {
		return c.numbers.Pop()
	} else {
		return 0
	}
}

func NewCalculator() *calculator {
	c := new(calculator)
	return c
}
