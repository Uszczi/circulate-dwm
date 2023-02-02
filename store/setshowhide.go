package store

import (
	"circulate/ty"
	"sync"
)

type showHide struct {
	sync.RWMutex
	selected ty.HWND
}

var setShowHide = &showHide{}

func SetSetHideShow(hwnd ty.HWND) {
	setShowHide.Lock()
	setShowHide.selected = hwnd
	setShowHide.Unlock()
}
func GetSetHideShow() ty.HWND {
	setShowHide.RLock()
	defer setShowHide.RUnlock()
	return setShowHide.selected
}
