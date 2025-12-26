// Package stack provides a simple LIFO stack implementation.
package stack

import (
	"fmt"
	"strings"
)

type Stack[T any] struct {
	items []T
}

type Stackable interface {
	Add(any)
	Pop() any
	Peek() any
	Len() int
	String() string
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Add(i T) {
	s.items = append(s.items, i)
}

func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}

	last := len(s.items) - 1
	item := s.items[last]
	s.items = s.items[:last]

	return item
}

func (s *Stack[T]) Peek() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func (s *Stack[T]) String() string {
	str := ""
	for _, item := range s.items {
		str += fmt.Sprintf("%v ", item)
	}
	return strings.ReplaceAll(str, " ", "")
}
