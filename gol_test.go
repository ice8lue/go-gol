package gogol

import (
	"testing"
)

func TestBoardGeneration(t *testing.T) {
	boardSizes := [4]int{10, 42, 100, 200}
	randomize := [2]bool{false, true}

	for _, rand := range randomize {
		for _, n := range boardSizes {
			board := GenerateBoard(n, n, rand)
			count := 0

			if len(board) != n {
				t.Error("Incorrect board height.")
			}

			for i := range board {
				if len(board[i]) != n {
					t.Error("Incorrect board width.")
				}

				for j := range board[i] {
					if board[i][j] == true {
						count++
					}
				}
			}

			if !rand && count > 0 {
				t.Error("Board should be empty.")
			}

			if rand && count == 0 {
				t.Error("Board should contain living cells.")
			}
		}
	}
}

func TestLivingNeighborCounting(t *testing.T) {
	board := [][]bool{
		{true, false, false},
		{true, false, true},
		{false, false, false},
	}

	positions := [][]int{
		{0, 0, 1},
		{1, 0, 3},
		{2, 0, 1},
		{0, 1, 1},
		{1, 1, 3},
		{2, 1, 0},
		{0, 2, 1},
		{1, 2, 2},
		{2, 2, 1},
	}

	for _, p := range positions {
		count := CountLivingNeighborCells(board, p[0], p[1])

		if count != p[2] {
			t.Errorf("Wrong number of living neighbors (counted %v, should be %v).", count, p[2])
		}
	}
}

func TestCalculateEvolution(t *testing.T) {
	board := [][]bool{
		{true, false, false, false, true},
		{false, true, true, false, false},
		{true, false, true, false, true},
		{false, false, false, true, false},
		{false, true, false, false, false},
	}

	nextGenBoard := CalculateEvolution(board)

	if nextGenBoard[4][0] != false {
		t.Error("Lonely cells should die.")
	}

	if nextGenBoard[1][1] != false {
		t.Error("Crowded cells should die.")
	}

	if nextGenBoard[1][4] != true {
		t.Error("Cell with three neighbors should become alive.")
	}

	if nextGenBoard[0][2] != true {
		t.Error("Cell with two neighbors should survive.")
	}
}

func TestCollectStatistics(t *testing.T) {
	board := [][]bool{
		{true, false, false, false, true},
		{false, true, true, false, false},
		{true, false, true, false, true},
		{false, false, false, true, false},
		{false, true, false, false, false},
	}

	stats := CollectStatistics(board)

	if stats.Alive != 9 {
		t.Error("Incorrect count of living cells.")
	}

	if stats.Dead != 16 {
		t.Error("Incorrect count of living cells.")
	}

	if stats.Total != 25 {
		t.Error("Incorrect count of total cells.")
	}
}
