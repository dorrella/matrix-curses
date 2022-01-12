package curses

import (
	"github.com/dorrella/matrix-curses/pkg/ringbuff"
	gc "github.com/gbin/goncurses"
)

type console struct {
	window *gc.Window
	buff   *ringbuff.RingBuff
}

//size is used for the ring buffer
//should be height of window
func newConsole(size int) *console {
	b := ringbuff.NewRingBuff(size)
	return &console{
		window: nil,
		buff:   b,
	}
}

//initialize console window
func (con *console) init(h, w, r, c int) {
	window, err := gc.NewWindow(h, w, r, c)
	if err != nil {
		panic(err)
	}
	con.window = window
	con.window.SetBackground(gc.ColorPair(COLOR_CONSOLE))
}

//adds message to console
func (con *console) AddMessage(s string) {
	win := con.window
	buf := con.buff

	//todo split on length >= width?
	buf.Add(s)

	//easier to clear the console once
	win.Erase()
	for i := 0; i < buf.GetSize(); i++ {
		line, _ := buf.Get(i)
		win.ColorOn(COLOR_CONSOLE)
		win.MovePrint(i, 0, line)
		win.ColorOff(COLOR_CONSOLE)
	}
	win.Refresh()

}
