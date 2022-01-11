package queens

type Board struct {
	board [8][8]bool
}

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) IsValid(r, c int) bool {
	if r < 0 || c < 0 {
		return false
	}

	if r >= 8 || c >= 8 {
		return false
	}
	return true
}

func (b *Board) IsSafe(r, c int) bool {
	if !b.IsValid(r, c) {
		return false
	}

	for i := 0; i < 8; i++ {
		if i != c {
			//check our row against other columns
			if b.IsValid(r, i) && b.board[r][i] {
				return false
			}
		}

		if i != r {
			//check our col against other rows
			if b.IsValid(i, c) && b.board[i][c] {
				return false
			}
		}

		//diags?
		//top left
		row := r - i
		col := c - i
		if b.IsValid(row, col) && b.board[row][col] {
			return false
		}

		//top right
		row = r - i
		col = c + i
		if b.IsValid(row, col) && b.board[row][col] {
			return false
		}

		//diags?
		//bottom left
		row = r + i
		col = c - i
		if b.IsValid(row, col) && b.board[row][col] {
			return false
		}

		//bottom right
		row = r + i
		col = c + i
		if b.IsValid(row, col) && b.board[row][col] {
			return false
		}
	}

	return true
}
