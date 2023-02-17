package usecase

import (
	"circulate/core"
	"circulate/ty"
	"circulate/win"
	"log"
)

func PrintWorkspaceDebug() {
	log.Println("[usecase.PrintWorkspaceDebug]")
	core.GetContainer().PrintDebugWorkspace()
}

func ClearWorkspace() {
	log.Println("[usecase.ClearWorkspace]")
	core.GetActiveWorkspace().WHWND = []ty.HWND{}
}

func MoveToWorkspace(hwnd ty.HWND, newWorkspace int) {
	log.Printf("[usecase.MoveToWorkspace] hwnd=%v, newWorkspace=%v\n", hwnd, newWorkspace)
	container := core.GetContainer()
	if newWorkspace != container.ActiveWorkspace {
		win.ShowWindow(hwnd, 6)
	}

	container.MoveToWorkspace(hwnd, newWorkspace)
	container.Workspaces[container.ActiveWorkspace-1].UpdateLayout()
}

func SwitchToWorkspace(newWorkspace int) {
	log.Printf("[usecase.SwitchToWorkspace] newWorkspace=%v\n", newWorkspace)
	container := core.GetContainer()
	container.SwitchToWorkspace(newWorkspace)
	container.Workspaces[container.ActiveWorkspace-1].UpdateLayout()
}
