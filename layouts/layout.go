package layouts

import (
	"circulate/ty"
	"circulate/win"
	"fmt"
)

var (
	monitorWidth  int
	monitorHeight int
)

func init() {
	monitorWidth, monitorHeight = win.GetDesktopDimentions()
}

type Layout interface {
	Add(ty.HWND)
	Calculate([]ty.HWND) []ty.RECT
}

func handleZeroOrOneWindow(amount int, invisibleBorder ty.RECT) []ty.RECT {
	if amount == 0 {
		return []ty.RECT{}
	}
	fmt.Println(invisibleBorder)
	fmt.Println("monitor_width", monitorWidth)
	fmt.Println("monitor_height", monitorHeight)

	return []ty.RECT{{
		Top:    -invisibleBorder.Top,
		Right:  monitorWidth + invisibleBorder.Left - invisibleBorder.Right,
		Bottom: monitorHeight + invisibleBorder.Top - invisibleBorder.Bottom,
		Left:   -invisibleBorder.Left,
	}}
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
	"floating": createFloatingLayout,
}

func createColumnsLayout() Layout {
	return &ColumnsLayout{}
}

func createRowsLayout() Layout {
	return &RowsLayout{}
}

func createFloatingLayout() Layout {
	return &FloatingLayout{}
}
