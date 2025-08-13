/* Sat Feb 15 08:28:56 PM IST 2025 */
/* By: Victor Nammour */
package ds

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	p []T // stack slice
}

func (s *Stack[T]) Init(size uint) {
	s.p = make([]T, 0, size)
}

func (s *Stack[T]) Push(b T) {
	s.p = append(s.p, b)
}

func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if len(s.p) == 0 {
		return zero, errors.New("stack underflow")
	}
	b := s.p[len(s.p)-1]
	s.p = s.p[:len(s.p)-1]
	return b, nil
}

func (s *Stack[T]) Top() (T, bool) {
	var zero T
	if len(s.p) == 0 {
		return zero, false
	}
	return s.p[len(s.p)-1], true
}

func (s *Stack[T]) Len() int {
	return len(s.p)
}

/*func (s *Stack[T]) Error() string {
	return "stack underflow"
}*/

func (s *Stack[T]) Print() {
	fmt.Printf("(index:value): ")
	for i := len(s.p) - 1; i >= 0; i-- {
		fmt.Printf("(%d:%v)", i, s.p[i])
		if i > 0 {
			fmt.Printf(", ")
		}
		if i == 0 {
			fmt.Printf("\n")
		}
	}
}
