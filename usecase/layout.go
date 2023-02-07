package usecase

import (
	"circulate/layouts"
	"circulate/store"
	"log"
)

func SetColumnLayout() {
	log.Printf("[usecase.SetColumnLayout]\n")
	workspace := store.GetActiveWorkspace()
	workspace.Layout = &layouts.ColumnLayout{}
	workspace.UpdateLayout()
}

func SetRowLayout() {
	log.Printf("[usecase.SetRowLayout]\n")
	workspace := store.GetActiveWorkspace()
	workspace.Layout = &layouts.RowLayout{}
	workspace.UpdateLayout()
}

func SetPreviousLayout() {
	log.Printf("[usecase.SetPreviousLayout]\n")
}

func SetNextLayout() {
	log.Printf("[usecase.SetNextLayout]\n")
}
