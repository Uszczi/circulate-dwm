package main

import (
	"fmt"
	"syscall"
	"unsafe"

	ww "github.com/jcollie/w32"
	"github.com/tadvi/winc/w32"
	"golang.org/x/sys/windows"
)

var (
	user32             = syscall.MustLoadDLL("user32.dll")
	procEnumWindows    = user32.MustFindProc("EnumWindows")
	procGetWindowTextW = user32.MustFindProc("GetWindowTextW")

	winuserDLL    = windows.NewLazyDLL("user32.dll")
	enumWindows   = winuserDLL.NewProc("EnumWindows")
	getWindowInfo = winuserDLL.NewProc("GetWindowInfo")
	isIconic      = winuserDLL.NewProc("IsIconic")
	// getww         = user32.MustFindProc("GetModuleFileNameA")
	// isWindow              = winuserDLL.NewProc("IsWindow")
	// isWindowVisible       = winuserDLL.NewProc("IsWindowVisible")
	// isIconic           = winuserDLL.NewProc("IsIconic")
	// isMinimized        = winuserDLL.NewProc("isMinimized")

	dwmapi                = windows.NewLazyDLL("dwmapi.dll")
	dwmGetWindowAttribute = dwmapi.NewProc("DwmGetWindowAttribute")
)

func GetWindowText(hwnd syscall.Handle, str *uint16, maxCount int32) (len int32, err error) {
	r0, _, e1 := syscall.Syscall(procGetWindowTextW.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(str)), uintptr(maxCount))
	len = int32(r0)
	if len == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

//     let containers = unsafe { &mut *(lparam.0 as *mut VecDeque<Container>) };
//
//     let is_visible = WindowsApi::is_window_visible(hwnd);

//     let is_window = WindowsApi::is_window(hwnd);
//     let is_minimized = WindowsApi::is_iconic(hwnd);
//
//     if is_visible && is_window && !is_minimized {
//         let window = Window { hwnd: hwnd.0 };
//
//         if let Ok(should_manage) = window.should_manage(None) {
//             if should_manage {
//                 let mut container = Container::default();
//                 container.windows_mut().push_back(window);
//                 containers.push_back(container);
//             }
//         }
//     }
//
//     true.into()
// }

func FindWindow(title string) (syscall.Handle, error) {
	var hwnd syscall.Handle
	// var windows w32.WINDOWPLACEMENT

	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		// var ass *uint16
		// var ass2 *uint16
		var dddd w32.WINDOWINFO
		var windowPlacement w32.WINDOWPLACEMENT
		// var windowRect w32.RECT

		// a, dd, e := dwmGetWindowAttribute.Call(uintptr(h), windows.DWMWA_CLOAKED, uintptr(unsafe.Pointer(ass)), 16)
		// a, _, e = dwmGetWindowAttribute.Call(uintptr(h), windows.DWMWA_CLOAK, uintptr(unsafe.Pointer(ass2)), 2)
		isWindowVisible := w32.IsWindowVisible(uintptr(h))

		w32.GetWindowInfo(uintptr(h), &dddd)
		// _, processId := w32.GetWindowThreadProcessId(uintptr(h))
		w32.GetWindowPlacement(uintptr(h), &windowPlacement)
		// windowRect = *w32.GetWindowRect(uintptr(h))
		isWindow := w32.IsWindow(uintptr(h))
		isWindowEnabled := w32.IsWindowEnabled(uintptr(h))
		// fmt.Println(getww.Call(uintptr(processId)))

		name, _ := ww.GetClassName(ww.HWND(h))

		// DwmGetWindowAttribute(hwnd, DWMWA_EXTENDED_FRAME_BOUNDS, &frame, sizeof(RECT));
		// i1, _, _ := isMinimized.Call(uintptr(h))

		b := make([]uint16, 200)
		_, err := GetWindowText(h, &b[0], int32(len(b)))

		if err != nil {
			// ignore the error
			return 1 // continue enumeration
		}
		windowTitle := w32.GetWindowText(uintptr(h))
		if isWindow && syscall.UTF16ToString(b) != "" && isWindowEnabled && isWindowVisible {
			if name != "Windows.UI.Core....CoreWindow" && windowTitle != "Program Manager" {
				fmt.Println("")
				fmt.Println("")
				// // fmt.Println("")
				// fmt.Println(syscall.UTF16ToString(b))
				fmt.Println(w32.GetWindowText(uintptr(h)))
				fmt.Println(isIconic.Call(uintptr(h)))
				fmt.Printf("%+v", windowPlacement)
				//
				// fmt.Printf("%+v\n", a)
				// fmt.Printf("%+v\n", dd)
				// fmt.Printf("%+v\n", e)
				// fmt.Println(processId)
				// fmt.Println(name, name2)
				// aew := w32.OpenProcess(uint32(processId), false, uint32(processId))
				// fmt.Println(aew)
				// fmt.Println("hello", ass2)
				// fmt.Println("hello", ass2)
				// fmt.Println(dddd.DwStyle)
				// fmt.Println(dddd)
				// fmt.Printf("%+v\n", dddd)
				// fmt.Printf("%+v", windowRect)
				// fmt.Printf("%+v", isWindowEnabled)
				aee := w32.WINDOWPLACEMENT{}

				if "Nowa karta - Google Chrome" == syscall.UTF16ToString(b) {
					w32.SetWindowPos(uintptr(h), 0, 100, 100, 500, 500, 0)
					w32.SetWindowPlacement(uintptr(h), &aee)
					fmt.Println("USTASWIQAMAAM")
					fmt.Println(w32.ShowWindow(uintptr(h), 1), "herer")
					// w32.SetForegroundWindow(uintptr(h))
					// fmt.Println(w32.SetActiveWindow(uintptr(h)), "herer")
					// fmt.Println(w32.GetLastError())

				}
				// w32.ShowWindow(uintptr(h), 1)

				// if "Battery Meter" == syscall.UTF16ToString(b) {
				// 	fmt.Println("batteryyyyyyyyyyy --------------")
				// 	w32.SetWindowPos(uintptr(h), 0, 100, 100, 500, 500, 0)
				// 	// w32.SetActiveWindow(uintptr(h))
				// 	// w32.SetForegroundWindow(uintptr(h))
				// 	fmt.Printf("%+v\n", windowPlacement)
				// 	fmt.Printf("%+v\n", windowRect)
				// 	fmt.Printf("%+v\n", dddd)
				// w32.ShowWindow(uintptr(h), 3)

				// }

			}
		}

		return 1 // continue enumeration
	})

	enumWindows.Call(cb, 0)
	if hwnd == 0 {
		return 0, nil
	}
	return hwnd, nil
}

func main() {
	const title = "Game"
	FindWindow(title)
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
