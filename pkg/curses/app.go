package curses

import (
	"fmt"
	"os"
	"os/signal"

	gc "github.com/gbin/goncurses"
)

//catch sigint
func catchSignal() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
	gc.End()
	os.Exit(0)
}

type App struct {
	matrix  *matrix
	console *console
}

func NewApp(rows, cols int) *App {
	a := &App{
		matrix: newMatrix(rows, cols),
	}

	return a
}

func (a *App) Run(log <-chan string, event <-chan *ChessEvent) {
	w, err := gc.Init()
	defer gc.End()
	if err != nil {
		panic("failed to init")
	}

	//catch sigterm
	go catchSignal()

	//disble cursor
	gc.Cursor(0)
	gc.Echo(false)

	//init colors
	setColors()
	w.SetBackground(gc.ColorPair(COLOR_BG))

	w.Refresh()

	//init windows
	a.matrix.init()
	a.matrix.window.Refresh()

	_, c := a.matrix.window.MaxYX()
	total_height, total_width := w.MaxYX()
	height := total_height - 2
	width := total_width - 3 - c
	msg := fmt.Sprintf("initializing %dx%d at (%d,%d)\n", height, width, 1, c+1)

	a.console = newConsole(height)
	a.console.init(height, width, 1, c+2)

	a.console.AddMessage(msg)
	a.console.AddMessage("")

	//wait
	for true {
		select {
		case line, ok := <-log:
			if !ok {
				break
			}
			a.console.AddMessage(line)
		case e, ok := <-event:
			if !ok {
				break
			}
			a.matrix.setBoxChar(e)
			a.matrix.window.Refresh()
		}
	}
}
