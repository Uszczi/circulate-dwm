package store

import "sync"

type showHide struct {
	sync.RWMutex
	selected uintptr
}

var setShowHide = &showHide{}

func SetSetHideShow(hwnd uintptr) {
	setShowHide.Lock()
	setShowHide.selected = hwnd
	setShowHide.Unlock()
}
func GetSetHideShow() uintptr {
	setShowHide.RLock()
	defer setShowHide.RUnlock()
	return setShowHide.selected
}
