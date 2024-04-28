package main

import (
	"fmt"
	"log"
)

import (
	"github.com/jluong102/fatchocobo/pkg/sudoku"
)

func GetSudokuBoard() string {
	response, err := sudoku.GetBoard()

	if err != nil {
		log.Printf("Unable to get board\n\tError: %s", err)
		return "Unable to get board"
	}

	board, err := sudoku.ParseBoardResponse(response)

	if err != nil {
		log.Printf("Unable to parse board\n\tError: %s", err)
		return "Failed to parse board"
	}

	return parseBoard(board)
}

func parseBoard(board *sudoku.Board) string {
	msg := "```\n"
	msg += "+-+-+-+-+-+-+-+-+-+\n"

	for _, i := range board.NewBoard.Grids[0].Value {
		for _, j := range i {
			msg += fmt.Sprintf("|%s", j)
		}

		msg += "|"
	}

	msg += "+-+-+-+-+-+-+-+-+-+\n"
	msg += "```"

	return msg
}
