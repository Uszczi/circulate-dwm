package store

import "sync"

var f = &foo2{}

func SetActiveLayout(name string) {
	f.Lock()
	f.ActiveLayout = name
	f.Unlock()
}
func GetActiveLayout() string {
	f.RLock()
	defer f.RUnlock()
	return f.ActiveLayout
}

func Incr() {
	f.incr()
}

func Count() int {
	return f.count()
}

type foo2 struct {
	sync.RWMutex
	count2       int
	ActiveLayout string
}

func (f *foo2) incr() {
	f.Lock()
	f.count2++
	f.Unlock()
}

func (f *foo2) count() int {
	f.RLock()
	defer f.RUnlock()
	return f.count2
}
