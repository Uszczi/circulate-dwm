package layouts

import (
	"circulate/ty"
	"fmt"
)

type RowsLayout struct{}

func (*RowsLayout) Add(ty.HWND) {
	return
}

func (rl *RowsLayout) Calculate(windows []ty.HWND) []ty.RECT {
	amount := len(windows)
	if amount == 0 {
		return []ty.RECT{}
	}

	invi := calculateWindowsInvisibleBorder(windows[0])
	if amount == 1 {
		return handleSingleWindow(amount, invi)
	}

	height := int32(monitorHeight) / int32(amount)
	fmt.Println(monitorHeight, amount, height)

	left := int32(0)
	top := int32(0)
	right := monitorWidth
	bottom := height

	var invisibleBorder ty.RECT
	result := []ty.RECT{}
	for _, hwnd := range windows {
		invisibleBorder = calculateWindowsInvisibleBorder(hwnd)
		rr := ty.RECT{
			Top:    -invisibleBorder.Top + int(top),
			Right:  invisibleBorder.Left - invisibleBorder.Right + int(right),
			Bottom: int(bottom),
			Left:   int(-invisibleBorder.Left) + int(left),
		}

		result = append(result, rr)
		top += height
	}
	return result
}
