package curses

import (
	gc "github.com/gbin/goncurses"
)

const (
	COLOR_BG      int16 = 1
	COLOR_MATRIX        = 2
	COLOR_ENTRY         = 3
	COLOR_CONSOLE       = 4
)

//set color pairs
func setColors() {
	//colors
	if !gc.HasColors() {
		panic("no color terminal")
	}
	err := gc.StartColor()
	if err != nil {
		panic(err)
	}

	//set background
	err = gc.InitPair(COLOR_BG, gc.C_WHITE, gc.C_BLACK)
	if err != nil {
		panic(err)
	}

	//border
	err = gc.InitPair(COLOR_MATRIX, gc.C_WHITE, gc.C_BLACK)
	if err != nil {
		panic(err)
	}

	//entry color
	err = gc.InitPair(COLOR_ENTRY, gc.C_RED, gc.C_BLACK)
	if err != nil {
		panic(err)
	}

	//console
	err = gc.InitPair(COLOR_CONSOLE, gc.C_BLACK, gc.C_WHITE)
	if err != nil {
		panic(err)
	}
}
