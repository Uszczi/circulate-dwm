package layouts

import (
	"circulate/ty"

	jw32 "github.com/jcollie/w32"
	"github.com/tadvi/winc/w32"
)

type RowLayout struct {
	Windows []ty.HWND
}

func (*RowLayout) Add(ty.HWND) {
	return
}

func (rl *RowLayout) Calculate() []RECT {
	windows := rl.Windows
	amount := int32(len(windows))
	result := []RECT{}

	monitor_width := int32(w32.GetSystemMetrics(0))
	monitor_height := int32(w32.GetSystemMetrics(1) - 37)

	if amount == 1 {
		return append(result, RECT{Left: 0, Top: 0, Right: monitor_width, Bottom: monitor_height})
	}

	height := monitor_height / amount

	left := int32(0)
	top := int32(0)
	right := monitor_width
	bottom := height

	for _, h := range windows {
		_frame, _ := jw32.DwmGetWindowAttribute(jw32.HWND(h), jw32.DWMWA_EXTENDED_FRAME_BOUNDS)
		frame, ok := _frame.(*jw32.RECT)
		windowRect := jw32.GetWindowRect(jw32.HWND(h))

		if ok {
			w_left := windowRect.Left - frame.Left + left
			w_top := windowRect.Top - frame.Top + top
			w_right := windowRect.Right - frame.Right + right
			w_bottom := windowRect.Bottom - frame.Bottom + bottom

			result = append(result, RECT{Left: w_left, Top: w_top, Right: w_right, Bottom: w_bottom})
			top += height
		}

	}

	return result
}
