package layouts

import (
	"circulate/ty"
)

type Layout interface {
	Add(ty.HWND)
	Calculate([]ty.HWND) []ty.RECT
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
