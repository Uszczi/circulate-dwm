package core

import (
	"circulate/ty"
	"circulate/win"
	"fmt"
	"syscall"

	"github.com/jcollie/w32"
)

func PrintDebugWindowNew(hwnd ty.HWND) {
	windowText := win.GetWindowText(hwnd)
	className, _ := win.GetClassName(hwnd)
	_, processId := w32.GetWindowThreadProcessId(w32.HWND(hwnd))
	ha := w32.OpenProcess(w32.PROCESS_QUERY_INFORMATION, false, uint32(processId))

	bufferSize := uint32(256)
	b := make([]uint16, bufferSize)
	win.QueryFullProcessImageName(syscall.Handle(ha), 0, &b[0], &bufferSize)

	exe := syscall.UTF16ToString(b)

	fmt.Printf("hwnd: %v, title:%v, exe: %v, className: %v\n", hwnd, windowText, exe, className)

}

func PrintDebugWindow(hwnd ty.HWND) {
	isWindowVisible := win.IsWindowVisible(hwnd)
	isWindow := win.IsWindow(hwnd)
	isWindowEnabled := win.IsWindowEnabled(hwnd)
	windowText := win.GetWindowText(hwnd)
	className, _ := win.GetClassName(hwnd)
	isWindowIconic := win.IsWindowIconic(hwnd)

	gwlEXStyle := win.GetWindowLongPtr(hwnd, win.GWL_EXSTYLE)
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
	fmt.Printf("gwlEXStyle |  0x%x\n", gwlEXStyle)
	fmt.Printf("gwlStyle | 0x%x\n", gwlStyle)

	inn := func(i bool) int {
		if i {
			return 1
		}
		return 0
	}

	fmt.Printf("WS_CHILD | %+v\n", inn(win.WS_CHILD&gwlStyle != 0))
	fmt.Printf("WS_POPUPWINDOW | %+v\n", inn(win.WS_POPUPWINDOW&gwlStyle != 0))
	fmt.Printf("WS_CHILDWINDOW | %+v\n", inn(win.WS_CHILDWINDOW&gwlStyle != 0))

	fmt.Printf("rect %+v\n", windowRect)
	fmt.Printf("clientRect %+v\n", clientRect)
	fmt.Printf("placement.Flags %+v\n", placement.Flags)
	fmt.Printf("placement.ShowCmd %+v\n", placement.ShowCmd)
	fmt.Printf("placement.PtMinPosition %+v\n", placement.PtMinPosition)
	fmt.Printf("placement.PtMaxPosition %+v\n", placement.PtMaxPosition)
	fmt.Printf("placement.RcNormalPosition %+v\n", placement.RcNormalPosition)
	fmt.Printf("frame %+v\n", frame)
	fmt.Printf("-------\n")
}
