package set

import "sync"

type UintSet struct {
	mu sync.RWMutex
	mp map[uint]struct{}
}

func NewUintSet() *UintSet {
	return &UintSet{mp: map[uint]struct{}{}}
}
func (set *UintSet) Add(data uint) {
	set.mu.Lock()
	set.mp[data] = struct{}{}
	set.mu.Unlock()
}

func (set *UintSet) List() []uint {
	var list []uint
	for i, _ := range set.mp {
		list = append(list, i)
	}
	return list
}

func (set *UintSet)DiffSlice(data []uint)[]uint {
	var diff []uint
	for _, da := range data {
		if _,ok:=set.mp[da];!ok {
			diff = append(diff, da)
		}
	}
	return diff
}

func (set *UintSet)GetAddItems(data []int64)[]uint {
	var diff []uint
	for _, da := range data {
		if _,ok:=set.mp[uint(da)];!ok {
			diff = append(diff, uint(da))
		}
	}
	return diff
}