package main

import (
	//"fmt"

	"dorrella.com/eight-queens/pkg/curses"
	"dorrella.com/eight-queens/pkg/queens"
)

func main() {

	log := make(chan string)
	app := curses.NewApp(8, 8)
	q := queens.NewQueens(log)

	go q.Run()
	app.Run(log)
}
