package layouts

import (
	"circulate/ty"
	"circulate/win"
	"fmt"

	jw32 "github.com/jcollie/w32"
)

var (
	AllLayouts    = [...]string{"rows", "columns", "floating"}
	monitorWidth  int
	monitorHeight int
	WindowGap     = ty.RECT{Left: 7, Top: 10, Bottom: 10, Right: 7}
)

func init() {
	monitorWidth, monitorHeight = win.GetDesktopDimentions()
}

type Layout interface {
	Add(ty.HWND)
	Calculate([]ty.HWND) []ty.RECT
}

func handleSingleWindow(amount int, invisibleBorder ty.RECT) []ty.RECT {
	fmt.Printf("%+v\n", invisibleBorder)
	return []ty.RECT{{
		Left:   -invisibleBorder.Left + WindowGap.Left,
		Top:    -invisibleBorder.Top + WindowGap.Top,
		Right:  monitorWidth + invisibleBorder.Left - invisibleBorder.Right - WindowGap.Right - WindowGap.Left,
		Bottom: monitorHeight - invisibleBorder.Bottom - WindowGap.Bottom - WindowGap.Top,
	}}
}

func calculateWindowsInvisibleBorder(hwnd ty.HWND) ty.RECT {
	_frame, _ := jw32.DwmGetWindowAttribute(jw32.HWND(hwnd), jw32.DWMWA_EXTENDED_FRAME_BOUNDS)
	frame, _ := _frame.(*jw32.RECT)
	windowRect := jw32.GetWindowRect(jw32.HWND(hwnd))
	invisibleWindowsBorder := ty.RECT{
		Top:    int(frame.Top) - int(windowRect.Top),
		Right:  int(frame.Right) - int(windowRect.Right),
		Bottom: int(frame.Bottom) - int(windowRect.Bottom),
		Left:   int(frame.Left) - int(windowRect.Left),
	}
	return invisibleWindowsBorder
}

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
