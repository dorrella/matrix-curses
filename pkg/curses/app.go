package curses

import (
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
	matrix *matrix
}

func NewApp(rows, cols int) *App {
	a := &App{
		matrix: newMatrix(rows, cols),
	}

	return a
}

func (a *App) Run() {
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
	w.SetBackground(gc.ColorPair(color_BG))

	w.Refresh()

	//init windows
	a.matrix.init()
	a.matrix.window.Refresh()

	_, c := a.matrix.window.MaxYX()
	total_height, total_width := w.MaxYX()
	height := total_height - 2
	con := newConsole(height)
	width := total_width - 3 - c
	con.init(height, width, 1, c+2)
	con.window.Printf("%d %d %d %d\n", height, width, 1, c+1)
	con.window.Refresh()

	//wait
	//_ = a.matrix.window.GetChar()
	_ = con.window.GetChar()
}
