package core

import (
	"circulate/ty"
	"circulate/win"
	"fmt"
	"log"
	"syscall"

	jw32 "github.com/jcollie/w32"
	"github.com/tadvi/winc/w32"
)

const (
	EVENT_SYSTEM_FOREGROUND = 3
	WINEVENT_OUTOFCONTEXT   = 0
	WINEVENT_INCONTEXT      = 4
	WINEVENT_SKIPOWNPROCESS = 2
	WINEVENT_SKIPOWNTHREAD  = 1
)

func ActiveWinEventHook(hWinEventHook win.HWINEVENTHOOK, event uint32, hwnd uintptr, idObject int32, idChild int32, idEventThread uint32, dwmsEventTime uint32) uintptr {
	if !isElibible(syscall.Handle(hwnd)) {
		return 0
	}

	log.Print("\n\n\n")
	log.Println("event:", event)
	log.Println("hWinEventHook:", hWinEventHook)
	log.Println("hwnd:", hwnd)
	log.Println("idObject:", idObject)
	log.Println("idChild:", idChild)
	log.Println("idEventThread:", idEventThread)
	log.Println("dwmsEventTime:", dwmsEventTime)

	isWindow := w32.IsWindow(uintptr(hwnd))
	log.Println("isWindow: ", isWindow)
	log.Println("title: ", w32.GetWindowText(uintptr(hwnd)))
	log.Println("IsWindowEnabled: ", w32.IsWindowEnabled(uintptr(hwnd)))

	a, b := w32.GetWindowThreadProcessId(uintptr(hwnd))
	log.Println("GetWindowThreadProcessId:", a, b)
	var windowInfo w32.WINDOWINFO
	w32.GetWindowInfo(uintptr(hwnd), &windowInfo)
	fmt.Printf("GetWindowInfo: %+v\n", windowInfo)
	var windowPlacement w32.WINDOWPLACEMENT
	w32.GetWindowPlacement(uintptr(hwnd), &windowPlacement)
	fmt.Printf("GetWindowPlacement: %+v\n", windowPlacement)
	topWindow := jw32.GetTopWindow(jw32.HWND(hwnd))

	fmt.Printf("GetTopWindow: %#+v\n", topWindow)

	GWL_EXSTYLE := w32.GetWindowLong(uintptr(hwnd), w32.GWL_EXSTYLE)
	fmt.Printf("GWL_EXSTYLE: %#v\n", GWL_EXSTYLE)

	GWL_STYLE := w32.GetWindowLong(uintptr(hwnd), w32.GWL_STYLE)
	fmt.Printf("GWL_STYLE: %#v\n", GWL_STYLE)

	GW_OWNER := jw32.GetWindow(jw32.HWND(hwnd), jw32.GW_OWNER)
	fmt.Printf("GW_OWNER: %#v\n", GW_OWNER)

	fmt.Printf("IsWindowVisible: %#v\n", jw32.IsWindowVisible(jw32.HWND(hwnd)))

	windows := GetWindows()
	var ww []ty.HWND

	for _, v := range windows {
		ww = append(ww, ty.HWND(v))
	}

	// positions := layouts.CalculateColumns(ww)
	// SetWindows(ww, positions)

	return 0

}
func Handler(str string) {
	fmt.Println("handerrr", str)
}

func RunWindowsServer() {
	fmt.Println("Start Windows Server")
	w32.GetModuleHandle("") // [TODO] check what it is

	// [TODO] what is 0x8000
	// [TODO] check last argument
	winEvHook := win.SetWinEventHook(0x8000, 0x8000, 0, ActiveWinEventHook, 0, 0, WINEVENT_OUTOFCONTEXT|WINEVENT_SKIPOWNPROCESS)
	defer w32.UnhookWindowsHookEx(w32.HANDLE(winEvHook))

	for {
		var msg w32.MSG
		if m := w32.GetMessage(&msg, 0, 0, 0); m != 0 {
			w32.TranslateMessage(&msg)
			w32.DispatchMessage(&msg)
		}
	}
}
