package queens

import (
	"fmt"
	"time"

	"dorrella.com/eight-queens/pkg/curses"
)

type Queens struct {
	b     *Board                  //chess board
	log   chan string             //channels
	event chan *curses.ChessEvent //chesboard event
}

func NewQueens(log chan string, event chan *curses.ChessEvent) *Queens {
	q := Queens{
		b:     NewBoard(),
		log:   log,
		event: event,
	}

	return &q
}

func (q *Queens) Run() {
	_ = q.step(0)
	close(q.log)
	close(q.event)
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
			q.add(row, i)

			if q.step(row + 1) {
				return true
			}

			q.log <- fmt.Sprintf("failed %d, %d\n", row, i)
			q.del(row, i)
		}

	}

	return false
}

func (q *Queens) add(row, col int) {
	q.b.board[row][col] = true
	ce := curses.ChessEvent{
		Color: curses.COLOR_ENTRY,
		Char:  'Q',
		Row:   row,
		Col:   col,
	}
	q.event <- &ce
}

func (q *Queens) del(row, col int) {
	q.b.board[row][col] = false
	ce := curses.ChessEvent{
		Color: curses.COLOR_ENTRY,
		Char:  ' ',
		Row:   row,
		Col:   col,
	}
	q.event <- &ce
}
