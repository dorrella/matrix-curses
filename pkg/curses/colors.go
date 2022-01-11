package curses

import (
	gc "github.com/gbin/goncurses"
)

const (
	color_BG     int16 = 1
	color_Matrix       = 2
	color_Entry        = 3
)

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
	err = gc.InitPair(color_BG, gc.C_WHITE, gc.C_BLACK)
	if err != nil {
		panic(err)
	}

	//second color
	err = gc.InitPair(color_Matrix, gc.C_WHITE, gc.C_WHITE)
	if err != nil {
		panic(err)
	}

	//second color
	err = gc.InitPair(color_Entry, gc.C_YELLOW, gc.C_BLACK)
	if err != nil {
		panic(err)
	}
}
