package usecase

import (
	"circulate/store"
	"circulate/ty"
	"circulate/win"
	"log"
)

func PrintWorkspaceDebug() {
	log.Println("[usecase.PrintWorkspaceDebug]")
	store.GetContainer().PrintDebugWorkspace()
}

func ClearWorkspace() {
	log.Println("[usecase.ClearWorkspace]")
	store.GetActiveWorkspace().WHWND = []ty.HWND{}

}

func MoveToWorkspace(hwnd ty.HWND, newWorkspace int) {
	log.Printf("[usecase.MoveToWorkspace] hwnd=%v, newWorkspace=%v\n", hwnd, newWorkspace)
	container := store.GetContainer()
	if newWorkspace != container.ActiveWorkspace {
		win.ShowWindow(hwnd, 6)
	}

	container.MoveToWorkspace(hwnd, newWorkspace)
	container.Workspaces[container.ActiveWorkspace].UpdateLayout()
}

func SwitchToWorkspace(newWorkspace int) {
	log.Printf("[usecase.SwitchToWorkspace] newWorkspace=%v\n", newWorkspace)
	container := store.GetContainer()
	container.SwitchToWorkspace(newWorkspace)
}
