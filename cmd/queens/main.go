package main

import (
	//"fmt"

	"dorrella.com/eight-queens/pkg/curses"
	"dorrella.com/eight-queens/pkg/queens"
)

func asdf(e chan *curses.ChessEvent) {
	ce := curses.ChessEvent{
		Color: curses.COLOR_ENTRY,
		Char:  'A',
		Row:   0,
		Col:   0,
	}

	e <- &ce
}

func main() {

	log := make(chan string)
	event := make(chan *curses.ChessEvent)
	app := curses.NewApp(8, 8)
	q := queens.NewQueens(log, event)

	go q.Run()
	//go asdf(event)
	app.Run(log, event)
}
