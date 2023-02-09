package core

import (
	"circulate/ty"
	"circulate/win"
	"fmt"

	"github.com/jcollie/w32"
)

func PrintDebugWindow(hwnd ty.HWND) {
	return
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

	fmt.Printf("WS_OVERLAPPED | %+v\n", inn(win.WS_OVERLAPPED&gwlStyle != 0))
	fmt.Printf("WS_POPUP | %+v\n", inn(win.WS_POPUP&gwlStyle != 0))
	fmt.Printf("WS_CHILD | %+v\n", inn(win.WS_CHILD&gwlStyle != 0))
	fmt.Printf("WS_MINIMIZE | %+v\n", inn(win.WS_MINIMIZE&gwlStyle != 0))
	fmt.Printf("WS_VISIBLE | %+v\n", inn(win.WS_VISIBLE&gwlStyle != 0))
	fmt.Printf("WS_DISABLED | %+v\n", inn(win.WS_DISABLED&gwlStyle != 0))
	fmt.Printf("WS_CLIPSIBLINGS | %+v\n", inn(win.WS_CLIPSIBLINGS&gwlStyle != 0))
	fmt.Printf("WS_CLIPCHILDREN | %+v\n", inn(win.WS_CLIPCHILDREN&gwlStyle != 0))
	fmt.Printf("WS_MAXIMIZE | %+v\n", inn(win.WS_MAXIMIZE&gwlStyle != 0))
	fmt.Printf("WS_CAPTION | %+v\n", inn(win.WS_CAPTION&gwlStyle != 0))
	fmt.Printf("WS_BORDER | %+v\n", inn(win.WS_BORDER&gwlStyle != 0))
	fmt.Printf("WS_DLGFRAME | %+v\n", inn(win.WS_DLGFRAME&gwlStyle != 0))
	fmt.Printf("WS_VSCROLL | %+v\n", inn(win.WS_VSCROLL&gwlStyle != 0))
	fmt.Printf("WS_HSCROLL | %+v\n", inn(win.WS_HSCROLL&gwlStyle != 0))
	fmt.Printf("WS_SYSMENU | %+v\n", inn(win.WS_SYSMENU&gwlStyle != 0))
	fmt.Printf("WS_THICKFRAME | %+v\n", inn(win.WS_THICKFRAME&gwlStyle != 0))
	fmt.Printf("WS_GROUP | %+v\n", inn(win.WS_GROUP&gwlStyle != 0))
	fmt.Printf("WS_TABSTOP | %+v\n", inn(win.WS_TABSTOP&gwlStyle != 0))
	fmt.Printf("WS_MINIMIZEBOX | %+v\n", inn(win.WS_MINIMIZEBOX&gwlStyle != 0))
	fmt.Printf("WS_MAXIMIZEBOX | %+v\n", inn(win.WS_MAXIMIZEBOX&gwlStyle != 0))
	fmt.Printf("WS_TILED | %+v\n", inn(win.WS_TILED&gwlStyle != 0))
	fmt.Printf("WS_ICONIC | %+v\n", inn(win.WS_ICONIC&gwlStyle != 0))
	fmt.Printf("WS_SIZEBOX | %+v\n", inn(win.WS_SIZEBOX&gwlStyle != 0))
	fmt.Printf("WS_OVERLAPPEDWINDOW | %+v\n", inn(win.WS_OVERLAPPEDWINDOW&gwlStyle != 0))
	fmt.Printf("WS_POPUPWINDOW | %+v\n", inn(win.WS_POPUPWINDOW&gwlStyle != 0))
	fmt.Printf("WS_CHILDWINDOW | %+v\n", inn(win.WS_CHILDWINDOW&gwlStyle != 0))

	fmt.Printf("WS_EX_TOOLWINDOW |  %+v\n", inn(win.WS_EX_TOOLWINDOW&gwlEXStyle != 0))
	fmt.Printf("WS_EX_TOPMOST %+v\n", inn(win.WS_EX_TOPMOST&gwlEXStyle != 0))
	fmt.Printf("WS_EX_DLGMODALFRAME |  %+v\n", inn(win.WS_EX_DLGMODALFRAME&gwlEXStyle != 0))
	fmt.Printf("WS_EX_NOPARENTNOTIFY |  %+v\n", inn(win.WS_EX_NOPARENTNOTIFY&gwlEXStyle != 0))
	fmt.Printf("WS_EX_TOPMOST |  %+v\n", inn(win.WS_EX_TOPMOST&gwlEXStyle != 0))
	fmt.Printf("WS_EX_ACCEPTFILES |  %+v\n", inn(win.WS_EX_ACCEPTFILES&gwlEXStyle != 0))
	fmt.Printf("WS_EX_TRANSPARENT |  %+v\n", inn(win.WS_EX_TRANSPARENT&gwlEXStyle != 0))
	fmt.Printf("WS_EX_MDICHILD |  %+v\n", inn(win.WS_EX_MDICHILD&gwlEXStyle != 0))
	fmt.Printf("WS_EX_WINDOWEDGE |  %+v\n", inn(win.WS_EX_WINDOWEDGE&gwlEXStyle != 0))
	fmt.Printf("WS_EX_CLIENTEDGE |  %+v\n", inn(win.WS_EX_CLIENTEDGE&gwlEXStyle != 0))
	fmt.Printf("WS_EX_CONTEXTHELP |  %+v\n", inn(win.WS_EX_CONTEXTHELP&gwlEXStyle != 0))
	fmt.Printf("WS_EX_RIGHT |  %+v\n", inn(win.WS_EX_RIGHT&gwlEXStyle != 0))
	fmt.Printf("WS_EX_LEFT |  %+v\n", inn(win.WS_EX_LEFT&gwlEXStyle != 0))
	fmt.Printf("WS_EX_RTLREADING |  %+v\n", inn(win.WS_EX_RTLREADING&gwlEXStyle != 0))
	fmt.Printf("WS_EX_LTRREADING |  %+v\n", inn(win.WS_EX_LTRREADING&gwlEXStyle != 0))
	fmt.Printf("WS_EX_LEFTSCROLLBAR |  %+v\n", inn(win.WS_EX_LEFTSCROLLBAR&gwlEXStyle != 0))
	fmt.Printf("WS_EX_RIGHTSCROLLBAR |  %+v\n", inn(win.WS_EX_RIGHTSCROLLBAR&gwlEXStyle != 0))
	fmt.Printf("WS_EX_CONTROLPARENT |  %+v\n", inn(win.WS_EX_CONTROLPARENT&gwlEXStyle != 0))
	fmt.Printf("WS_EX_APPWINDOW |  %+v\n", inn(win.WS_EX_APPWINDOW&gwlEXStyle != 0))
	fmt.Printf("WS_EX_OVERLAPPEDWINDOW |  %+v\n", inn(win.WS_EX_OVERLAPPEDWINDOW&gwlEXStyle != 0))
	fmt.Printf("WS_EX_PALETTEWINDOW |  %+v\n", inn(win.WS_EX_PALETTEWINDOW&gwlEXStyle != 0))
	fmt.Printf("WS_EX_LAYERED |  %+v\n", inn(win.WS_EX_LAYERED&gwlEXStyle != 0))
	fmt.Printf("WS_EX_NOINHERITLAYOUT |  %+v\n", inn(win.WS_EX_NOINHERITLAYOUT&gwlEXStyle != 0))
	fmt.Printf("WS_EX_LAYOUTRTL |  %+v\n", inn(win.WS_EX_LAYOUTRTL&gwlEXStyle != 0))
	fmt.Printf("WS_EX_NOACTIVATE |  %+v\n", inn(win.WS_EX_NOACTIVATE&gwlEXStyle != 0))

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
