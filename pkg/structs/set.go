package structs

import (
	"fmt"
	"strings"
)

type Set[T comparable] struct {
	data map[T]struct{}
}

// NewSet creates new Set[T] from slice
func NewSet[T comparable](data []T) Set[T] {
	var newSet Set[T]
	newSet.data = make(map[T]struct{})
	for _, elem := range data {
		newSet.Add(elem)
	}
	return newSet
}

func (set *Set[T]) Size() int {
	return len(set.data)
}

func (set *Set[T]) Add(val T) {
	set.data[val] = struct{}{}
}

func (set *Set[T]) Remove(val T) {
	delete(set.data, val)
}

func (set *Set[T]) Find(val T) bool {
	_, ok := set.data[val]
	return ok
}

func SetUnion[T comparable](l, r Set[T]) Set[T] {
	var unionSet Set[T] = NewSet(make([]T, 0))
	for k := range l.data {
		unionSet.Add(k)
	}
	for k := range r.data {
		unionSet.Add(k)
	}
	return unionSet
}

func SetIntersection[T comparable](l, r Set[T]) Set[T] {
	var setIntersection Set[T] = NewSet(make([]T, 0))
	if len(l.data) > len(r.data) {
		l, r = r, l
	}
	for k := range l.data {
		if r.Find(k) {
			setIntersection.Add(k)
		}
	}
	return setIntersection
}

func (set *Set[T]) String() string {
	var res strings.Builder
	for k := range set.data {
		res.WriteString(fmt.Sprintf("%v", k))
	}
	return res.String()
}