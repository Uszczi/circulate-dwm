package core

import (
	"fmt"


	jw32 "github.com/jcollie/w32"
	"github.com/tadvi/winc/w32"
)

var (
	isIconic    = user32.NewProc("IsIconic")
)


func PrintDebugWindow(h uintptr) {
	isWindowIconic, _, _ := isIconic.Call(uintptr(h))
    isWindowsVisible := w32.IsWindowVisible(h)
	windowText := w32.GetWindowText(uintptr(h))
	windowRect := w32.GetWindowRect(uintptr(h))
	frame, _ := jw32.DwmGetWindowAttribute(jw32.HWND(h), jw32.DWMWA_EXTENDED_FRAME_BOUNDS)

	fmt.Printf("\nGetWindowText: %+v\n", windowText)
	fmt.Printf("isIconic: %+v\n", isWindowIconic)
	fmt.Printf("isWindowsVisible: %+v\n", isWindowsVisible)
	fmt.Printf("GetWindowRect: %+v\n", windowRect)
	fmt.Printf("DWMWA_EXTENDED_FRAME_BOUNDS: %+v\n\n", frame)

}
