package usecase

import (
	"circulate/store"
	"circulate/ty"
	"circulate/win"
	"log"
)

func PrintWorkspaceDebug() {
	store.GetContainer().PrintDebugWorkspace()
}

func ClearWorkspace() {
	store.GetActiveWorkspace().WHWND = []ty.HWND{}

}

// foregroundWindow := win.GetForegroundWindow()

func MoveToWorkspace(hwnd ty.HWND, newWorkspace int) {

	log.Printf("Moving %+v windows to %+v workspace\n", hwnd, newWorkspace)
	container := store.GetContainer()
	if newWorkspace != container.ActiveWorkspace {
		win.ShowWindow(hwnd, 6)
	}

	container.MoveToWorkspace(hwnd, newWorkspace)
	container.Workspaces[container.ActiveWorkspace].UpdateLayout()
}

func SwitchToWorkspace(newWorkspace int) {
	container := store.GetContainer()
	container.SwitchToWorkspace(newWorkspace)
}
