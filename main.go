package main

import (
	"fmt"

	"dorrella.com/eight-queens/pkg/curses"
)

func main() {
	fmt.Println("hello")

	app := curses.NewApp(8, 8)
	app.Run()
}
