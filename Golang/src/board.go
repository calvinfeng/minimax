package main

func NewBoard() Board {
	newBoard := Board{}
	for i := range newBoard.Grid {
		for j := range newBoard.Grid[i] {
			newBoard.Grid[i][j] = "_"
		}
	}
	return newBoard
}

type Board struct {
	Grid [3][3] string
}

func (b *Board) GetCopy() Board {
    newBoard := NewBoard()
    for i := range b.Grid {
        for j := range b.Grid[i] {
            newBoard.Grid[i][j] = b.Grid[i][j]
        }
    }
    return newBoard
}

func (board *Board) PlaceMark(i int, j int, mark string) {
	board.Grid[i][j] = mark
}

func (board *Board) IsOver() bool {
	return (board.isWon() || board.isTied())
}

func (board *Board) GetAvailablePos() [][2]int {
    availPos := [][2]int{}
    for i := range board.Grid {
        for j := range board.Grid[i] {
            if board.Grid[i][j] == "_" {
                availPos = append(availPos, [2]int{i, j})
            }
        }
    }
    return availPos
}

func (board *Board) Winner() string {
	xStreak := [3]string{"X", "X", "X"}
	oStreak := [3]string{"O", "O", "O"}

	rows := board.rows()
	for i := range rows {
		if rows[i] == xStreak {
			return "X"
		}
		if rows[i] == oStreak {
			return "O"
		}
	}

	columns := board.columns()
	for i := range columns {
		if columns[i] == xStreak {
			return "X"
		}
		if columns[i] == oStreak {
			return "O"
		}
	}

	diagonals := board.diagonals()
	for i := range diagonals {
		if diagonals[i] == xStreak {
			return "X"
		}
		if diagonals[i] == oStreak {
			return "O"
		}
	}

	return ""
}

func (board Board) String() string {
	var boardStr string = "    "
	for i := range board.Grid {
		for j := range board.Grid[i] {
			boardStr += board.Grid[i][j] + " "
		}
		boardStr += "\n    "
	}
	return boardStr
}

func (board *Board) rows() [3][3]string {
	return board.Grid
}

func (board *Board) columns() [3][3]string {
	columns := [3][3]string{}
	for i := range board.Grid {
		for j := range board.Grid[i] {
			columns[j][i] = board.Grid[i][j]
		}
	}
	return columns
}

func (board *Board) diagonals() [2][3]string {
	diagonals := [2][3]string{}
	for i := range board.Grid {
		diagonals[0][i] = board.Grid[i][i]
	}

	for i := range board.Grid {
		diagonals[1][i] = board.Grid[i][2 - i]
	}

	return diagonals
}


func (board *Board) isWon() bool {
	if board.Winner() != "" {
		return true
	}
	return false
}

func (board *Board) isTied() bool {
	spaceCount := 0
	for i := range board.Grid {
		for j := range board.Grid[i] {
			if board.Grid[i][j] == "_" {
				spaceCount += 1
			}
		}
	}

	if spaceCount == 0 && board.Winner() == "" {
		return true
	}
	return false
}
