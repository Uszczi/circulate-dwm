package usecase

import (
	"circulate/core"
	"circulate/layouts"
	"circulate/store"
)

func SetColumnLayout() {
	workspace := store.GetActiveWorkspace()

	var layout layouts.Layout
	layout = &layouts.ColumnLayout{Windows: workspace.Windows}

	positions := layout.Calculate()
	core.SetWindows(workspace.Windows, positions)
}

func SetRowLayout() {
	workspace := store.GetActiveWorkspace()

	var layout layouts.Layout
	layout = &layouts.RowLayout{Windows: workspace.Windows}

	positions := layout.Calculate()
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
