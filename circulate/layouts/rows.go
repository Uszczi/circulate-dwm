// package layouts
//
// import (
// 	"syscall"
//
// 	jw32 "github.com/jcollie/w32"
// 	"github.com/tadvi/winc/w32"
// )
//
// func CalculateRows(container []syscall.Handle) []w32.RECT {
// 	amount := len(container)
// 	result := []w32.RECT{}
//
// 	monitor_width := w32.GetSystemMetrics(0)
// 	monitor_height := w32.GetSystemMetrics(1) - 37
//
// 	if amount == 1 {
// 		return append(result, w32.RECT{Left: 0, Top: 0, Right: int32(monitor_width), Bottom: int32(monitor_height)})
// 	}
//
// 	height := monitor_height / amount
//
// 	left := 0
// 	top := 0
// 	right := monitor_width
// 	bottom := top + height
//
// 	for _, h := range container {
// 		frame, _ := jw32.DwmGetWindowAttribute(jw32.HWND(h), jw32.DWMWA_EXTENDED_FRAME_BOUNDS)
// 		windowRect := jw32.GetWindowRect(jw32.HWND(h))
//
// 		frame2, ok := frame.(*jw32.RECT)
// 		if ok {
// 			w_left := int(windowRect.Left) - int(frame2.Left) + left
// 			w_top := int(windowRect.Top) - int(frame2.Top) + top
// 			w_right := 2*(int(windowRect.Right)-int(frame2.Right)) + right
// 			w_bottom := 1*(int(windowRect.Bottom)-int(frame2.Bottom)) + height
//
// 			result = append(result, w32.RECT{Left: int32(w_left), Top: int32(w_top), Right: int32(w_right), Bottom: int32(w_bottom)})
// 			top += height
// 			bottom += height
//
// 		}
//
// 	}
//
// 	return result
// }

package layouts

import (
	jw32 "github.com/jcollie/w32"
	"github.com/tadvi/winc/w32"
)

func CalculateRows(windows []uintptr) []RECT {
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
