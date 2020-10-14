package main

import "fmt"

type Stack struct {
	intSlice []int
}

func main() {
	var s *Stack = new(Stack)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
}

// We take an integer and add to the top of the stack
func (s *Stack) Push(i int) {
	s.intSlice = append(s.intSlice, i)
}

// Peek returns the top item from the stack
func (s *Stack) Peek() int {
	return s.intSlice[len(s.intSlice)-1]
}

// Pop removes and returns the top item
func (s *Stack) Pop() int {
	top := s.Peek()
	s.intSlice = s.intSlice[:len(s.intSlice)-1]

	return top
}
