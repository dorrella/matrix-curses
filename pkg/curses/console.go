package curses

import (
	"dorrella.com/eight-queens/pkg/ringbuff"
	gc "github.com/gbin/goncurses"
)

type console struct {
	window *gc.Window
	buff   *ringbuff.RingBuff
}

func newConsole(size int) *console {
	b := ringbuff.NewRingBuff(size)
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
	con.window.SetBackground(gc.ColorPair(COLOR_CONSOLE))
}

func (con *console) AddMessage(s string) {
	win := con.window
	buf := con.buff

	//todo split on length?
	buf.Add(s)

	win.Erase()
	for i := 0; i < buf.GetSize(); i++ {
		//fix
		line, _ := buf.Get(i)
		win.ColorOn(COLOR_CONSOLE)
		win.MovePrint(i, 0, line)
		win.ColorOff(COLOR_CONSOLE)
	}
	win.Refresh()

}
