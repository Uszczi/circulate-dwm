package core

import (
	"circulate/ty"
	"circulate/win"
	"fmt"

	"github.com/tadvi/winc/w32"
)

const (
	EVENT_SYSTEM_FOREGROUND = 3
	WINEVENT_OUTOFCONTEXT   = 0
	WINEVENT_INCONTEXT      = 4
	WINEVENT_SKIPOWNPROCESS = 2
	WINEVENT_SKIPOWNTHREAD  = 1
	EVENT_OBJECT_CREATE     = 0x8000
	EVENT_OBJECT_DESTROY    = 0x8001
	EVENT_OBJECT_FOCUS      = 0x8005
	EVENT_OBJECT_HIDE       = 0x8003
	EVENT_OBJECT_SHOW       = 0x8002
)

func ActiveWinEventHook(hWinEventHook win.HWINEVENTHOOK, event uint32, hwnd uintptr, idObject int32, idChild int32, idEventThread uint32, dwmsEventTime uint32) uintptr {
	if !isElibible(ty.HWND(hwnd)) {
		return 0
	}
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
	winEvHook := win.SetWinEventHook(EVENT_OBJECT_CREATE, EVENT_OBJECT_CREATE, 0, ActiveWinEventHook, 0, 0, 0)
	defer w32.UnhookWindowsHookEx(w32.HANDLE(winEvHook))

	for {
		var msg w32.MSG
		if m := w32.GetMessage(&msg, 0, 0, 0); m != 0 {
			w32.TranslateMessage(&msg)
			w32.DispatchMessage(&msg)
		}
	}
}
