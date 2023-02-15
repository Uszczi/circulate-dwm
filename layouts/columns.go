package layouts

import (
	"circulate/ty"

	jw32 "github.com/jcollie/w32"
)

type ColumnsLayout struct{}

func (*ColumnsLayout) Add(ty.HWND) {
	return
}

func (cl *ColumnsLayout) Calculate(windows []ty.HWND) []ty.RECT {
	amount := len(windows)
	if amount == 0 || amount == 1 {
		return handleZeroOrOneWindow(amount, ty.RECT{})
	}

	width := int32(monitorWidth) / int32(amount)

	h := windows[0]
	_frame, _ := jw32.DwmGetWindowAttribute(jw32.HWND(h), jw32.DWMWA_EXTENDED_FRAME_BOUNDS)
	frame, _ := _frame.(*jw32.RECT)
	windowRect := jw32.GetWindowRect(jw32.HWND(h))
	invisibleWindowsBorder := ty.RECT{
		Top:    int(frame.Top) - int(windowRect.Top),
		Right:  int(frame.Right) - int(windowRect.Right),
		Bottom: int(frame.Bottom) - int(windowRect.Bottom),
		Left:   int(frame.Left) - int(windowRect.Left),
	}
	invisibleBorder := invisibleWindowsBorder

	left := int32(0)
	top := int32(0)
	right := width
	bottom := monitorHeight

	result := []ty.RECT{}
	for range windows {
		rr := ty.RECT{
			Top:    -invisibleBorder.Top + int(top),
			Right:  invisibleBorder.Left - invisibleBorder.Right + int(right),
			Bottom: invisibleBorder.Top - invisibleBorder.Bottom + int(bottom),
			Left:   int(-invisibleBorder.Left) + int(left),
		}

		result = append(result, rr)
		left += width
		// right += width

	}

	return result
}
