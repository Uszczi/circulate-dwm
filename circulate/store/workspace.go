package store

import (
	"circulate/circulate/ty"
	"circulate/circulate/win"
	"fmt"
	"sync"
)

type workspacesStore struct {
	sync.RWMutex
	workspaces       []*workspaceStore
	active_workspace int
}

type workspaceStore struct {
	w []ty.HWND
}

var workspaces = &workspacesStore{active_workspace: 0, workspaces: []*workspaceStore{
	&workspaceStore{}, &workspaceStore{}, &workspaceStore{}, &workspaceStore{}, &workspaceStore{}, &workspaceStore{}}}

func PrintDebugWorkspace() {
	for _, workspace := range workspaces.workspaces {
		fmt.Println(workspace.w)
	}
	return
}
func MoveToWorkspace(hwnd ty.HWND, workspaceName int) {
	fmt.Println("Move to worksapce storee")
	workspaces.workspaces[workspaceName].w = append(workspaces.workspaces[workspaceName].w, hwnd)
}

func SwitchToLayout(workspaceName int) {
	for _, workspace := range workspaces.workspaces {
		for _, hwnd := range workspace.w {
			win.ShowWindow(hwnd, 6)
		}
	}
	for _, hwnd := range workspaces.workspaces[workspaceName].w {
		win.ShowWindow(hwnd, 1)
	}
}
