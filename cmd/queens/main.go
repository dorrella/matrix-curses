package main

import (
	"github.com/dorrella/matrix-curses/pkg/curses"
	"github.com/dorrella/matrix-curses/pkg/queens"
)

// Visualize 8 queens with ncurses
func main() {

	log := make(chan string)
	event := make(chan *curses.MatrixEvent)
	app := curses.NewApp(8, 8)
	q := queens.NewQueens(log, event)

	go q.Run()
	app.Run(log, event)
}
