package core

import (
	"circulate/circulate/layouts"
	"syscall"

	"github.com/tadvi/winc/w32"
	jw32 "github.com/jcollie/w32"
	"golang.org/x/sys/windows"
)

var (
	user32      = windows.NewLazyDLL("user32.dll")
	enumWindows = user32.NewProc("EnumWindows")
)


func getWindows() []uintptr {
    container := []uintptr{}

	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		if !isElibible(h) {
			return 1
		}

		container = append(container, uintptr(h))
		return 1
	})

	_, _, _ = enumWindows.Call(cb, 0)
    return container
}


func isElibible(h syscall.Handle) bool {
	isWindowVisible := w32.IsWindowVisible(uintptr(h))
	isWindow := w32.IsWindow(uintptr(h))
	isWindowEnabled := w32.IsWindowEnabled(uintptr(h))
	windowText := w32.GetWindowText(uintptr(h))
	className, _ := jw32.GetClassName(jw32.HWND(h))

	if !isWindow ||
		!isWindowEnabled ||
		!isWindowVisible ||
		className == "Windows.UI.Core.CoreWindow" ||
		windowText == "" ||
		windowText == "Program Manager" ||
		windowText == "Calculator" ||
		windowText == "Add an account" ||
		windowText == "Settings" {

		return false
	}

	return true

}

func GetWindows() []uintptr  {
    return getWindows()

}

func SetWindows(windows []uintptr, rects []layouts.RECT)  {
    for i, hwnd := range windows {
        rect := rects[i]

		jw32.SetWindowPos(jw32.HWND(hwnd), jw32.HWND_NOTOPMOST, int(rect.Left), int(rect.Top), int(rect.Right ), int(rect.Bottom ), jw32.SWP_NOACTIVATE|0x0020)
    }


}
