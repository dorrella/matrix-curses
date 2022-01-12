package curses

import (
	gc "github.com/gbin/goncurses"
)

func getBoxSize(r, c int) (int, int) {
	//err on 0x0?
	new_r := r*2 + 1
	new_c := c*2 + 1
	return new_r, new_c
}

type matrix struct {
	window *gc.Window // window
	matrix [][]rune   // matrix
	rows   int
	cols   int
}

func newMatrix(rows, cols int) *matrix {
	new_h, new_w := getBoxSize(rows, cols)

	m := make([][]rune, new_h)
	for r := 0; r < new_h; r++ {
		m[r] = make([]rune, new_w)
		for c := 0; c < new_w; c++ {
			m[r][c] = ' '
		}
	}

	return &matrix{
		window: nil,
		matrix: m,
		rows:   rows,
		cols:   cols,
	}
}

func (m *matrix) init() {
	new_h, new_w := getBoxSize(m.rows, m.cols)
	window, err := gc.NewWindow(new_h, new_w, 1, 1)
	if err != nil {
		panic(err)
	}
	m.window = window
	m.window.SetBackground(gc.ColorPair(COLOR_MATRIX))

	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			ce := ChessEvent{
				Row:   r,
				Col:   c,
				Char:  m.matrix[r][c],
				Color: COLOR_ENTRY,
			}
			m.setBoxChar(&ce)
		}
	}

	end_r := new_h - 1
	end_c := new_w - 1
	for r := 0; r < new_h; r++ {
		even_row := r%2 == 0
		for c := 0; c < new_w; c++ {
			even_col := c%2 == 0
			if !even_row && !even_col {
				//non box space
				continue
			}

			//if we prit a bullet, something is not being handled
			var char gc.Char = gc.ACS_BULLET

			if r == 0 && c == 0 {
				// top left corner
				// should be gc.ACS_ULCORNER, but there is a bug
				char = gc.ACS_LLCORNER
			} else if r == 0 && c == end_c {
				// top right corner
				char = gc.ACS_URCORNER
			} else if r == end_r && c == 0 {
				// bottom left corner
				// should be gc.ACS_LLCORNER
				char = gc.ACS_ULCORNER
			} else if r == end_r && c == end_c {
				// bottom right corner
				char = gc.ACS_LRCORNER
			} else if r == 0 {
				//top edge
				char = gc.ACS_HLINE
				if even_col {
					//top edge tee
					char = gc.ACS_TTEE
				}
			} else if r == end_r {
				//bottom edge
				char = gc.ACS_HLINE
				if even_col {
					//bottom edge tee
					char = gc.ACS_BTEE
				}
			} else if c == 0 {
				//left edge
				char = gc.ACS_VLINE
				if even_row {
					//left edge tee
					char = gc.ACS_LTEE
				}
			} else if c == end_c {
				//right edge
				char = gc.ACS_VLINE
				if even_row {
					//right edge tee
					char = gc.ACS_RTEE
				}
			} else {
				//inner edges
				if even_row && even_col {
					// +
					char = gc.ACS_PLUS
				} else if even_row {
					// even rows are horizontal lines
					char = gc.ACS_HLINE
				} else {
					// rest are verticle
					char = gc.ACS_VLINE
				}
			}

			m.window.ColorOn(COLOR_MATRIX)
			//m.window.MovePrintf(r, c, "%s", char)
			m.window.Move(r, c)
			m.window.AddChar(char)
			m.window.ColorOff(COLOR_MATRIX)
		}
	}

}

func (m *matrix) setBoxChar(ce *ChessEvent) {
	height, width := m.window.MaxYX()
	new_r := ce.Row*2 + 1
	new_c := ce.Col*2 + 1

	if new_r > height {
		panic(new_r)
	}

	if new_c > width {
		panic(new_c)
	}

	m.matrix[ce.Row][ce.Col] = ce.Char

	m.window.ColorOn(ce.Color)
	m.window.MovePrintf(new_r, new_c, "%c", ce.Char)
	m.window.ColorOff(ce.Color)
}
