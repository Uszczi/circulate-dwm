package core

import (
	"circulate/ty"
	"circulate/win"
	"fmt"
)

func PrintDebugWindow(hwnd ty.HWND) {
	isWindowVisible := win.IsWindowVisible(hwnd)
	isWindow := win.IsWindow(hwnd)
	isWindowEnabled := win.IsWindowEnabled(hwnd)
	windowText := win.GetWindowText(hwnd)
	className, _ := win.GetClassName(hwnd)
	isWindowIconic := win.IsWindowIconic(hwnd)

	gwlExStyle := win.GetWindowLongPtr(hwnd, win.GWL_EXSTYLE)
	gwlStyle := win.GetWindowLongPtr(hwnd, win.GWL_STYLE)

	fmt.Printf("-------\n")
	fmt.Printf("%+v\n", hwnd)
	fmt.Printf("windowText: %+v\n", windowText)
	fmt.Printf("isWindowVisible: %+v\n", isWindowVisible)
	fmt.Printf("isWindow: %+v\n", isWindow)
	fmt.Printf("isWindowEnabled: %+v\n", isWindowEnabled)
	fmt.Printf("className: %+v\n", className)
	fmt.Printf("isWindowIconic: %+v\n", isWindowIconic)
	fmt.Printf("GWL_EXSTYLE 0x%x\n", gwlExStyle)
	fmt.Printf("GWL_STYLE 0x%x\n", gwlStyle)
	fmt.Printf("WS_EX_TOPMOST %+v\n", win.WS_EX_TOPMOST&gwlExStyle)
	fmt.Printf("-------\n")

}
