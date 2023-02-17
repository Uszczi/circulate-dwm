package layouts

import (
	"circulate/ty"
)

type FloatingLayout struct{}

func (*FloatingLayout) Add(ty.HWND) {
	return
}

func (rl *FloatingLayout) Calculate(windows []ty.HWND) []ty.RECT {
	return []ty.RECT{}
}
