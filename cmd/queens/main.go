package main

import (
	"dorrella.com/eight-queens/pkg/curses"
)

func asdf(log chan<- string) {

	log <- "test"
	log <- "test2"
}

func main() {
	log := make(chan string)
	app := curses.NewApp(8, 8)

	go asdf(log)

	app.Run(log)
}
