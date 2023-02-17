package store

import (
	"circulate/core"
	"circulate/ty"
)

var w *core.Container

var WindowGap = ty.RECT{Left: 10, Top: 10, Bottom: 10, Right: 10}

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
