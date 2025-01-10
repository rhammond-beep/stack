package stack

import (
	"errors"
)

/*
The union type includes all of Go's natively supported types
*/
type StackElement interface {
	int | int8 | int16 | rune | int64 | string | bool | interface{}
}

type StackImpl[SE StackElement] interface {
	Pop() SE
	Push(se SE)
}

type Stack[SE StackElement] struct {
	data []SE
}

/*
Push new element onto the end of the stack
*/
func (s *Stack[SE]) Push(se SE) {
	s.data = append(s.data, se)
}

/*
Pop and element from the stack, or if no elements on stack, return
a populated error message
*/
func (s *Stack[SE]) Pop() (*SE, error) {
	if len(s.data) == 0 {
		return nil, errors.New("Stack Empty")
	}

	n := len(s.data) - 1
	element := s.data[n]
	s.data = s.data[:n]
	return &element, nil
}
