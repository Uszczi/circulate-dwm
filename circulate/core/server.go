package core

import (
	"fmt"
	"log"
	"syscall"

	"github.com/asaskevich/EventBus"
	jw32 "github.com/jcollie/w32"
	"github.com/tadvi/winc/w32"
	"golang.org/x/sys/windows"
)

var (
	modkernel32 = windows.NewLazyDLL("kernel32.dll")

	procSetWinEventHook = user32.NewProc("SetWinEventHook")
	procUnhookWinEvent  = user32.NewProc("UnhookWinEvent")

	procGetModuleHandle = modkernel32.NewProc("GetModuleHandleW")
)

type WINEVENTPROC func(hWinEventHook HWINEVENTHOOK, event uint32, hwnd HWND, idObject int32, idChild int32, idEventThread uint32, dwmsEventTime uint32) uintptr

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

const (
	//~ EVENT_SYSTEM_FOREGROUND DWORD = 0x0003
	//~ WINEVENT_OUTOFCONTEXT  DWORD = 0x0000
	//~ WINEVENT_INCONTEXT   = 0x0004
	EVENT_SYSTEM_FOREGROUND = 3
	WINEVENT_OUTOFCONTEXT   = 0
	WINEVENT_INCONTEXT      = 4
	WINEVENT_SKIPOWNPROCESS = 2
	WINEVENT_SKIPOWNTHREAD  = 1
)

func ActiveWinEventHook(hWinEventHook HWINEVENTHOOK, event uint32, hwnd HWND, idObject int32, idChild int32, idEventThread uint32, dwmsEventTime uint32) uintptr {
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
	log.Println("isWindow: ", jw32.IsWindow(jw32.HWND(hwnd)))
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

	return 0

}
func Handler(str string) {
	fmt.Println("handerrr", str)
}

func Main() {
	client := EventBus.NewClient(":2015", "/_client_bus_", EventBus.New())
	fmt.Println(client)
	client.Start()
	fmt.Println("started")
	client.Subscribe("main:calculator", Handler, ":2015", "/_client_bus_")
	fmt.Println("started")
	client.Stop()

	// hinst := w32.GetModuleHandle("")
	//
	// fmt.Println(hinst)
	//
	// winEvHook := SetWinEventHook(0x8000, 0x8000, 0, ActiveWinEventHook, 0, 0, WINEVENT_OUTOFCONTEXT|WINEVENT_SKIPOWNPROCESS)
	// defer w32.UnhookWindowsHookEx(w32.HANDLE(winEvHook))
	//
	// for {
	// 	var msg w32.MSG
	// 	if m := w32.GetMessage(&msg, 0, 0, 0); m != 0 {
	// 		w32.TranslateMessage(&msg)
	// 		w32.DispatchMessage(&msg)
	// 	}
	// }
}

func SetWinEventHook(eventMin DWORD, eventMax DWORD, hmodWinEventProc HMODULE, pfnWinEventProc WINEVENTPROC, idProcess DWORD, idThread DWORD, dwFlags DWORD) HWINEVENTHOOK {
	log.Println("procSetWinEventHook S")
	pfnWinEventProcCallback := syscall.NewCallback(pfnWinEventProc)
	ret, ret2, err := procSetWinEventHook.Call(
		uintptr(eventMin),
		uintptr(eventMax),
		uintptr(hmodWinEventProc),
		pfnWinEventProcCallback,
		uintptr(idProcess),
		uintptr(idThread),
		uintptr(dwFlags),
	)

	log.Printf("%#v", err)
	log.Printf("%#v", ret)
	log.Printf("%#v", ret2)
	log.Println("procSetWinEventHook E")
	return HWINEVENTHOOK(ret)
}

// mn =
