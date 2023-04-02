package core

import (
	"circulate/ty"
	"circulate/win"
	"fmt"
	"log"

	"github.com/tadvi/winc/w32"
)

type BorderWindow ty.HWND

var borderWidth = 3

func (bw *BorderWindow) SetPosition(rect ty.RECT) {
	log.Printf("BorderWindow.SetPosition rect=%+v\n", rect)
	w32.SetWindowPos(w32.HWND(*bw), w32.HWND_TOPMOST, rect.Left, rect.Top, rect.Right, rect.Bottom, w32.SWP_SHOWWINDOW|w32.SWP_NOACTIVATE)
}

func (bw *BorderWindow) Hide() {
	win.ShowWindow(ty.HWND(*bw), 0)
}

func SetBorderWindow(hwnd ty.HWND) {
	log.Printf("BorderWindow.SetWindow border: %v hwnd: %+v\n", Borderwindow, hwnd)
	PrintDebugWindowNew(hwnd)
	if Borderwindow == nil {
		CreateBorderWindow()
	}

	rect := w32.GetWindowRect(uintptr(hwnd))
	if win.GetWindowText(hwnd) == "" || rect.Top < -100 {
		log.Printf("[SetBorderWindow] hidding")
		Borderwindow.Hide()
		return
	}

	penWidth := 5
	rec2 := ty.RECT{Left: int(rect.Left), Top: int(rect.Top) - penWidth, Right: int(rect.Right - rect.Left), Bottom: int(rect.Bottom-rect.Top) + borderWidth}
	Borderwindow.SetPosition(rec2)

	var paint w32.PAINTSTRUCT
	w32.BeginPaint(uintptr(*Borderwindow), &paint)

	pen := win.CreatePen(0, uint(10), 0x00ff00)
	brush := win.CreateSolidBrush(0)

	w32.SelectObject(paint.Hdc, pen)
	w32.SelectObject(paint.Hdc, brush)
	w32.Rectangle(paint.Hdc, 0, 0, int32(rec2.Right), int32(rec2.Bottom))
	w32.EndPaint(uintptr(*Borderwindow), &paint)

	fmt.Println(rect)
}

var (
	className  = "CirculateBorder"
	windowText = className
)
var Borderwindow *BorderWindow

func CreateBorderWindow() {
	border := BorderWindow(win.CreateBorderWindow(className))
	Borderwindow = &border

	fmt.Println("Window border hwnd ", border)

	for {
		// TODO how does it work?
		var msg w32.MSG
		if m := w32.GetMessage(&msg, 0, 0, 0); m != 0 {
			w32.TranslateMessage(&msg)
			w32.DispatchMessage(&msg)
		}
	}
}
