package set

import (
	"sync"
)

type Set[T comparable] struct {
	mu    sync.RWMutex
	mp    map[T]struct{}
	total int
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{mp: make(map[T]struct{})}
}
func (set *Set[T]) Add(data ...T) {
	set.mu.Lock()
	defer set.mu.Unlock()
	set.total += len(data)
	for _, v := range data {
		set.mp[v] = struct{}{}
	}

}

func (set *Set[T]) List() []T {
	var list []T
	for i := range set.mp {
		list = append(list, i)
	}
	return list
}
func (set *Set[T]) Len() int {
	return set.total
}

func (set *Set[T]) DiffSlice(data ...T) []T {
	var diff []T
	for _, da := range data {
		if _, ok := set.mp[da]; !ok {
			diff = append(diff, da)
		}
	}

	return diff
}

func (set *Set[T]) GetNewitems(data []T) []T {
	var diff []T
	for _, da := range data {
		if _, ok := set.mp[da]; !ok {
			diff = append(diff, da)
		}
	}
	return diff
}
