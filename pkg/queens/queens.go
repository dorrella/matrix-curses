package queens

import (
	"fmt"
	"time"
)

type Queens struct {
	b   *Board      //chess board
	log chan string //channels
}

func NewQueens(log chan string) *Queens {
	q := Queens{
		b:   NewBoard(),
		log: log,
	}

	return &q
}

func (q *Queens) Run() {
	_ = q.step(0)
}

func (q *Queens) step(row int) bool {
	time.Sleep(time.Millisecond * 500)
	if row >= 8 {
		q.log <- "success"
		return true
	}

	for i := 0; i < 8; i++ {
		if q.b.IsSafe(row, i) {
			q.log <- fmt.Sprintf("trying %d,%d\n", row, i)
			q.b.board[row][i] = true

			if q.step(row + 1) {
				return true
			}

			q.log <- fmt.Sprintf("failed %d, %d\n", row, i)
			q.b.board[row][i] = false
		}

	}

	return false
}
