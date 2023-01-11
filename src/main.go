package main

import (
	"fmt"
	"syscall"

	jw32 "github.com/jcollie/w32"
	"github.com/tadvi/winc/w32"
	"golang.org/x/sys/windows"
)

var (
	user32      = windows.NewLazyDLL("user32.dll")
	isIconic    = user32.NewProc("IsIconic")
	enumWindows = user32.NewProc("EnumWindows")
)

var (
	container = []syscall.Handle{}
)

func printDebugWindow(h syscall.Handle) {
	isWindowIconic, _, _ := isIconic.Call(uintptr(h))
	windowText := w32.GetWindowText(uintptr(h))

	fmt.Print("\n\n")
	fmt.Printf("GetWindowText: %+v\n", windowText)
	fmt.Printf("isIconic: %+v\n", isWindowIconic)

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
		windowText == "" ||
		className == "Windows.UI.Core.CoreWindow" ||
		windowText == "Program Manager" {
		return false
	}
	return true

}

func start() {
	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		if !isElibible(h) {
			return 1
		}

		printDebugWindow(h)
		container = append(container, h)

		return 1
	})

	enumWindows.Call(cb, 0)

	fmt.Printf("container: %+v\n", container)

}

func main() {
	start()
}
