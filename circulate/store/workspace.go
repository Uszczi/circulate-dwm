package store

import (
	"circulate/circulate/ty"
	"circulate/circulate/win"
	"fmt"
	"sync"
)

type workspace struct {
	sync.RWMutex
	windows []ty.HWND
}

var workspaces = []*workspace{&workspace{}, &workspace{}}

func PrintDebugWorkspace(hwnd ty.HWND) {
	for _, workspace := range workspaces {
		fmt.Println(workspace.windows)
	}
	return
}
func MoveToWorkspace(index int) {
	foregroundWindow := win.GetForegroundWindow()

	workspaces[index].windows = append(workspaces[index].windows, foregroundWindow)
}
