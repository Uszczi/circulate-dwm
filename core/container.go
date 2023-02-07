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
	fmt.Printf("\nActiveWorkspace: %#v\n", cr.ActiveWorkspace+1)
	for _, workspace := range cr.Workspaces {
		fmt.Printf("%+v, %T\n", workspace, workspace.Layout)

		expected_setup := "[]"
		if len(workspace.WHWND) > 0 {
			rects := workspace.Layout.Calculate(workspace.WHWND)
			expected_setup = fmt.Sprint(rects)
		}

		fmt.Printf("Expected setup: %+v\n", expected_setup)
	}
}
