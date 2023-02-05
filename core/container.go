package core

import (
	"circulate/ty"
	"fmt"
)

type Container struct {
	ActiveWorkspace int
	Workspaces      []*Workspace
}

func (cr *Container) MoveToWorkspace(hwnd ty.HWND, workspaceName int) bool {
	for _, workspace := range cr.Workspaces {
		workspace.RemoveWindow(hwnd)
	}
	return cr.Workspaces[workspaceName].AddWindow(hwnd)
}

func (cr *Container) SwitchToWorkspace(workspaceID int) {
	cr.ActiveWorkspace = workspaceID

	for _, workspace := range cr.Workspaces {
		workspace.HideWorkspace()
	}
	cr.Workspaces[cr.ActiveWorkspace].ShowWorkspace()
}

func (cr *Container) PrintDebugWorkspace() {
	for _, workspace := range cr.Workspaces {
		fmt.Println(workspace.WHWND)
	}
	return
}
