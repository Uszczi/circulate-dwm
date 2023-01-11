package main

import (
	"fmt"
	"syscall"

	jw32 "github.com/jcollie/w32"
	"github.com/tadvi/winc/w32"
	"golang.org/x/sys/windows"
)

var (
	container = []syscall.Handle{}
)

var (
	user32      = windows.NewLazyDLL("user32.dll")
	enumWindows = user32.NewProc("EnumWindows")
	isIconic    = user32.NewProc("IsIconic")
)

func start() {
	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		var dddd w32.WINDOWINFO
		var windowPlacement w32.WINDOWPLACEMENT
		isWindowVisible := w32.IsWindowVisible(uintptr(h))
		w32.GetWindowInfo(uintptr(h), &dddd)
		w32.GetWindowPlacement(uintptr(h), &windowPlacement)
		isWindow := w32.IsWindow(uintptr(h))
		isWindowEnabled := w32.IsWindowEnabled(uintptr(h))

		name, _ := jw32.GetClassName(jw32.HWND(h))

		windowTitle := w32.GetWindowText(uintptr(h))
		if isWindow && isWindowEnabled && isWindowVisible {
			if name != "Windows.UI.Core....CoreWindow" && windowTitle != "Program Manager" {
				fmt.Println("")
				fmt.Println("")
				fmt.Println(w32.GetWindowText(uintptr(h)))
				fmt.Println(isIconic.Call(uintptr(h)))
				fmt.Printf("%+v", windowPlacement)

			}
		}

		return 1
	})

	enumWindows.Call(cb, 0)
}

func main() {
	start()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Found '%s' window: handle=0x%x\n", title, h)
}

// [] What is getclient
// [] What is ismanageble
// [] What is manage

// BOOL CALLBACK scan(HWND hwnd, LPARAM lParam) {
//   Client *c = getclient(hwnd);
//   if (c)
//     c->isalive = true;
//   else if (ismanageable(hwnd))
//     manage(hwnd);
//
