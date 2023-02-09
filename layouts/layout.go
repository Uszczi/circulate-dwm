package layouts

import (
	"circulate/ty"
	"circulate/win"

	"github.com/jcollie/w32"
)

const toolbar_height_px = 37

type Layout interface {
	Add(ty.HWND)
	Calculate([]ty.HWND) []ty.RECT
}

func handleZeroOrOneWindow(amount int) []ty.RECT {
	if amount == 0 {
		return []ty.RECT{}
	}

	monitor_width := int32(w32.GetSystemMetrics(win.SM_CXSCREEN))
	monitor_height := int32(w32.GetSystemMetrics(win.SM_CYSCREEN) - toolbar_height_px)
	return []ty.RECT{{Left: 0, Top: 0, Right: monitor_width, Bottom: monitor_height}}
}

// Wzorzec fabryka
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
