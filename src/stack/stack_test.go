package stack

import "testing"

//Exercise 1.2: Write a simple unit test for this package. You should at least test that a Pop works after a Push.
func TestPush(t *testing.T) {
	stack := new(Stack)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	if len(stack.contents) != 7 {
		t.Log("Number of elements should be 7!")
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	stack := new(Stack)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.Pop() != 3 {
		t.Log("First out should be 3!")
		t.Fail()
	}
	if stack.Pop() != 2 {
		t.Log("Second out should be 2!")
		t.Fail()
	}
	if stack.Pop() != 1 {
		t.Log("Last out should be 1!")
		t.Fail()
	}
}
