package usecase

import (
	"circulate/store"
	"circulate/win"
	"log"
)

func PrintWorkspaceDebug() {
	store.GetContainer().PrintDebugWorkspace()
}

func MoveToWorkspace(newWorkspace int) {
	foregroundWindow := win.GetForegroundWindow()
	log.Printf("Moving %+v windows to %+v workspace\n", foregroundWindow, newWorkspace)
	container := store.GetContainer()
	if newWorkspace != container.ActiveWorkspace {
		win.ShowWindow(foregroundWindow, 6)
	}

	container.MoveToWorkspace(foregroundWindow, newWorkspace)
	container.Workspaces[container.ActiveWorkspace].UpdateLayout()
}

func SwitchToWorkspace(newWorkspace int) {
	container := store.GetContainer()
	container.SwitchToWorkspace(newWorkspace)
}
