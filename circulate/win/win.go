package win

import (
	"circulate/circulate/ty"
	jw32 "github.com/jcollie/w32"
	_ "github.com/tadvi/winc/w32"
	"golang.org/x/sys/windows"
	_ "golang.org/x/sys/windows"
)

var (
	user32   = windows.NewLazyDLL("user32.dll")
	isIconic = user32.NewProc("IsIconic")
)

func GetActiveWindow() ty.HWND {
	return ty.HWND(jw32.GetActiveWindow())
}

func GetForegroundWindow() ty.HWND {
	return ty.HWND(jw32.GetForegroundWindow())
}

func IsWindowVisible(hwnd ty.HWND) bool {
	return jw32.IsWindowVisible(jw32.HWND(hwnd))
}

func IsWindowIconic(hwnd ty.HWND) uintptr {
	isIconic, _, _ := isIconic.Call(uintptr(hwnd))
	return isIconic
}
