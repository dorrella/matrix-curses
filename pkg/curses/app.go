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

	//wait
	a.matrix.window.Refresh()
	_ = a.matrix.window.GetChar()

}
