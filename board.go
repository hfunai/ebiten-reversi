package main

import (
	"fmt"
	"math/rand"
)

type Board [][]int

func NewBoard() *Board {
	board := make(Board, numCellsY)
	for row := 0; row < numCellsY; row++ {
		board[row] = make([]int, numCellsX)
		for col := 0; col < numCellsX; col++ {
			board[row][col] = PlayerNone
			// Initial Position
			if (row == 3 && col == 3) || (row == 4 && col == 4) {
				board[row][col] = PlayerWhite
			} else if (row == 3 && col == 4) || (row == 4 && col == 3) {
				board[row][col] = PlayerBlack
			}
		}
	}
	return &board
}

func (b *Board) findValidMoves(currentPlayer int) Moves {
	validMoves := make(Moves, 0)
	oppponent := currentPlayer ^ 1
	shifts := Shifts{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}

	for r := 0; r < numCellsY; r++ {
		for c := 0; c < numCellsX; c++ {
			isValidMove := false
			// Pass already placed
			if (*b)[r][c] != PlayerNone {
				continue
			}
			// 上下左右に相手の石があるかチェック
			for _, s := range shifts {
				dr := r + s.Row
				dc := c + s.Col
				foundOpponent := false
				for dr >= 0 && dr < numCellsY && dc >= 0 && dc < numCellsX {
					// 隣接した相手の石があった上でその先に自分の石があれば置ける
					if (*b)[dr][dc] == oppponent {
						foundOpponent = true
					} else if (*b)[dr][dc] == currentPlayer && foundOpponent {
						isValidMove = true
						break
					} else {
						break
					}
					dr += s.Row
					dc += s.Col
				}
			}
			if isValidMove {
				m := &Move{Row: r, Col: c}
				validMoves = append(validMoves, m)
			}
		}
	}

	return validMoves
}

func (b *Board) getFlippableStones(currentPlayer, row, col int) Flippables {
	fbs := make(Flippables, 0)
	op := currentPlayer ^ 1
	ss := Shifts{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}

	for _, s := range ss {
		dr := row + s.Row
		dc := col + s.Col
		found := false
		tmpfbs := make(Flippables, 0)

		for dr >= 0 && dr < numCellsY && dc >= 0 && dc < numCellsX {
			if (*b)[dr][dc] == op {
				found = true
				fb := &Flippable{Row: dr, Col: dc}
				tmpfbs = append(tmpfbs, fb)
			} else if (*b)[dr][dc] == currentPlayer && found {
				fbs = append(fbs, tmpfbs...)
				break
			} else {
				break
			}
			dr += s.Row
			dc += s.Col
		}
	}

	return fbs
}

func (b *Board) flipStones(currentPlayer, row, col int) {
	fbs := b.getFlippableStones(currentPlayer, row, col)
	for _, fb := range fbs {
		(*b)[fb.Row][fb.Col] = currentPlayer
	}
}

func (b *Board) canPlace(currentPlayer, row, col int) bool {
	vms := b.findValidMoves(currentPlayer)
	for _, vm := range vms {
		if row == vm.Row && col == vm.Col {
			return true
		}
	}
	return false
}

func (b *Board) placeStone(currentPlayer, row, col int) {
	(*b)[row][col] = currentPlayer
}

func (b *Board) placeAuto(currentPlayer int) {
	vms := b.findValidMoves(currentPlayer)
	num := rand.Intn(len(vms))
	vm := vms[num]

	b.placeStone(currentPlayer, vm.Row, vm.Col)
	b.flipStones(currentPlayer, vm.Row, vm.Col)
}

func (b *Board) showHint(currentPlayer int) {
	vms := b.findValidMoves(currentPlayer)
	fmt.Printf("%v\n", b)
	for _, vm := range vms {
		fmt.Printf("player: %d, row: %d, col: %d\n", currentPlayer, vm.Row, vm.Col)
	}
}
