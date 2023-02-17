package usecase

import (
	"circulate/core"
	"circulate/layouts"
	"fmt"
	"log"
)

func SetLayout(name string) error {
	log.Printf("[usecase.SetLayout], name=%v\n", name)
	newLayout, ok := layouts.CreateLayout(name)
	if !ok {
		log.Printf("Unvalid layout {%+v}\n", name)
		return fmt.Errorf("Unvalid layout {%v}", name)

	}
	workspace := core.GetActiveWorkspace()
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
