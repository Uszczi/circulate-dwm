package core

import (
	"circulate/ty"
	"circulate/win"
	"fmt"

	"github.com/jcollie/w32"
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

	windowRect := w32.GetWindowRect(w32.HWND(hwnd))
	clientRect := w32.GetClientRect(w32.HWND(hwnd))

	var placement w32.WINDOWPLACEMENT
	w32.GetWindowPlacement(w32.HWND(hwnd), &placement)

	_frame, _ := w32.DwmGetWindowAttribute(w32.HWND(hwnd), w32.DWMWA_EXTENDED_FRAME_BOUNDS)
	frame, _ := _frame.(*w32.RECT)

	fmt.Printf("-------\n")
	fmt.Printf("%+v\n", hwnd)
	fmt.Printf("windowText: %+v\n", windowText)
	fmt.Printf("isWindowVisible: %+v\n", isWindowVisible)
	fmt.Printf("isWindow: %+v\n", isWindow)
	fmt.Printf("isWindowEnabled: %+v\n", isWindowEnabled)
	fmt.Printf("className: %+v\n", className)
	fmt.Printf("isWindowIconic: %+v\n", isWindowIconic)
	fmt.Printf("GWL_EXSTYLE 0x%x\n", gwlExStyle)
	fmt.Printf("\tWS_EX_TOPMOST %+v\n", win.WS_EX_TOPMOST&gwlExStyle)
	fmt.Printf("GWL_STYLE 0x%x\n", gwlStyle)
	fmt.Printf("\tWS_EX_TOOLWINDOW: %+v\n", win.WS_EX_TOOLWINDOW&gwlStyle)

	fmt.Printf("WS_EX_DLGMODALFRAME: %+v\n", win.WS_EX_DLGMODALFRAME&gwlStyle)
	fmt.Printf("WS_EX_NOPARENTNOTIFY: %+v\n", win.WS_EX_NOPARENTNOTIFY&gwlStyle)
	fmt.Printf("WS_EX_TOPMOST: %+v\n", win.WS_EX_TOPMOST&gwlStyle)
	fmt.Printf("WS_EX_ACCEPTFILES: %+v\n", win.WS_EX_ACCEPTFILES&gwlStyle)
	fmt.Printf("WS_EX_TRANSPARENT: %+v\n", win.WS_EX_TRANSPARENT&gwlStyle)
	fmt.Printf("WS_EX_MDICHILD: %+v\n", win.WS_EX_MDICHILD&gwlStyle)
	fmt.Printf("WS_EX_WINDOWEDGE: %+v\n", win.WS_EX_WINDOWEDGE&gwlStyle)
	fmt.Printf("WS_EX_CLIENTEDGE: %+v\n", win.WS_EX_CLIENTEDGE&gwlStyle)
	fmt.Printf("WS_EX_CONTEXTHELP: %+v\n", win.WS_EX_CONTEXTHELP&gwlStyle)
	fmt.Printf("WS_EX_RIGHT: %+v\n", win.WS_EX_RIGHT&gwlStyle)
	fmt.Printf("WS_EX_LEFT: %+v\n", win.WS_EX_LEFT&gwlStyle)
	fmt.Printf("WS_EX_RTLREADING: %+v\n", win.WS_EX_RTLREADING&gwlStyle)
	fmt.Printf("WS_EX_LTRREADING: %+v\n", win.WS_EX_LTRREADING&gwlStyle)
	fmt.Printf("WS_EX_LEFTSCROLLBAR: %+v\n", win.WS_EX_LEFTSCROLLBAR&gwlStyle)
	fmt.Printf("WS_EX_RIGHTSCROLLBAR: %+v\n", win.WS_EX_RIGHTSCROLLBAR&gwlStyle)
	fmt.Printf("WS_EX_CONTROLPARENT: %+v\n", win.WS_EX_CONTROLPARENT&gwlStyle)
	fmt.Printf("WS_EX_APPWINDOW: %+v\n", win.WS_EX_APPWINDOW&gwlStyle)
	fmt.Printf("WS_EX_OVERLAPPEDWINDOW: %+v\n", win.WS_EX_OVERLAPPEDWINDOW&gwlStyle)
	fmt.Printf("WS_EX_PALETTEWINDOW: %+v\n", win.WS_EX_PALETTEWINDOW&gwlStyle)
	fmt.Printf("WS_EX_LAYERED: %+v\n", win.WS_EX_LAYERED&gwlStyle)
	fmt.Printf("WS_EX_NOINHERITLAYOUT: %+v\n", win.WS_EX_NOINHERITLAYOUT&gwlStyle)
	fmt.Printf("WS_EX_LAYOUTRTL: %+v\n", win.WS_EX_LAYOUTRTL&gwlStyle)
	fmt.Printf("WS_EX_NOACTIVATE: %+v\n", win.WS_EX_NOACTIVATE&gwlStyle)

	fmt.Printf("rect %+v\n", windowRect)
	fmt.Printf("clientRect %+v\n", clientRect)
	fmt.Printf("placement.Length %+v\n", placement.Length)
	fmt.Printf("placement.Flags %+v\n", placement.Flags)
	fmt.Printf("placement.ShowCmd %+v\n", placement.ShowCmd)
	fmt.Printf("placement.PtMinPosition %+v\n", placement.PtMinPosition)
	fmt.Printf("placement.PtMaxPosition %+v\n", placement.PtMaxPosition)
	fmt.Printf("placement.RcNormalPosition %+v\n", placement.RcNormalPosition)
	fmt.Printf("frame %+v\n", frame)
	fmt.Printf("-------\n")
}
