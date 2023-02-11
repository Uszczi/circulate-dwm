package usecase

import (
	"circulate/core"
	"circulate/layouts"
	"circulate/store"
	"log"
)

func Setup() {
	log.Printf("[usecase.Setup]\n")

	store.SetContainer(&core.Container{ActiveWorkspace: 1, Workspaces: []*core.Workspace{
		{Layout: &layouts.ColumnsLayout{}},
		{Layout: &layouts.ColumnsLayout{}},
		{Layout: &layouts.ColumnsLayout{}},
		{Layout: &layouts.ColumnsLayout{}},
		{Layout: &layouts.ColumnsLayout{}},
		{Layout: &layouts.ColumnsLayout{}},
		{Layout: &layouts.ColumnsLayout{}},
		{Layout: &layouts.ColumnsLayout{}},
		{Layout: &layouts.ColumnsLayout{}}}})

	hwnds := core.GetWindows()

	workspace := store.GetActiveWorkspace()
	for _, hwnd := range hwnds {
		log.Println("[usecase.Setup] adding hwnd: ", hwnd)
		workspace.AddWindow(hwnd)
	}
	workspace.UpdateLayout()
}
