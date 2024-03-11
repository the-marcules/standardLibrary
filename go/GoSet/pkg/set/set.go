package set

import (
	"fmt"
	"slices"
	"strings"
)

type Set[T comparable] struct {
	container []T
	Stringer  func(val T) string
}

func NewSet[T comparable](initialValue ...T) *Set[T] {
	s := Set[T]{
		container: []T{},
		Stringer: func(val T) string {
			return fmt.Sprint(val)
		},
	}

	for _, value := range initialValue {
		s.Add(value)
	}

	return &s
}

func (s *Set[T]) Add(value T) {
	if !s.IsValueInSet(value) {
		s.container = append(s.container, value)
	}
}

func (s *Set[T]) Remove(value T) {
	index := s.indexOfValue(value)
	if index != -1 {
		s.container = append(s.container[:index], s.container[index+1:]...)
	}
}

func (s *Set[T]) indexOfValue(value T) int {
	return slices.Index(s.container, value)
}

func (s *Set[T]) String() string {
	var str []string

	for _, item := range s.container {
		str = append(str, s.Stringer(item))
	}

	return strings.Join(str, "\n")
}

func (s *Set[T]) IsValueInSet(value T) bool {
	return slices.Contains(s.container, value)
}

func (s *Set[T]) ItemsCount() int {
	return len(s.container)
}

func (s *Set[T]) Clear() {
	s.container = []T{}
}
