package store

import (
	"circulate/core"
)

var w *core.Container

func SetContainer(container *core.Container) {
	w = container
}

func GetContainer() *core.Container {
	return w
}

func GetActiveWorkspace() *core.Workspace {
	// TODO move it from there
	return w.Workspaces[w.ActiveWorkspace-1]
}
