package connect4

import (
	"errors"
	"fmt"
)

type Board struct {
	positions [6][7]string
	connected int
}

type position struct {
	checked bool
	matched bool
}

func (b *Board) move(column int, row int, player string) (bool, error) {
	if position := b.positions[column][row]; position != "" {
		return false, errors.New(fmt.Sprintf("Position already taken by: %s", position))
	} else {
		b.positions[column][row] = player
	}

	b.CheckHorizontal(row, column)
	if b.connected >= 3 {
		return true, nil
	} else {
		b.connected = 0
	}

	b.checkVertical(column, row)
	if b.connected >= 4 {
		return true, nil
	} else {
		b.connected = 0
	}

	b.checkDiagonal(column, row)
	if b.connected >= 4 {
		return true, nil
	} else {
		b.connected = 0
	}

	return false, nil
}

// CheckHorizontal control function for horizonal matching
func (b *Board) CheckHorizontal(row int, column int) {
	b.checkLeft(row, column)
	if b.connected == 4 {
		return
	}

	b.checkRight(row, column)
}

// Recursive function to check positions to the left
func (b *Board) checkLeft(row int, column int) {
	if b.connected < 4 && column > 0 {
		if b.positions[row][column] == b.positions[row][column-1] {
			b.connected++
			b.checkLeft(row, column-1)
		}
	}
}

// Recursive function to check positions to the right
func (b *Board) checkRight(row int, column int) {
	if b.connected < 4 && column < 6 {
		if b.positions[row][column] == b.positions[row][column+1] {
			b.connected++
			b.checkRight(row, column+1)
		}
	}
}

func (b *Board) checkVertical(column int, row int) {
	if b.connected < 4 && row > 0 {
		if b.positions[column][row] == b.positions[column][row-1] {
			b.connected++
			b.checkVertical(column, row-1)
		}
	}

	if b.connected < 4 && row < 5 {
		if b.positions[column][row] == b.positions[column][row+1] {
			b.connected++
			b.checkVertical(column, row+1)
		}
	}
}

func (b *Board) checkDiagonal(column int, row int) {
	b.checkLR(column, row)
	if b.connected < 4 {
		b.connected = 0
		b.checkRL(column, row)
	}
}

func (b *Board) checkLR(column int, row int) {
	if (column > 0 && column < 6) && (row > 0 && row < 5) && b.connected < 4 {
		if b.positions[column][row] == b.positions[column-1][row-1] {
			b.connected++
			b.checkDiagonal(column-1, row-1)
		}
		if b.positions[column][row] == b.positions[column+1][row+1] {
			b.connected++
			b.checkDiagonal(column-1, row+1)
		}
	}
}

func (b *Board) checkRL(column int, row int) {
	if (column > 0 && column < 6) && (row > 0 && row < 5) && b.connected < 4 {
		if b.positions[column][row] == b.positions[column-1][row+1] {
			b.connected++
			b.checkDiagonal(column-1, row+1)
		}
		if b.positions[column][row] == b.positions[column+1][row-1] {
			b.connected++
			b.checkDiagonal(column+1, row-1)
		}
	}
}
