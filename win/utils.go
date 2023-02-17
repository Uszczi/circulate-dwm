package win

const toolbar_height_px = 40

func GetDesktopWidth() int {
	monitor_width := GetSystemMetrics(SM_CXSCREEN)
	return monitor_width
}

func GetDesktopHeight() int {
	monitor_height := GetSystemMetrics(SM_CYSCREEN)
	return monitor_height - toolbar_height_px
}

func GetDesktopDimentions() (int, int) {
	return GetDesktopWidth(), GetDesktopHeight()
}
