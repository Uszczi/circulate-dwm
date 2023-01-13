package core

import (
	"fmt"
	"syscall"

	"circulate/circulate/layouts"

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
	windowRect := w32.GetWindowRect(uintptr(h))
	frame, _ := jw32.DwmGetWindowAttribute(jw32.HWND(h), jw32.DWMWA_EXTENDED_FRAME_BOUNDS)

	fmt.Printf("GetWindowText: %+v\n", windowText)
	fmt.Printf("isIconic: %+v\n", isWindowIconic)
	fmt.Printf("GetWindowRect: %+v\n", windowRect)
	fmt.Printf("DWMWA_EXTENDED_FRAME_BOUNDS: %+v\n", frame)

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
		windowText == "Program Manager" ||
		windowText == "Calculator" ||
		windowText == "Add an account" ||
		windowText == "Settings" {

		return false
	}
	return true

}

func setupContainer() {
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

func calculatePositions() []w32.RECT {
	amount := len(container)
	result := []w32.RECT{}

	monitor_width := w32.GetSystemMetrics(0)
	monitor_height := w32.GetSystemMetrics(1) - 37

	if amount == 1 {
		return append(result, w32.RECT{Left: 0, Top: 0, Right: int32(monitor_width), Bottom: int32(monitor_height)})
	}

	width := monitor_width / amount
	fmt.Println("Widhttt", width)
	left := 0
	top := 0
	_ = top
	right := left + width
	fmt.Println("right", right)

	bottom := monitor_height
	_ = bottom

	for _, h := range container {
		fmt.Println(left, right)
		frame, _ := jw32.DwmGetWindowAttribute(jw32.HWND(h), jw32.DWMWA_EXTENDED_FRAME_BOUNDS)
		windowRect := jw32.GetWindowRect(jw32.HWND(h))

		frame2, ok := frame.(*jw32.RECT)
		if ok {
			w_left := int(windowRect.Left) - int(frame2.Left) + left
			w_top := int(windowRect.Top) - int(frame2.Top) + top
			w_right := 2*(int(windowRect.Right)-int(frame2.Right)) + width
			w_bottom := int(windowRect.Bottom) - int(frame2.Bottom) + bottom

			fmt.Println("")
			fmt.Println(frame)
			fmt.Println(windowRect)
			fmt.Println(w_left, w_top, w_right, w_bottom)
			fmt.Println(left, top, width, bottom)
			fmt.Println("")

			result = append(result, w32.RECT{Left: int32(w_left), Top: int32(w_top), Right: int32(w_right), Bottom: int32(w_bottom)})
			left += width
			right += width

		}

	}

	return result
}

func Start() {
	setupContainer()
	// positions := calculatePositions()
	positions := layouts.CalculateRows(container)
	_ = positions
	fmt.Println(positions)

	for i, h := range container {
		rect := positions[i]
		fmt.Printf("\n\n")
		printDebugWindow(h)
		fmt.Println(rect)
		// fmt.Println("GWL_EXSTYLE", w32.GetWindowLong(uintptr(h), w32.GWL_EXSTYLE))
		// fmt.Println("GWL_STYLE", w32.GetWindowLong(uintptr(h), w32.GWL_STYLE))

		// https://learn.microsoft.com/en-us/windows/win32/winmsg/window-styles
		// w32.SetWindowLong(uintptr(h), w32.GWL_EXSTYLE, 256)
		// w32.SetWindowLong(uintptr(h), w32.GWL, 256)
		// w32.SetWindowLong(uintptr(h), w32.GWL_EXSTYLE, w32.WS_BORDER)
		//
		jw32.SetWindowPos(jw32.HWND(h), jw32.HWND_NOTOPMOST, int(rect.Left), int(rect.Top), int(rect.Right), int(rect.Bottom), jw32.SWP_NOACTIVATE|0x0020)
		fmt.Println(jw32.GetLastError())
		// w32.MoveWindow(uintptr(h), int(rect.Left), int(rect.Top), int(rect.Right), int(rect.Bottom), true)
		// jw32.SetWindowPos
		// w32.SetWindowPos
		// fmt.Println(positions[i])
	}

	fmt.Printf("container: %+v\n", container)
}
