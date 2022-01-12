package curses

//wrapper for updating events to the matrix
type MatrixEvent struct {
	Color int16
	Char  rune
	Row   int
	Col   int
}
