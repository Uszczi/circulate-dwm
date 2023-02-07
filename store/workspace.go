package store

import (
	"circulate/core"
	"circulate/layouts"
)

var w = &core.Container{ActiveWorkspace: 0, Workspaces: []*core.Workspace{
	{Layout: &layouts.ColumnLayout{}},
	{Layout: &layouts.ColumnLayout{}},
	{Layout: &layouts.ColumnLayout{}},
	{Layout: &layouts.ColumnLayout{}},
	{Layout: &layouts.ColumnLayout{}},
	{Layout: &layouts.ColumnLayout{}},
}}

func GetContainer() *core.Container {
	return w
}

func GetActiveWorkspace() *core.Workspace {
	return w.Workspaces[w.ActiveWorkspace]
}
