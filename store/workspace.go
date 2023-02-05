package store

import (
	"circulate/core"
)

var w = &core.Container{ActiveWorkspace: 0, Workspaces: []*core.Workspace{
	{}, {}, {}, {}, {}, {},
}}

func GetContainer() *core.Container {
	return w
}
func GetActiveWorkspace() *core.Workspace {
	return w.Workspaces[w.ActiveWorkspace]
}
