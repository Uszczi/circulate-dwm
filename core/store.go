package core

import (
	"circulate/ty"
)

var w *Container

var WindowGap = ty.RECT{Left: 10, Top: 10, Bottom: 10, Right: 10}

func SetContainer(container *Container) {
	w = container
}

func GetContainer() *Container {
	return w
}

func GetActiveWorkspace() *Workspace {
	// TODO move it from there
	return w.Workspaces[w.ActiveWorkspace-1]
}
