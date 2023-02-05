package usecase

import (
	"circulate/core"
	"circulate/layouts"
	"circulate/store"
)

func SetColumnLayout() {
	workspace := store.GetActiveWorkspace()

	var layout layouts.Layout
	layout = &layouts.ColumnLayout{}

	if len(workspace.WHWND) > 0 {
		positions := layout.Calculate(workspace.WHWND)
		core.SetWindows(workspace.WHWND, positions)
	}
}

func SetRowLayout() {
	workspace := store.GetActiveWorkspace()

	var layout layouts.Layout
	layout = &layouts.RowLayout{}

	positions := layout.Calculate(workspace.WHWND)
	core.SetWindows(workspace.WHWND, positions)
}

func SetPreviousLayout() {
	return
}

func SetNextLayout() {
	return
}
