package win

import (
	"circulate/ty"
	"log"
	"syscall"
	"unsafe"

	"github.com/jcollie/w32"
)

const toolbar_height_px = 40

func GetDesktopWidth() int {
	monitor_width := GetSystemMetrics(SM_CXSCREEN)
	return monitor_width
}

func GetDesktopHeight() int {
	monitor_height := GetSystemMetrics(SM_CYSCREEN)
	return monitor_height - toolbar_height_px
}

func GetDesktopDimentions() (int, int) {
	return GetDesktopWidth(), GetDesktopHeight()
}

func GetWindowByClassName(name string) ty.HWND {
	// TODO find a better way for this
	var res ty.HWND
	callback := func(hwnd ty.HWND) {
		if GetWindowText(hwnd) == name {
			res = hwnd
		}
	}
	EnumWindows(callback)
	return res
}

// TODO this cannot be there
func callback(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) uintptr {
	// TODO check docs about this
	switch msg {
	case WM_DESTROY:
		DestroyWindow(hwnd)
	case WM_CLOSE:
		PostQuitMessage(0)
	default:
		ret := DefWindowProc(hwnd, msg, wparam, lparam)
		return ret
	}
	return 0
}

func CreateBorderWindow(name string) ty.HWND {
	moduleHandle := GetModuleHandle("")
	brush := CreateSolidBrush(0x000000)

	wcx := w32.WNDCLASSEX{
		WndProc:    syscall.NewCallback(callback),
		Instance:   w32.HINSTANCE(moduleHandle),
		Cursor:     0,
		Background: w32.HBRUSH(brush),
		ClassName:  syscall.StringToUTF16Ptr(name),
	}
	wcx.Size = uint32(unsafe.Sizeof(wcx))

	w32.RegisterClassEx(&wcx)

	CreateWindow(
		WS_EX_TOOLWINDOW|WS_EX_LAYERED,
		name,
		name,
		WS_POPUP|WS_SYSMENU|WS_MAXIMIZEBOX|WS_MINIMIZEBOX,
		0,
		0,
		0,
		0,
		0,
		0,
		syscall.Handle(moduleHandle),
	)
	hwnd := GetWindowByClassName(name)

	// TODO try to understand SetLayeredWindowAttributes
	ok, err := SetLayeredWindowAttributes(hwnd, 0x000000, 0x255, 1)
	if ok == 0 {
		log.Printf("SetLayeredWindowAttributes failed: [%v]", err)
	}

	return hwnd
}
