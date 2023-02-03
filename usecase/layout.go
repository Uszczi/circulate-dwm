package usecase

import (
	"circulate/core"
	"circulate/layouts"
	"circulate/store"
)

func SetColumnLayout() {
	workspace := store.GetActiveWorkspace()

	positions := layouts.CalculateColumns(workspace.Windows)
	core.SetWindows(workspace.Windows, positions)
}

func SetRowLayout() {
	workspace := store.GetActiveWorkspace()

	positions := layouts.CalculateRows(workspace.Windows)
	core.SetWindows(workspace.Windows, positions)
}

func SetPreviousLayout() {
	return
}

func SetNextLayout() {
	return
}

func SwitchToLayout(workspaceName int) {
	store.SwitchToLayout(workspaceName)
}
