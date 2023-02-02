package usecase

import (
	"circulate/core"
	"circulate/layouts"
	"circulate/store"
)

func SetColumnLayout() {
	workspace := store.GetActiveWorkspace()

	positions := layouts.CalculateColumns(workspace.W)
	core.SetWindows(workspace.W, positions)
}

func SetRowLayout() {
	workspace := store.GetActiveWorkspace()

	positions := layouts.CalculateRows(workspace.W)
	core.SetWindows(workspace.W, positions)
}

func SetPreviousLayout() {
	return
}

func SetNextLayout() {
	return
}
