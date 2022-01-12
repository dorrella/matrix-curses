package queens

import (
	"fmt"
	"time"

	"github.com/dorrella/matrix-curses/pkg/curses"
)

type Queens struct {
	b     *Board                   //chess board
	log   chan string              //channels
	event chan *curses.MatrixEvent //chesboard event
}

// Return wapper around interface
func NewQueens(log chan string, event chan *curses.MatrixEvent) *Queens {
	q := Queens{
		b:     NewBoard(),
		log:   log,
		event: event,
	}

	return &q
}

//Run 8 queens
//todo, choose 1 pos
func (q *Queens) Run() {
	_ = q.step(0)
	close(q.log)
	close(q.event)
}

//helper function to do the work
//true on success
func (q *Queens) step(row int) bool {
	time.Sleep(time.Millisecond * 500)
	//exit condition is starting row 9 of an 8x8 board
	if row >= 8 {
		q.log <- "success"
		return true
	}

	for i := 0; i < 8; i++ {
		//only iterate if it is safe to place the queen in the next spot
		if q.b.IsSafe(row, i) {
			q.log <- fmt.Sprintf("trying %d,%d\n", row, i)
			q.add(row, i)

			if q.step(row + 1) {
				//success
				return true
			}

			// remove queen and continue
			q.log <- fmt.Sprintf("failed %d, %d\n", row, i)
			q.del(row, i)
		}

	}

	// we have exhausted all open spaces on this row
	return false
}

//update position
func (q *Queens) add(row, col int) {
	q.b.board[row][col] = true

	me := curses.MatrixEvent{
		Color: curses.COLOR_ENTRY,
		Char:  'Q',
		Row:   row,
		Col:   col,
	}

	q.event <- &me
}

//undo add
func (q *Queens) del(row, col int) {
	q.b.board[row][col] = false

	me := curses.MatrixEvent{
		Color: curses.COLOR_ENTRY,
		Char:  ' ',
		Row:   row,
		Col:   col,
	}

	q.event <- &me
}
