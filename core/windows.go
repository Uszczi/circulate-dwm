package core

import (
	"circulate/ty"
	"circulate/win"
)

var excludedClassNames = []string{
	"Windows.UI.Core.CoreWindow",
	"MSTaskListWClass",
	"Windows.UI.Composition.DesktopWindowContentBridge",
}

func getWindows() []ty.HWND {
	container := []ty.HWND{}

	callback := func(hwnd ty.HWND) {
		if !IsElibible(hwnd) {
			return
		}

		container = append(container, hwnd)
		return
	}

	win.EnumWindows(callback)
	return container
}

func IsElibible(hwnd ty.HWND) bool {
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
