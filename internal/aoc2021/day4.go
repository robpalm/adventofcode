package aoc2021

import (
	"bufio"
	"robpalm/adventofcode/internal/utils"
	"strings"
)

type Board struct {
	Squares  [5][5]int
	Finished bool
}

func Day4_1(scanner *bufio.Scanner) int {
	return day4(scanner, true)
}

func Day4_2(scanner *bufio.Scanner) int {
	return day4(scanner, false)
}

func day4(scanner *bufio.Scanner, first bool) int {
	var draws []int
	var boards []Board

	scanner.Scan()
	for _, draw := range strings.Split(scanner.Text(), ",") {
		draws = append(draws, utils.ToInt(draw))
	}

	for scanner.Scan() {
		boards = append(boards, setupBoard(scanner))
	}

	boardSum := 0
	for _, draw := range draws {
		for boardIndex := range boards {
			board := &boards[boardIndex]
			if board.Finished {
				continue
			}
			mark(board, draw)

			if won(board) {
				boardSum = sum(board) * draw
				if first {
					return boardSum
				}
			}
		}
	}
	return boardSum
}

func setupBoard(scanner *bufio.Scanner) (b Board) {
	b.Finished = false
	for y := 0; y < 5; y++ {
		scanner.Scan()
		row := strings.ReplaceAll(strings.TrimSpace(scanner.Text()), "  ", " ")
		for x, nr := range strings.Split(row, " ") {
			b.Squares[y][x] = utils.ToInt(nr)
		}
	}
	return
}

func mark(board *Board, draw int) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if board.Squares[y][x] == draw {
				board.Squares[y][x] = 0
				return
			}
		}
	}
}

func won(board *Board) bool {
	for a := 0; a < 5; a++ {
		var winX, winY, winD1, winD2 = true, true, true, true
		for b := 0; b < 5; b++ {
			if board.Squares[a][b] != 0 {
				winX = false
			}
			if board.Squares[b][a] != 0 {
				winY = false
			}
			if board.Squares[b][b] != 0 {
				winD1 = false
			}
			if board.Squares[b][4-b] != 0 {
				winD2 = false
			}
		}
		if winX || winY || winD1 || winD2 {
			board.Finished = true
			return true
		}
	}
	return false
}

func sum(board *Board) (sum int) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			sum += board.Squares[y][x]
		}
	}
	return
}
