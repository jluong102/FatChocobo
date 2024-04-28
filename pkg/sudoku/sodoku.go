package sudoku

import (
	"encoding/json"
	"net/http"
	"io"
)

const SUDOKU_URL string = "https://sudoku-api.vercel.app/api/dosuku"

type Board struct {
	Newboard struct {
		Grids []struct {
			Value      [][]int `json:"value"`
			Solution   [][]int `json:"solution"`
			Difficulty string  `json:"difficulty"`
		} `json:"grids"`
		Results int    `json:"results"`
		Message string `json:"message"`
	} `json:"newboard"`
}

func GetBoard() (*http.Response, error) {
	request := http.NewRequest(http.MethodGet, SUDOKU_URL, nil)

	client := http.Client{}
	return client.Do(request)
}

func ParseBoardResponse(request *http.Response) (*board, error) {
	board := new(Board)

	raw, err := io.ReadAll(request.Body)
	defer request.Body.Close()	

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw, board)

	return board, err	
}
