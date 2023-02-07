package core

import (
	"circulate/layouts"
	"circulate/ty"
	"circulate/win"
	"log"

	jw32 "github.com/jcollie/w32"
	"github.com/tadvi/winc/w32"
)

func getWindows() []uintptr {
	container := []uintptr{}

	callback := func(hwnd ty.HWND) {
		if !isElibible(hwnd) {
			return
		}

		container = append(container, uintptr(hwnd))
		return
	}

	win.EnumWindows(callback)
	return container
}

func isElibible(h ty.HWND) bool {

	isWindowVisible := w32.IsWindowVisible(uintptr(h))
	isWindow := w32.IsWindow(uintptr(h))
	isWindowEnabled := w32.IsWindowEnabled(uintptr(h))
	windowText := w32.GetWindowText(uintptr(h))
	className, _ := jw32.GetClassName(jw32.HWND(h))
	isWindowIconic := win.IsWindowIconic(h)
	if !isWindow ||
		!isWindowEnabled ||
		!isWindowVisible ||
		isWindowIconic == 1 ||
		className == "Windows.UI.Core.CoreWindow" ||
		windowText == "" ||
		windowText == "Program Manager" ||
		windowText == "Calculator" ||
		windowText == "Add an account" ||
		windowText == "Settings" {

		return false
	}

	log.Printf("\n\n")
	log.Println("windowText: ", windowText)
	log.Println("HWND: ", h)
	log.Println("isWindowVisible: ", isWindowVisible)
	log.Println("isWindow: ", isWindow)
	log.Println("isWindowEnabled: ", isWindowEnabled)
	log.Println("className: ", className)
	log.Println("isWindowIconic: ", isWindowIconic)
	log.Println("getActivewindows: ", jw32.GetActiveWindow())
	log.Println("GetForegroundWindow: ", jw32.GetForegroundWindow())

	return true

}

func GetWindows() []uintptr {
	return getWindows()

}

func SetWindows(windows []ty.HWND, rects []layouts.RECT) {
	for i, hwnd := range windows {
		rect := rects[i]

		jw32.SetWindowPos(jw32.HWND(hwnd), jw32.HWND_NOTOPMOST, int(rect.Left), int(rect.Top), int(rect.Right), int(rect.Bottom), jw32.SWP_NOACTIVATE|0x0020)
	}

}
