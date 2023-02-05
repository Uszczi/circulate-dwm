package usecase

import (
	"circulate/store"
	"circulate/win"
)

func PrintWorkspaceDebug() {
	store.GetContainer().PrintDebugWorkspace()
}

func MoveToWorkspace(newWorkspace int) {
	foregroundWindow := win.GetForegroundWindow()
	container := store.GetContainer()
	container.MoveToWorkspace(foregroundWindow, newWorkspace)
	container.Workspaces[container.ActiveWorkspace].UpdateLayout()
}

func SwitchToWorkspace(newWorkspace int) {
	container := store.GetContainer()
	container.SwitchToWorkspace(newWorkspace)
}
