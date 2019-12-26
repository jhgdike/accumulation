package stack

import "errors"

type Item interface {}

type Stack interface {
	Push(v Item) error
	Pop() (Item, error)
}

var (
	emptyError = errors.New("stack is empty")
	fullError = errors.New("stack is full")
)
