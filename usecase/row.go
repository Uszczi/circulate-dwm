package usecase

import (
	"circulate/core"
	"circulate/layouts"
)

func SetRowLayout() {
	windows := core.GetWindows()
	positions := layouts.CalculateRows(windows)
	for _, h := range windows {
		core.PrintDebugWindow(h)
	}
	core.SetWindows(windows, positions)
}
