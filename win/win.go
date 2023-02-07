package win

import (
	"circulate/ty"
	"syscall"

	jw32 "github.com/jcollie/w32"
	_ "github.com/tadvi/winc/w32"
	"golang.org/x/sys/windows"
	_ "golang.org/x/sys/windows"
)

var (
	user32              = windows.NewLazyDLL("user32.dll")
	isIconic            = user32.NewProc("IsIconic")
	procSetWinEventHook = user32.NewProc("SetWinEventHook")
	enumWindows         = user32.NewProc("EnumWindows")
)

type WINEVENTPROC func(hWinEventHook HWINEVENTHOOK, event uint32, hwnd uintptr, idObject int32, idChild int32, idEventThread uint32, dwmsEventTime uint32) uintptr

type (
	HANDLE        uintptr
	HINSTANCE     HANDLE
	HHOOK         HANDLE
	HMODULE       HANDLE
	HWINEVENTHOOK HANDLE
	DWORD         uint32
	INT           int
	WPARAM        uintptr
	LPARAM        uintptr
	LRESULT       uintptr
	HWND          HANDLE
	UINT          uint32
	BOOL          int32
	ULONG_PTR     uintptr
	LONG          int32
	LPWSTR        *WCHAR
	WCHAR         uint16
)

func GetActiveWindow() ty.HWND {
	return ty.HWND(jw32.GetActiveWindow())
}

func GetForegroundWindow() ty.HWND {
	return ty.HWND(jw32.GetForegroundWindow())
}

func IsWindow(hwnd ty.HWND) bool {
	return jw32.IsWindow(jw32.HWND(hwnd))
}

func IsWindowVisible(hwnd ty.HWND) bool {
	return jw32.IsWindowVisible(jw32.HWND(hwnd))
}

func IsWindowEnabled(hwnd ty.HWND) bool {
	return jw32.IsWindowEnabled(jw32.HWND(hwnd))
}

func GetWindowText(hwnd ty.HWND) string {
	return jw32.GetWindowText(jw32.HWND(hwnd))
}

func GetClassName(hwnd ty.HWND) (string, bool) {
	return jw32.GetClassName(jw32.HWND(hwnd))
}

func IsWindowIconic(hwnd ty.HWND) uintptr {
	isIconic, _, _ := isIconic.Call(uintptr(hwnd))
	return isIconic
}

func GetWindowLongPtr(hwnd ty.HWND, index int) uintptr {
	return jw32.GetWindowLongPtr(jw32.HWND(hwnd), index)
}

func ShowWindow(hwnd ty.HWND, cmdshow int) bool {
	result := jw32.ShowWindow(jw32.HWND(hwnd), cmdshow)
	return result
}

func SetWinEventHook(eventMin DWORD, eventMax DWORD, hmodWinEventProc HMODULE, pfnWinEventProc WINEVENTPROC, idProcess DWORD, idThread DWORD, dwFlags DWORD) HWINEVENTHOOK {
	pfnWinEventProcCallback := syscall.NewCallback(pfnWinEventProc)
	ret, _, _ := procSetWinEventHook.Call(
		uintptr(eventMin),
		uintptr(eventMax),
		uintptr(hmodWinEventProc),
		pfnWinEventProcCallback,
		uintptr(idProcess),
		uintptr(idThread),
		uintptr(dwFlags),
	)

	return HWINEVENTHOOK(ret)
}

func EnumWindows(callback func(ty.HWND)) {
	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		callback(ty.HWND(h))
		return 0
	})
	_, _, _ = enumWindows.Call(cb, 0)
}
