package core

import (
	"circulate/circulate/store"
	"circulate/circulate/win"
	"fmt"

	"github.com/jcollie/w32"
)

func UseSetHowHide() {
	hwnd := store.GetSetHideShow()
	if hwnd == 0 {
		hwnd = win.GetForegroundWindow()
		store.SetSetHideShow(hwnd)
	}

	if win.IsWindowVisible(hwnd) && win.IsWindowIconic(hwnd) == 0 {
		fmt.Println("okno jest widoczne")
		w32.ShowWindow(w32.HWND(hwnd), 0)
	} else {
		fmt.Println("okno jest ne widoczne")
		w32.ShowWindow(w32.HWND(hwnd), 1)
	}

}
