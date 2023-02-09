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
	amount := len(windows)
	if amount == 0 || amount == 1 {
		return handleZeroOrOneWindow(amount)
	}

	monitor_width := int32(w32.GetSystemMetrics(0))
	monitor_height := int32(w32.GetSystemMetrics(1) - 37)

	width := int32(monitor_width) / int32(amount)

	left := int32(0)
	top := int32(0)
	right := left + width
	bottom := monitor_height

	result := []ty.RECT{}
	for _, h := range windows {
		_frame, _ := jw32.DwmGetWindowAttribute(jw32.HWND(h), jw32.DWMWA_EXTENDED_FRAME_BOUNDS)
		frame, ok := _frame.(*jw32.RECT)
		windowRect := jw32.GetWindowRect(jw32.HWND(h))

		if ok {
			w_left := windowRect.Left - frame.Left + left
			w_top := windowRect.Top - frame.Top + top
			w_right := 2*(windowRect.Right-frame.Right) + width
			w_bottom := windowRect.Bottom - frame.Bottom + bottom

			result = append(result, ty.RECT{Left: int(w_left), Top: int(w_top), Right: int(w_right), Bottom: int(w_bottom)})
			left += width
			right += width

		}

	}

	return result
}
