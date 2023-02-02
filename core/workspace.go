package core

import (
	"circulate/store"
	"circulate/win"
)

func PrintWorkspaceDebug() {
	store.PrintDebugWorkspace()
}

func MoveToWorkspace(workspace int) {
	foregroundWindow := win.GetForegroundWindow()

	store.MoveToWorkspace(foregroundWindow, workspace)
}

func ShowWorkspace(workspace int) {
	foregroundWindow := win.GetForegroundWindow()

	store.MoveToWorkspace(foregroundWindow, workspace)
}

func HideWorkspace(workspace int) {
	foregroundWindow := win.GetForegroundWindow()

	store.MoveToWorkspace(foregroundWindow, workspace)
}
