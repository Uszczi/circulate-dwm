package usecase

import (
	"circulate/layouts"
	"circulate/store"
	"fmt"
	"log"
)

func SetLayout(name string) error {
	log.Printf("[usecase.SetLayout], name=%v\n", name)
	newLayout, ok := layouts.CreateLayout(name)
	if !ok {
		return fmt.Errorf("Unvalid layout {%v}", name)

	}
	workspace := store.GetActiveWorkspace()
	workspace.Layout = newLayout
	workspace.UpdateLayout()
	return nil
}

func SetPreviousLayout() {
	log.Printf("[usecase.SetPreviousLayout]\n")
}

func SetNextLayout() {
	log.Printf("[usecase.SetNextLayout]\n")
}
