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
	W      []ty.HWND
	layout string
}

var workspaces = &workspacesStore{active_workspace: 0, workspaces: []*workspaceStore{
	&workspaceStore{}, &workspaceStore{}, &workspaceStore{}, &workspaceStore{}, &workspaceStore{}, &workspaceStore{}}}

func PrintDebugWorkspace() {
	for _, workspace := range workspaces.workspaces {
		fmt.Println(workspace.W)
	}
	return
}
func MoveToWorkspace(hwnd ty.HWND, workspaceName int) {
	fmt.Println("Move to worksapce storee")
	workspaces.workspaces[workspaceName].W = append(workspaces.workspaces[workspaceName].W, hwnd)
}

func SwitchToLayout(workspaceName int) {
	for _, workspace := range workspaces.workspaces {
		for _, hwnd := range workspace.W {
			win.ShowWindow(hwnd, 6)
		}
	}
	for _, hwnd := range workspaces.workspaces[workspaceName].W {
		win.ShowWindow(hwnd, 1)
	}
}

func GetActiveWorkspace() *workspaceStore {
	return workspaces.workspaces[workspaces.active_workspace]
}
