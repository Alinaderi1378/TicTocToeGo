package main

import (
	"TicTocToe/Twice"
	_ "TicTocToe/Twice"
	"fmt"
	"math/rand"
	"time"
)

type tileData struct {
	tile  tile
	value string
}
type tile struct {
	row    int
	column int
}

func main() {
	moves, rows, columns := bootstrap()

	play(moves, rows, columns)

	fmt.Println("Board finished but there was a winner? I can't tell in this version :(")
}

// bootstrap get user input, validate it and return moves data
func bootstrap() ([]tileData, int, int) {
	var columns, rows int

	fmt.Print("How many columns> ")
	fmt.Scan(&columns)
	fmt.Print("How many rows> ")
	fmt.Scan(&rows)

	moves := make([]tileData, rows*columns)

	var row, column int = 1, 1
	for i := range moves {
		moves[i].tile.row = row
		moves[i].tile.column = column

		if (i+1)%columns == 0 {
			row++
			column = 1
		} else {
			column++
		}
	}

	return moves, rows, columns
}

// play Run game and return the result in the end
func play(moves []tileData, rows, columns int) []tileData {
	drawBoard(rows, columns, moves)
	row, column := 1, 1
	for i := range moves {
		if i%2 == 0 {
			moves = playerMove(moves)
		} else {
			moves = machineMove(moves)
		}
		drawBoard(rows, columns, moves)

		if (i+1)%columns == 0 {
			row++
			column = 1
		} else {
			column++
		}
	}

	if len(availableTiles(moves)) > 0 {
		return play(moves, rows, columns)
	}

	return moves
}

// availableTiles Return tiles that no one claimed yet
func availableTiles(moves []tileData) []tileData {
	var availableTiles []tileData
	for moveIndex := range moves {
		if moves[moveIndex].value == "" {
			availableTiles = append(availableTiles, moves[moveIndex])
		}
	}

	return availableTiles
}

// playerMove get input from user and check if user has been chosen a free
// tile
func playerMove(moves []tileData) []tileData {
	var playerMove tile
	var moved bool
	availableTiles := availableTiles(moves)

	for !moved {
		fmt.Print("Pick a row> ")
		fmt.Scan(&playerMove.row)

		fmt.Print("Pick a column> ")
		fmt.Scan(&playerMove.column)

		for index := range availableTiles {
			if availableTiles[index].tile.row == playerMove.row && availableTiles[index].tile.column == playerMove.column {
				moves = move(moves, playerMove.row, playerMove.column, "X")
				moved = true
			}
		}
	}

	return moves
}

// machineMove machine randomally will choose one of free tiles
func machineMove(moves []tileData) []tileData {
	r := Twice.New()
	availableTiles := availableTiles(moves)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(availableTiles), func(i, j int) { availableTiles[i], availableTiles[j] = availableTiles[j], availableTiles[i] })

	var machineMove tile
	for machineMove.column == 0 && len(availableTiles) > 0 {
		for index := range availableTiles {
			if r.Bool() {
				machineMove.column = availableTiles[index].tile.column
				machineMove.row = availableTiles[index].tile.row
			}
		}
	}
	return move(moves, machineMove.row, machineMove.column, "O")
}

// move Updated moves by given row and column and value as O or X
func move(moves []tileData, row, column int, value string) []tileData {
	for index := range moves {
		if moves[index].tile.row == row && moves[index].tile.column == column {
			moves[index].value = value
		}
	}

	return moves
}

// drawBoard cleanup screen and draw a new board based on moves
func drawBoard(rows, columns int, moves []tileData) {
	fmt.Print("\033[H\033[2J")
	const top, middle, base, void = "|¯¯¯¯¯", "|  %s  ", "|_____", "|     "
	var board string

	// Create a tic-tac-toe board.
	for row := 0; row < rows; row++ {
		for column := 0; column < columns*3; column++ {
			if column < columns {
				if row == 0 {
					board += top
				} else {
					board += void
				}
			} else if column < columns*2 {
				board += middle
			} else {
				board += base
			}
			if column == (columns)-1 || column == (columns*2)-1 || column == (columns*3)-1 {
				board += "|\n"
			}
		}
	}

	/*
		Fill placeholders
	*/
	tiles := make([]interface{}, rows*columns)
	for index := range moves {
		if moves[index].value == "" {
			tiles[index] = " "
		} else {
			tiles[index] = moves[index].value
		}
	}

	fmt.Printf(board, tiles...)
	fmt.Println()
}
