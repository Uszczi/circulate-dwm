package layouts

import (
	"circulate/ty"
)

type Layout interface {
	Add(ty.HWND)
	Calculate([]ty.HWND) []ty.RECT
}

var AllLayouts = [...]string{"columns", "rows", "floating"}

func CreateLayout(name string) (Layout, bool) {
	layout, ok := layoutToCreate[name]
	if !ok {
		return nil, false
	}
	return layout(), ok
}

var layoutToCreate = map[string]func() Layout{
	"columns":  createFloatingLayout,
	"rows":     createRowsLayout,
	"floating": createColumnsLayout}

func createColumnsLayout() Layout {
	return &ColumnsLayout{}
}

func createRowsLayout() Layout {
	return &RowsLayout{}
}

func createFloatingLayout() Layout {
	return &FloatingLayout{}
}
