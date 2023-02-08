package layouts

import (
	"circulate/ty"

	jw32 "github.com/jcollie/w32"
	"github.com/tadvi/winc/w32"
)

type ColumnsLayout struct{}

func (*ColumnsLayout) Add(ty.HWND) {
	return
}

func (cl *ColumnsLayout) Calculate(windows []ty.HWND) []ty.RECT {
	amount := int(len(windows))
	result := []ty.RECT{}

	if amount == 0 {
		return result
	}

	monitor_width := int32(w32.GetSystemMetrics(0))
	monitor_height := int32(w32.GetSystemMetrics(1))
	if amount == 1 {
		return append(result, ty.RECT{Left: 0, Top: 0, Right: monitor_width, Bottom: monitor_height})

	}

	width := int32(monitor_width) / int32(amount)

	left := int32(0)
	top := int32(0)
	right := left + width
	bottom := monitor_height

	for _, h := range windows {
		_frame, _ := jw32.DwmGetWindowAttribute(jw32.HWND(h), jw32.DWMWA_EXTENDED_FRAME_BOUNDS)
		frame, ok := _frame.(*jw32.RECT)
		windowRect := jw32.GetWindowRect(jw32.HWND(h))

		if ok {
			w_left := windowRect.Left - frame.Left + left
			w_top := windowRect.Top - frame.Top + top
			w_right := 2*(windowRect.Right-frame.Right) + width
			w_bottom := windowRect.Bottom - frame.Bottom + bottom

			result = append(result, ty.RECT{Left: w_left, Top: w_top, Right: w_right, Bottom: w_bottom})
			left += width
			right += width

		}

	}

	return result
}
