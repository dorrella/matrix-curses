package curses

type ChessEvent struct {
	Color int16
	Char  rune
	Row   int
	Col   int
}
