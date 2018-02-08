package main

import "fmt"
import "rpc"

//Exercise 2: Create a reverse polish calculator. Use your stack package.
func main() {
	c := rpc.NewCalculator()
	c.Add(5)
	c.Add(1)
	c.Add(2)
	c.Calc("+")
	c.Add(4)
	c.Calc("*")
	c.Calc("+")
	c.Add(3)
	c.Calc("-")
	fmt.Println(c.Get())
}
