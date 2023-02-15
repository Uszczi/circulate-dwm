package layouts

import (
	"circulate/ty"
	"fmt"

	jw32 "github.com/jcollie/w32"
	"github.com/tadvi/winc/w32"
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
	monitor_width := int32(w32.GetSystemMetrics(0))
	monitor_height := int32(w32.GetSystemMetrics(1) - 37)
	invisibleBorder := invisibleWindowsBorder

	fmt.Println("monitor_width", monitor_width)
	fmt.Println("monitor_height", monitor_height)

	if amount == 1 {
		retust := handleZeroOrOneWindow(amount, invisibleWindowsBorder)
		fmt.Println(monitor_width)
		fmt.Println(monitor_height)

		fmt.Println("handleZeroOrOneWindow", retust)
		return retust
	}

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
		fmt.Println("Najpierw fame ozniej windowRect")
		fmt.Println(frame)
		fmt.Println(windowRect)

		if ok {
			// w_left := windowRect.Left - frame.Left + left
			// w_top := windowRect.Top - frame.Top + top
			// w_right := windowRect.Right - frame.Right + right
			// w_bottom := windowRect.Bottom - frame.Bottom + bottom

			rr := ty.RECT{
				Top:    -invisibleBorder.Top + int(top),
				Right:  invisibleBorder.Left - invisibleBorder.Right + int(right),
				Bottom: invisibleBorder.Top - invisibleBorder.Bottom + int(bottom),
				Left:   int(-invisibleBorder.Left) + int(left),
			}

			result = append(result, rr)
			top += height
		}

	}

	return result
}
