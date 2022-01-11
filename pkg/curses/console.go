package curses

import (
	gc "github.com/gbin/goncurses"
)

type console struct {
	window *gc.Window
	buff   []string
}

func newConsole(size int) *console {
	b := make([]string, size)
	return &console{
		window: nil,
		buff:   b,
	}
}

func (con *console) init(h, w, r, c int) {
	window, err := gc.NewWindow(h, w, r, c)
	if err != nil {
		panic(err)
	}
	con.window = window
	con.window.SetBackground(gc.ColorPair(color_Console))
}
