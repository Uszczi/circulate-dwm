package core

// import (
// 	"fmt"
// 	"syscall"
//
// 	"circulate/circulate/layouts"
//
// 	jw32 "github.com/jcollie/w32"
// 	"github.com/tadvi/winc/w32"
// 	"golang.org/x/sys/windows"
// )
//
// var (
// 	user32      = windows.NewLazyDLL("user32.dll")
// 	isIconic    = user32.NewProc("IsIconic")
// 	enumWindows = user32.NewProc("EnumWindows")
// )
//
//
//
//
//
// func Start() {
// 	setupContainer()
// 	// positions := calculatePositions()
// 	positions := layouts.CalculateRows(container)
// 	_ = positions
// 	fmt.Println(positions)
//
// 	for i, h := range container {
// 		rect := positions[i]
// 		fmt.Printf("\n\n")
// 		printDebugWindow(h)
// 		fmt.Println(rect)
// 		// fmt.Println("GWL_EXSTYLE", w32.GetWindowLong(uintptr(h), w32.GWL_EXSTYLE))
// 		// fmt.Println("GWL_STYLE", w32.GetWindowLong(uintptr(h), w32.GWL_STYLE))
//
// 		// https://learn.microsoft.com/en-us/windows/win32/winmsg/window-styles
// 		// w32.SetWindowLong(uintptr(h), w32.GWL_EXSTYLE, 256)
// 		// w32.SetWindowLong(uintptr(h), w32.GWL, 256)
// 		// w32.SetWindowLong(uintptr(h), w32.GWL_EXSTYLE, w32.WS_BORDER)
// 		//
// 		jw32.SetWindowPos(jw32.HWND(h), jw32.HWND_NOTOPMOST, int(rect.Left), int(rect.Top), int(rect.Right), int(rect.Bottom), jw32.SWP_NOACTIVATE|0x0020)
// 		fmt.Println(jw32.GetLastError())
// 		// w32.MoveWindow(uintptr(h), int(rect.Left), int(rect.Top), int(rect.Right), int(rect.Bottom), true)
// 		// jw32.SetWindowPos
// 		// w32.SetWindowPos
// 		// fmt.Println(positions[i])
// 	}
//
// 	fmt.Printf("container: %+v\n", container)
// }
