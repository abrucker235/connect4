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
	if b.connected == 3 {
		return true, nil
	}
	b.connected = 0

	b.CheckVertical(row, column)
	if b.connected == 3 {
		return true, nil
	}
	b.connected = 0

	b.CheckDiagonal(row, column)
	if b.connected == 3 {
		return true, nil
	}
	b.connected = 0

	return false, nil
}

// CheckHorizontal control function
func (b *Board) CheckHorizontal(row int, column int) {
	b.checkLeft(row, column)
	if b.connected == 3 {
		return
	}

	b.checkRight(row, column)
}

// Recursive function to check positions to the left
func (b *Board) checkLeft(row int, column int) {
	if b.connected < 3 && column > 0 {
		if b.positions[row][column] == b.positions[row][column-1] {
			b.connected++
			b.checkLeft(row, column-1)
		}
	}
}

// Recursive function to check positions to the right
func (b *Board) checkRight(row int, column int) {
	if b.connected < 3 && column < 6 {
		if b.positions[row][column] == b.positions[row][column+1] {
			b.connected++
			b.checkRight(row, column+1)
		}
	}
}

// CheckVertical control function
func (b *Board) CheckVertical(row int, column int) {
	if row >= 3 {
		for i := row; i > 0; i-- {
			if b.connected <= 3 {
				if b.positions[row][column] == b.positions[row-1][column] {
					b.connected++
				}
			}
		}
	}
}

// CheckDiagonal control function
func (b *Board) CheckDiagonal(row int, column int) {
	b.checkDiagonalDownLeft(row, column)
	b.checkDiagonalUpRight(row, column)
	if b.connected != 3 {
		b.connected = 0
	}

	b.checkDiagonalDownRight(row, column)
	b.checkDiagonalUpLeft(row, column)
	if b.connected != 3 {
		b.connected = 0
	}
}

func (b *Board) checkDiagonalDownLeft(row int, column int) {
	if column > 0 && row > 0 && b.connected < 3 {
		if b.positions[row][column] == b.positions[row-1][column-1] {
			b.connected++
			b.checkDiagonalDownLeft(row-1, column-1)
		}
	}
}

func (b *Board) checkDiagonalUpRight(row int, column int) {
	if column < 6 && row < 5 && b.connected < 3 {
		if b.positions[row][column] == b.positions[row+1][column+1] {
			b.connected++
			b.checkDiagonalUpRight(row+1, column+1)
		}
	}
}

func (b *Board) checkDiagonalUpLeft(row int, column int) {
	if column > 0 && row < 5 && b.connected < 3 {
		if b.positions[row][column] == b.positions[row+1][column-1] {
			b.connected++
			b.checkDiagonalUpLeft(row+1, column-1)
		}
	}
}

func (b *Board) checkDiagonalDownRight(row int, column int) {
	if column < 6 && row > 0 && b.connected < 3 {
		if b.positions[row][column] == b.positions[row-1][column+1] {
			b.connected++
			b.checkDiagonalDownRight(row-1, column+1)
		}
	}
}
