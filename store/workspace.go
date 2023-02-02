package store

import (
	"circulate/ty"
	"circulate/win"
	"fmt"
	"sync"
)

type workspacesStore struct {
	sync.RWMutex
	workspaces       []*workspaceStore
	active_workspace int
}

type workspaceStore struct {
	Windows []ty.HWND
	layout  string
}

var w = &workspacesStore{active_workspace: 1, workspaces: []*workspaceStore{
	{}, {}, {}, {}, {}, {},
}}

func PrintDebugWorkspace() {
	for _, workspace := range w.workspaces {
		fmt.Println(workspace.Windows)
	}
	return
}

func MoveToWorkspace(hwnd ty.HWND, workspaceName int) {
	w.workspaces[workspaceName].Windows = append(w.workspaces[workspaceName].Windows, hwnd)
}

func SwitchToLayout(workspaceName int) {
	w.active_workspace = workspaceName

	for _, workspace := range w.workspaces {
		for _, hwnd := range workspace.Windows {
			win.ShowWindow(hwnd, 6)
		}
	}
	for _, hwnd := range w.workspaces[workspaceName].Windows {
		win.ShowWindow(hwnd, 1)
	}
}

func GetActiveWorkspace() *workspaceStore {
	return w.workspaces[w.active_workspace]
}
