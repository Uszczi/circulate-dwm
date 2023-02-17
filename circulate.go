package main

import (
	"circulate/cmd"
	"circulate/ty"
	"circulate/win"
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"github.com/tadvi/winc/w32"
)

var kernel32 = syscall.NewLazyDLL("kernel32.dll")

const TRANSPARENCY_COLOUR = 0

// Flags for SetLayeredWindowAttributes.
const (
	LWA_COLORKEY = 0x1
	LWA_ALPHA    = 0x2
)

var (
	user32 = syscall.NewLazyDLL("user32.dll")

	pCreateWindowExW  = user32.NewProc("CreateWindowExW")
	pDefWindowProcW   = user32.NewProc("DefWindowProcW")
	pDestroyWindow    = user32.NewProc("DestroyWindow")
	pDispatchMessageW = user32.NewProc("DispatchMessageW")
	pGetMessageW      = user32.NewProc("GetMessageW")
	pLoadCursorW      = user32.NewProc("LoadCursorW")
	pPostQuitMessage  = user32.NewProc("PostQuitMessage")
	pRegisterClassExW = user32.NewProc("RegisterClassExW")
	pTranslateMessage = user32.NewProc("TranslateMessage")
)

const (
	cSW_SHOW        = 5
	cSW_USE_DEFAULT = 0x80000000
)

const (
	cWS_MAXIMIZE_BOX = 0x00010000
	cWS_MINIMIZEBOX  = 0x00020000
	cWS_THICKFRAME   = 0x00040000
	cWS_SYSMENU      = 0x00080000
	cWS_CAPTION      = 0x00C00000
	cWS_VISIBLE      = 0x10000000

	cWS_OVERLAPPEDWINDOW = 0x00CF0000
)

func createWindow(className, windowName string, style uint32, eXStyle uint32, x, y, width, height int64, parent, menu, instance syscall.Handle) (syscall.Handle, error) {
	ret, _, err := pCreateWindowExW.Call(
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
		return 0, err
	}
	return syscall.Handle(ret), nil
}

const (
	cWM_DESTROY = 0x0002
	cWM_CLOSE   = 0x0010
)

func defWindowProc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) uintptr {
	ret, _, _ := pDefWindowProcW.Call(
		uintptr(hwnd),
		uintptr(msg),
		uintptr(wparam),
		uintptr(lparam),
	)
	return uintptr(ret)
}

func destroyWindow(hwnd syscall.Handle) error {
	ret, _, err := pDestroyWindow.Call(uintptr(hwnd))
	if ret == 0 {
		return err
	}
	return nil
}

type tPOINT struct {
	x, y int32
}

type tMSG struct {
	hwnd    syscall.Handle
	message uint32
	wParam  uintptr
	lParam  uintptr
	time    uint32
	pt      tPOINT
}

func dispatchMessage(msg *tMSG) {
	pDispatchMessageW.Call(uintptr(unsafe.Pointer(msg)))
}

func getMessage(msg *tMSG, hwnd syscall.Handle, msgFilterMin, msgFilterMax uint32) (bool, error) {
	ret, _, err := pGetMessageW.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(hwnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax),
	)
	if int32(ret) == -1 {
		return false, err
	}
	return int32(ret) != 0, nil
}

func postQuitMessage(exitCode int32) {
	pPostQuitMessage.Call(uintptr(exitCode))
}

const (
	cCOLOR_WINDOW = 5
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

func registerClassEx(wcx *tWNDCLASSEXW) (uint16, error) {
	ret, _, err := pRegisterClassExW.Call(
		uintptr(unsafe.Pointer(wcx)),
	)
	if ret == 0 {
		return 0, err
	}
	return uint16(ret), nil
}

func translateMessage(msg *tMSG) {
	pTranslateMessage.Call(uintptr(unsafe.Pointer(msg)))
}

func main() {
	if len(os.Args) > 1 {
		cmd.Execute()
		return
	}
	className := "CirculateBorder"
	windowText := "CirculateBorder"
	moduleHandle := win.GetModuleHandle("")

	modgdi32 := syscall.NewLazyDLL("gdi32.dll")
	user32 := syscall.NewLazyDLL("user32.dll")
	procCreateSolidBrush := modgdi32.NewProc("CreateSolidBrush")
	procSetLayeredWindowAttributes := user32.NewProc("SetLayeredWindowAttributes")

	brush, _, _ := procCreateSolidBrush.Call(TRANSPARENCY_COLOUR)
	fmt.Println(brush)

	fn := func(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) uintptr {
		switch msg {
		case cWM_CLOSE:
			destroyWindow(hwnd)
		case cWM_DESTROY:
			postQuitMessage(0)
		default:
			ret := defWindowProc(hwnd, msg, wparam, lparam)
			return ret
		}
		return 0
	}

	wcx := w32.WNDCLASSEX{
		WndProc:    syscall.NewCallback(fn),
		Instance:   uintptr(moduleHandle),
		Cursor:     0,
		Background: brush,
		ClassName:  syscall.StringToUTF16Ptr(className),
	}

	wcx.Size = uint32(unsafe.Sizeof(wcx))
	w32.RegisterClassEx(&wcx)

	createWindow(
		className,
		windowText,
		w32.WS_POPUP|w32.WS_SYSMENU|w32.WS_MAXIMIZEBOX|w32.WS_MINIMIZEBOX,
		w32.WS_EX_TOOLWINDOW|w32.WS_EX_LAYERED,
		100,
		100,
		300,
		300,
		0,
		0,
		syscall.Handle(moduleHandle),
	)

	var wn ty.HWND
	callback := func(hwnd ty.HWND) {
		if win.GetWindowText(hwnd) == "CirculateBorder" {
			wn = hwnd
		}
	}
	win.EnumWindows(callback)

	procSetLayeredWindowAttributes.Call(uintptr(wn), 0xffffff, 0x255, 1)

	w32.SetWindowPos(w32.HWND(wn), w32.HWND_NOTOPMOST, 100, 100, 100, 100, w32.SWP_SHOWWINDOW|w32.SWP_NOACTIVATE)

	for {
		msg := tMSG{}
		getMessage(&msg, 0, 0, 0)
		translateMessage(&msg)
		dispatchMessage(&msg)
	}
}
