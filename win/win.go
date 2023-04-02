package win

import (
	"circulate/ty"
	"syscall"
	"unsafe"

	jw32 "github.com/jcollie/w32"
	"github.com/tadvi/winc/w32"
	"golang.org/x/sys/windows"
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

type tWNDCLASSEXW struct {
	size       uint32
	style      uint32
	wndProc    uintptr
	clsExtra   int32
	wndExtra   int32
	instance   syscall.Handle
	icon       syscall.Handle
	cursor     syscall.Handle
	background syscall.Handle
	menuName   *uint16
	className  *uint16
	iconSm     syscall.Handle
}

var (
	user32                     = windows.NewLazyDLL("user32.dll")
	AllowSetForegroundWindow   = user32.NewProc("AllowSetForegroundWindow") // Move this to func
	AttachThreadInput          = user32.NewProc("AttachThreadInput")        // Move this to func
	enumWindows                = user32.NewProc("EnumWindows")
	isIconic                   = user32.NewProc("IsIconic")
	procSetWinEventHook        = user32.NewProc("SetWinEventHook")
	createWindowExW            = user32.NewProc("CreateWindowExW")
	defWindowProcW             = user32.NewProc("DefWindowProcW")
	destroyWindow              = user32.NewProc("DestroyWindow")
	dispatcheW                 = user32.NewProc("GetMessageW")
	LoadCursorW                = user32.NewProc("LoadCursorW")
	postQuitMessage            = user32.NewProc("PostQuitMessage")
	registerClassExW           = user32.NewProc("RegisterClassExW")
	setLayeredWindowAttributes = user32.NewProc("SetLayeredWindowAttributes")

	modgdi32         = syscall.NewLazyDLL("gdi32.dll")
	createSolidBrush = modgdi32.NewProc("CreateSolidBrush")
	createPen        = modgdi32.NewProc("CreatePen")

	modkernel32                = syscall.NewLazyDLL("kernel32.dll")
	queryFullProcessImageNameW = modkernel32.NewProc("QueryFullProcessImageNameW")
)

func CreateWindow(eXStyle uint32, className, windowName string, style uint32, x, y, width, height int64, parent, menu, instance syscall.Handle) error {
	ret, _, err := createWindowExW.Call(
		uintptr(eXStyle),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(className))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(windowName))),
		uintptr(style),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(parent),
		uintptr(menu),
		uintptr(instance),
		uintptr(0),
	)
	if ret == 0 {
		return err
	}
	return nil
}

func DefWindowProc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) uintptr {
	// TODO what it does?
	ret, _, _ := defWindowProcW.Call(
		uintptr(hwnd),
		uintptr(msg),
		uintptr(wparam),
		uintptr(lparam),
	)
	return uintptr(ret)
}

func DestroyWindow(hwnd syscall.Handle) error {
	ret, _, err := destroyWindow.Call(uintptr(hwnd))
	if ret == 0 {
		return err
	}
	return nil
}

func PostQuitMessage(exitCode int32) {
	postQuitMessage.Call(uintptr(exitCode))
}

func RegisterClassEx(wcx *tWNDCLASSEXW) (uint16, error) {
	ret, _, err := registerClassExW.Call(
		uintptr(unsafe.Pointer(wcx)),
	)
	if ret == 0 {
		return 0, err
	}
	return uint16(ret), nil
}

func CreateSolidBrush(color uint32) w32.HBRUSH {
	brush, _, _ := createSolidBrush.Call(uintptr(color))
	return w32.HBRUSH(brush)
}

func CreatePen(penType uint, width uint, color uintptr) uintptr {
	pen, _, _ := createPen.Call(0, uintptr(width), uintptr(color))
	return pen
}

func SetLayeredWindowAttributes(hwnd ty.HWND, color uint, notsure uint, notsure2 uint) (uintptr, error) {
	res, _, err := setLayeredWindowAttributes.Call(uintptr(hwnd), uintptr(color), uintptr(notsure), uintptr(notsure2))
	return res, err
}

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

func GetSystemMetrics(index int) int {
	result := jw32.GetSystemMetrics(index)
	return result
}

func GetModuleHandle(modulename string) HMODULE {
	result := jw32.GetModuleHandle(modulename)
	return HMODULE(result)
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
		return 1
	})
	_, _, _ = enumWindows.Call(cb, 0)
}

func QueryFullProcessImageName(process syscall.Handle, flags uint32, exeName *uint16, size *uint32) error {
	r1, _, e1 := syscall.Syscall6(
		queryFullProcessImageNameW.Addr(),
		4,
		uintptr(process),
		uintptr(flags),
		uintptr(unsafe.Pointer(exeName)),
		uintptr(unsafe.Pointer(size)),
		0,
		0)
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}
