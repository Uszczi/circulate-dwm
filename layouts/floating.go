package layouts

import (
	"circulate/ty"
)

type FloatingLayout struct{}

func (*FloatingLayout) Add(ty.HWND) {
	return
}

func (rl *FloatingLayout) Calculate(windows []ty.HWND) []ty.RECT {
	amount := len(windows)
	if amount == 0 || amount == 1 {
		return handleZeroOrOneWindow(amount, ty.RECT{})
	}

	return []ty.RECT{}
}
