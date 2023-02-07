package core

import (
	"circulate/ty"
	"circulate/win"
	"fmt"

	jw32 "github.com/jcollie/w32"
	"github.com/tadvi/winc/w32"
)

func PrintDebugWindow(h uintptr) {
	isWindowIconic := win.IsWindowIconic(ty.HWND(h))
	isWindowsVisible := w32.IsWindowVisible(h)
	windowText := w32.GetWindowText(uintptr(h))
	windowRect := w32.GetWindowRect(uintptr(h))
	frame, _ := jw32.DwmGetWindowAttribute(jw32.HWND(h), jw32.DWMWA_EXTENDED_FRAME_BOUNDS)

	fmt.Printf("\nGetWindowText: %+v\n", windowText)
	fmt.Printf("isIconic: %+v\n", isWindowIconic)
	fmt.Printf("IsWindowVisible: %+v\n", isWindowsVisible)
	fmt.Printf("GetWindowRect: %+v\n", windowRect)
	fmt.Printf("DWMWA_EXTENDED_FRAME_BOUNDS: %+v\n\n", frame)

}
