package usecase

import (
	"circulate/layouts"
	"circulate/store"
)

func SetColumnLayout() {
	workspace := store.GetActiveWorkspace()
	workspace.Layout = &layouts.ColumnLayout{}
	workspace.UpdateLayout()
}

func SetRowLayout() {
	workspace := store.GetActiveWorkspace()
	workspace.Layout = &layouts.RowLayout{}
	workspace.UpdateLayout()
}

func SetPreviousLayout() {
	return
}

func SetNextLayout() {
	return
}
