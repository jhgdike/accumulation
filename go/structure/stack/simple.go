package stack

import (
	"sync"
)

type ItemStack struct {
	items []Item
	lock sync.RWMutex
	size int
}

func NewItemStack(size int) Stack {
	s := &ItemStack{
		items: []Item{},
		lock: sync.RWMutex{},
		size: size,
	}
	return s
}

func (s *ItemStack) Push(v Item) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.items) >= s.size {
		return fullError
	}
	s.items = append(s.items, v)
	return nil
}

func (s *ItemStack) Pop() (Item, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.items) == 0 {
		return nil, emptyError
	}
	item := s.items[len(s.items) - 1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}
