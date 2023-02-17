package layouts

import (
	"circulate/ty"
)

type ColumnsLayout struct{}

func (*ColumnsLayout) Add(ty.HWND) {
	return
}

func (cl *ColumnsLayout) Calculate(windows []ty.HWND) []ty.RECT {
	amount := len(windows)
	if amount == 0 {
		return []ty.RECT{}
	}

	invi := calculateWindowsInvisibleBorder(windows[0])
	if amount == 1 {
		return handleSingleWindow(amount, invi)
	}

	width := int32(monitorWidth) / int32(amount)

	left := int32(0)
	top := int32(0)
	right := width
	bottom := monitorHeight

	var invisibleBorder ty.RECT
	result := []ty.RECT{}
	for _, hwnd := range windows {
		invisibleBorder = calculateWindowsInvisibleBorder(hwnd)
		rr := ty.RECT{
			Left:   int(-invisibleBorder.Left) + int(left) + WindowGap.Left,
			Top:    -invisibleBorder.Top + int(top) + WindowGap.Top,
			Right:  invisibleBorder.Left - invisibleBorder.Right + int(right) - WindowGap.Right - WindowGap.Left,
			Bottom: int(bottom) - WindowGap.Bottom - WindowGap.Top,
		}

		result = append(result, rr)
		left += width
		// right += width

	}

	return result
}
