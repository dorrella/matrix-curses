package main

import (
	"dorrella.com/eight-queens/pkg/curses"
	"dorrella.com/eight-queens/pkg/queens"
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
