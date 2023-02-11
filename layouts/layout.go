package layouts

import (
	"circulate/ty"
	"circulate/win"
)

var monitorWidth int
var monitorHeight int

func init() {
	monitorWidth, monitorHeight = win.GetDesktopDimentions()
}

type Layout interface {
	Add(ty.HWND)
	Calculate([]ty.HWND) []ty.RECT
}

func handleZeroOrOneWindow(amount int) []ty.RECT {
	if amount == 0 {
		return []ty.RECT{}
	}

	return []ty.RECT{{Left: 0, Top: 0, Right: monitorWidth, Bottom: monitorHeight}}
}

var AllLayouts = [...]string{"rows", "columns", "floating"}

func CreateLayout(name string) (Layout, bool) {
	layout, ok := layoutToCreate[name]
	if !ok {
		return nil, false
	}
	return layout(), ok
}

var layoutToCreate = map[string]func() Layout{
	"columns":  createColumnsLayout,
	"rows":     createRowsLayout,
	"floating": createFloatingLayout}

func createColumnsLayout() Layout {
	return &ColumnsLayout{}
}

func createRowsLayout() Layout {
	return &RowsLayout{}
}

func createFloatingLayout() Layout {
	return &FloatingLayout{}
}
