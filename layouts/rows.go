package layouts

import (
	"circulate/ty"

	jw32 "github.com/jcollie/w32"
	"github.com/tadvi/winc/w32"
)

type RowsLayout struct{}

func (*RowsLayout) Add(ty.HWND) {
	return
}

func (rl *RowsLayout) Calculate(windows []ty.HWND) []ty.RECT {
	amount := len(windows)
	if amount == 0 || amount == 1 {
		return handleZeroOrOneWindow(amount)
	}

	monitor_width := int32(w32.GetSystemMetrics(0))
	monitor_height := int32(w32.GetSystemMetrics(1) - 37)

	height := monitor_height / int32(amount)

	left := int32(0)
	top := int32(0)
	right := monitor_width
	bottom := height

	result := []ty.RECT{}
	for _, h := range windows {
		_frame, _ := jw32.DwmGetWindowAttribute(jw32.HWND(h), jw32.DWMWA_EXTENDED_FRAME_BOUNDS)
		frame, ok := _frame.(*jw32.RECT)
		windowRect := jw32.GetWindowRect(jw32.HWND(h))

		if ok {
			w_left := windowRect.Left - frame.Left + left
			w_top := windowRect.Top - frame.Top + top
			w_right := windowRect.Right - frame.Right + right
			w_bottom := windowRect.Bottom - frame.Bottom + bottom

			result = append(result, ty.RECT{Left: int(w_left), Top: int(w_top), Right: int(w_right), Bottom: int(w_bottom)})
			top += height
		}

	}

	return result
}
