package core

import (
	"circulate/layouts"
	"circulate/ty"
	"circulate/win"

	jw32 "github.com/jcollie/w32"
)

type Workspace struct {
	WHWND  []ty.HWND
	Layout layouts.Layout
}

func (ws *Workspace) RemoveWindow(hwnd ty.HWND) bool {
	for k, wHWND := range ws.WHWND {
		if wHWND == hwnd {
			ws.WHWND = append(ws.WHWND[:k], ws.WHWND[k+1:]...)
			return true
		}
	}
	return false
}

func (ws *Workspace) AddWindow(hwnd ty.HWND) bool {
	for _, wHWND := range ws.WHWND {
		if wHWND == hwnd {
			return false
		}
	}

	ws.WHWND = append(ws.WHWND, hwnd)
	return true
}

func (ws *Workspace) UpdateLayout() {
	if len(ws.WHWND) == 0 {
		return
	}

	rects := ws.Layout.Calculate(ws.WHWND)
	for i, hwnd := range ws.WHWND {
		rect := rects[i]
		jw32.SetWindowPos(jw32.HWND(hwnd), jw32.HWND_NOTOPMOST, int(rect.Left), int(rect.Top), int(rect.Right), int(rect.Bottom), jw32.SWP_NOACTIVATE|0x0020)
	}
}

func (ws *Workspace) ShowWorkspace() {
	for _, wHWND := range ws.WHWND {
		win.ShowWindow(wHWND, 1)
	}
}

func (ws *Workspace) HideWorkspace() {
	for _, wHWND := range ws.WHWND {
		win.ShowWindow(wHWND, 6)
	}
}
