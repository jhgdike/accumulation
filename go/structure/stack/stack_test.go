package stack

import "testing"

func TestStack(t *testing.T) {
	stack := NewItemStack(5)
	for i:= 0; i < 5; i ++ {
		stack.Push(1)
	}
	err := stack.Push(6)
	if err == nil {
		t.Fatal("stack should be full, but still can be fill")
	}
	t.Log(err)
	for i := 0; i < 5; i ++ {
		v, _ := stack.Pop()
		t.Logf("stack top value: %d", v)
	}
	v, err := stack.Pop()
	if err == nil {
		t.Fatal("stack should be empty, but still can be pop")
	}
	t.Log(v)
}
