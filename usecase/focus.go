package usecase

import (
	"circulate/store"
	"circulate/ty"
	"circulate/win"
	"log"
	"syscall"
	"time"

	"github.com/jcollie/w32"
	"golang.org/x/sys/windows"
)

var (
	libuser32         *windows.LazyDLL
	attachThreadInput *windows.LazyProc
)

func init() {
	libuser32 = windows.NewLazySystemDLL("user32.dll")
	attachThreadInput = libuser32.NewProc("AttachThreadInput")
}

func FocusNext() {
	log.Println("[usecase.FocusNext]")
	activeWindow := win.GetForegroundWindow()
	log.Println("[usecase.FocusNext] activeWindow:", activeWindow)

	activeWindowIndex := 0
	hwnds := store.GetActiveWorkspace().WHWND
	log.Println(hwnds)
	for i, hwnd := range hwnds {
		if hwnd == activeWindow {
			activeWindowIndex = i
		}
	}

	nextWindowHwnd := hwnds[(activeWindowIndex+1)%len(hwnds)]
	log.Println("nextWindowHwnd", nextWindowHwnd)
	// ok := win.ShowWindow(nextWindowHwnd, win.SW_SHOW)

	// ok := w32.SetForegroundWindow(w32.HWND(nextWindowHwnd))
	// log.Println("[usecase.FocusNext] SetForegroundWindow result:", ok)
	forceForegroundWindow(int(nextWindowHwnd))
	// w32.MouseInput(w32.MOUSEINPUT{Dx: -1000, Dy: 500})
	// w32.MOUSEEVENTF_MOVE
}

func forceForegroundWindow(HWND int) {
	// w32.GetFocus()
	windowThreadProcessId, _ := w32.GetWindowThreadProcessId(w32.GetForegroundWindow())
	currentThreadId, _ := w32.GetWindowThreadProcessId(w32.HWND(HWND))
	// currentThreadId := w32.GetCurrentThread()
	// w32.SetFocus(w32.HWND(HWND))
	log.Println(windowThreadProcessId, currentThreadId)

	// CONST_SW_SHOW := 5

	w1, w2, w3 := AttachThreadInput(windowThreadProcessId, currentThreadId, 1)
	log.Println("attach to input", w1, w2, w3)

	r, r2, r3 := win.AllowSetForegroundWindow.Call(0xFFFFFFFF)
	log.Println(r, r2, r3)

	ok := w32.SetForegroundWindow(w32.HWND(HWND))
	log.Println("[usecase.FocusNext] SetForegroundWindow result:", ok)
	// win.ShowWindow(ty.HWND(HWND), CONST_SW_SHOW)
	time.Sleep(200 * time.Microsecond)

	// time.Sleep(1 * time.Second)
	e1, e2, e3 := AttachThreadInput(windowThreadProcessId, currentThreadId, 0)
	log.Println("detach from input", e1, e2, e3)
	erro := w32.GetLastError()
	log.Println(erro)
}

func AttachThreadInput(idAttach w32.HANDLE, idAttachTo w32.HANDLE, fAttach int) (uintptr, uintptr, error) {
	ret, r2, r3 := syscall.Syscall(attachThreadInput.Addr(), 3,
		uintptr(idAttach),
		uintptr(idAttachTo),
		uintptr(fAttach))

	return ret, r2, r3
}

func FocusPrevious() {
	log.Println("[usecase.FocusPrevious]")
	activeWindow := win.GetForegroundWindow()
	log.Println("[usecase.FocusPrevious] activeWindow:", activeWindow)

	activeWindowIndex := 0
	hwnds := store.GetActiveWorkspace().WHWND
	log.Println(hwnds)
	for i, hwnd := range hwnds {
		if hwnd == activeWindow {
			activeWindowIndex = i
		}
	}

	var nextWindowHwnd ty.HWND
	if activeWindowIndex == 0 {
		nextWindowHwnd = hwnds[len(hwnds)-1]
	} else {
		nextWindowHwnd = hwnds[(activeWindowIndex-1)%len(hwnds)]
	}

	log.Println("nextWindowHwnd", nextWindowHwnd)
	// ok := win.ShowWindow(nextWindowHwnd, win.SW_SHOW)

	// ok := w32.SetForegroundWindow(w32.HWND(nextWindowHwnd))
	// log.Println("[usecase.FocusNext] SetForegroundWindow result:", ok)
	r, r2, r3 := win.AllowSetForegroundWindow.Call(0xFFFFFFFF)
	log.Println(r, r2, r3)
	forceForegroundWindow(int(nextWindowHwnd))
	// w32.MouseInput(w32.MOUSEINPUT{Dx: -1000, Dy: 500})
	// w32.MOUSEEVENTF_MOVE
}
