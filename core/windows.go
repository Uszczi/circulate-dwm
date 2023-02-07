package core

import (
	"circulate/layouts"
	"circulate/ty"
	"circulate/win"

	jw32 "github.com/jcollie/w32"
)

var excludedClassNames = []string{"Windows.UI.Core.CoreWindow", "MSTaskListWClass"}

func getWindows() []ty.HWND {
	container := []ty.HWND{}

	callback := func(hwnd ty.HWND) {
		if !isElibible(hwnd) {
			return
		}

		container = append(container, hwnd)
		return
	}

	win.EnumWindows(callback)
	return container
}

func isElibible(hwnd ty.HWND) bool {
	isWindowVisible := win.IsWindowVisible(hwnd)
	isWindow := win.IsWindow(hwnd)
	isWindowEnabled := win.IsWindowEnabled(hwnd)
	windowText := win.GetWindowText(hwnd)
	className, _ := win.GetClassName(hwnd)
	isWindowIconic := win.IsWindowIconic(hwnd)
	gwlExStyle := win.GetWindowLongPtr(hwnd, win.GWL_EXSTYLE)

	if !isWindow ||
		!isWindowEnabled ||
		!isWindowVisible ||
		win.WS_EX_TOPMOST&gwlExStyle != 0 ||
		isWindowIconic == 1 ||
		className == "Windows.UI.Core.CoreWindow" ||
		windowText == "" ||
		windowText == "Program Manager" ||
		windowText == "Calculator" ||
		windowText == "Add an account" ||
		windowText == "Settings" {

		return false
	}
	for _, name := range excludedClassNames {
		if name == className {
			return false
		}
	}

	PrintDebugWindow(hwnd)
	return true

}

func GetWindows() []ty.HWND {
	return getWindows()

}

func SetWindows(windows []ty.HWND, rects []layouts.RECT) {
	for i, hwnd := range windows {
		rect := rects[i]

		jw32.SetWindowPos(jw32.HWND(hwnd), jw32.HWND_NOTOPMOST, int(rect.Left), int(rect.Top), int(rect.Right), int(rect.Bottom), jw32.SWP_NOACTIVATE|0x0020)
	}

}
