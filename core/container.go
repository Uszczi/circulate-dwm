package core

import (
	"circulate/ty"
	"circulate/win"
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
	return cr.Workspaces[workspaceName-1].AddWindow(hwnd)
}

func (cr *Container) SwitchToWorkspace(workspaceID int) {
	cr.ActiveWorkspace = workspaceID

	for _, workspace := range cr.Workspaces {
		workspace.HideWorkspace()
	}
	cr.Workspaces[cr.ActiveWorkspace-1].ShowWorkspace()
}

func (cr *Container) PrintDebugWorkspace() {
	fmt.Printf("\nActiveWorkspace: %#v\n", cr.ActiveWorkspace)
	for _, workspace := range cr.Workspaces {
		fmt.Printf("%+v, %T\n", workspace, workspace.Layout)

		for _, hwnd := range workspace.WHWND {
			fmt.Printf("%v: %+v\n", hwnd, win.GetWindowText(hwnd))
		}

		expected_setup := "[]"
		if len(workspace.WHWND) > 0 {
			rects := workspace.Layout.Calculate(workspace.WHWND)
			expected_setup = fmt.Sprint(rects)
		}

		fmt.Printf("Expected setup: %+v\n\n", expected_setup)
	}
}
