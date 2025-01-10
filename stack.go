package stack

import (
	"errors"
)

/*
Just a wrapper around the empty interface (essentially accept any type)
*/
type StackElement interface {
	interface{}
}

type StackImpl[SE StackElement] interface {
	Pop() SE
	Push(se SE)
}

type Stack[SE StackElement] struct {
	data []SE
}

/*
Push new element onto the top of the stack
*/
func (s *Stack[SE]) Push(se SE) {
	s.data = append(s.data, se)
}

/*
Pop the top element off the stack, if in elements present then return an error
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
