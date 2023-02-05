package core

import (
	"circulate/layouts"
	"circulate/ty"
	"circulate/win"
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
	return
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
