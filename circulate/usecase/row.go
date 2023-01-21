package usecase

import (
    "circulate/circulate/layouts"
    "circulate/circulate/core"

)

func SetRowLayout() {
    windows := core.GetWindows()
    positions := layouts.CalculateRows(windows)
    for _, h := range windows {
        core.PrintDebugWindow(h)
    }
    core.SetWindows(windows, positions)
}
