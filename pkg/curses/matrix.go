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
	m.window.SetBackground(gc.ColorPair(color_Matrix))

	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			m.setBoxChar(r, c, m.matrix[r][c])
		}
	}
}

func (m *matrix) setBoxChar(r, c int, value rune) {
	height, width := m.window.MaxYX()
	new_r := r*2 + 1
	new_c := c*2 + 1

	if new_r > height {
		panic(new_r)
	}

	if new_c > width {
		panic(new_c)
	}

	m.window.ColorOn(color_Entry)
	m.window.MovePrintf(new_r, new_c, "%c", value)
	m.window.ColorOff(color_Entry)
}
