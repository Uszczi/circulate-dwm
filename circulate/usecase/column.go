package usecase

import (
	"circulate/circulate/core"
	"circulate/circulate/layouts"
)

func SetColumnLayout() {
	windows := core.GetWindows()
	positions := layouts.CalculateColumns(windows)
	for _, h := range windows {
		core.PrintDebugWindow(h)
	}
	core.SetWindows(windows, positions)
}
