package usecase

import (
    "circulate/circulate/layouts"
    "circulate/circulate/core"

)

func SetColumnLayout() {
    windows := core.GetWindows()
    positions := layouts.CalculateColumns(windows)
    for _, h := range windows {
        core.PrintDebugWindow(h)
    }
    core.SetWindows(windows, positions)
}
