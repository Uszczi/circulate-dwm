package usecase

import (
	"circulate/core"
	"circulate/store"
	"log"
)

func Setup() {
	log.Printf("[usecase.Setup]\n")

	hwnds := core.GetWindows()
	log.Println("[usecase.Setup]")

	workspace := store.GetActiveWorkspace()
	for _, hwnd := range hwnds {
		workspace.AddWindow(hwnd)
	}
	workspace.UpdateLayout()
}