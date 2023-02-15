package core

import (
	"circulate/ty"
	"circulate/win"
)

var excludedClassNames = []string{
	"MSTaskListWClass",
	"Windows.UI.Composition.DesktopWindowContentBridge",
	"Windows.UI.Core.CoreWindow",
	"ToolbarWindow32",
	"SysListView32",
	"SysTreeView32",
}

func IsElibible(hwnd ty.HWND) bool {
	isWindowVisible := win.IsWindowVisible(hwnd)
	isWindow := win.IsWindow(hwnd)
	isWindowEnabled := win.IsWindowEnabled(hwnd)
	windowText := win.GetWindowText(hwnd)
	className, _ := win.GetClassName(hwnd)
	isWindowIconic := win.IsWindowIconic(hwnd)
	gwlExStyle := win.GetWindowLongPtr(hwnd, win.GWL_EXSTYLE)
	// gwlStyle := win.GetWindowLongPtr(hwnd, win.GWL_STYLE)

	if handleUnclearWindows(hwnd, className, windowText) {
		return true
	}

	if !isWindow ||
		!isWindowEnabled ||
		!isWindowVisible ||
		isWindowIconic == 1 ||
		win.WS_EX_TOPMOST&gwlExStyle != 0 ||
		win.WS_EX_TOOLWINDOW&gwlExStyle != 0 ||
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

func handleUnclearWindows(hwnd ty.HWND, className string, windowText string) bool {
	if className == "XLMAIN" && windowText != "" {
		return true
	}

	if className == "CabinetWClass" {
		return true
	}

	return false
}

func GetWindows() []ty.HWND {
	return getWindows()
}
