package sudoku

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	board := NewSudoku()
	for r := 1; r <= 9; r++ {
		for c := 1; c <= 9; c++ {
			got, err := board.GetValue(r, c)
			if err != nil {
				t.Error(err)
			}
			if got != 0 {
				t.Errorf("Row %d:Col %d is not 0 (%d)!", r, c, board.grid[r][c])
			}
		}
	}
}

func TestRowColFail(t *testing.T) {
	board := NewSudoku()
	got, err := board.GetValue(1, 10)
	if err != nil && got == 0 {
		// Expected result
		return
	}
	t.Error("Test should have failed!")
}

func TestRowColValueFail(t *testing.T) {
	board := NewSudoku()
	err := board.SetValue(1, 8, 11)
	if err != nil {
		// Expected error
		return
	}
	t.Error("Test should have failed!")
}

func TestSetCellValue(t *testing.T) {
	board := NewSudoku()
	for r := 1; r <= 9; r++ {
		for c := 1; c <= 9; c++ {
			for v := 1; v <= 9; v++ {
				board.SetValue(r, c, v)
				got, err := board.GetValue(r, c)
				if err != nil {
					t.Error(err)
				}
				if got != v {
					t.Errorf("Row %d:Col %d should have value %d but have value %d!", r, c, v, got)
				}
			}
		}
	}
}

func TestGiven(t *testing.T) {
	board := NewSudoku()
	for r := 1; r <= 9; r++ {
		for c := 1; c <= 9; c++ {
			for v := 1; v <= 9; v++ {
				board.SetGiven(r, c, v)
				got, err := board.GetGiven(r, c)
				if err != nil {
					t.Error(err)
				}
				if got != true {
					t.Errorf("Row %d:Col %d should have value %d but have value %v!", r, c, v, got)
				}
			}
		}
	}
}

func TestBoxNotations(t *testing.T) {
	board := NewSudoku()
	for r := 1; r <= 9; r++ {
		for c := 1; c <= 9; c++ {
			for v := 1; v <= 9; v++ {
				board.ToggleBoxNotation(r, c, v)
			}
			got, err := board.GetBoxNotation(r, c)
			if err != nil {
				t.Error(err)
			}
			for v := 1; v <= 9; v++ {
				if got[v] != true {
					t.Errorf("Row %d:Col %d should have box notation for %d!", r, c, v)
				}
			}
		}
	}
}

func TestCellNotations(t *testing.T) {
	board := NewSudoku()
	for r := 1; r <= 9; r++ {
		for c := 1; c <= 9; c++ {
			for v := 1; v <= 9; v++ {
				board.ToggleCellNotation(r, c, v)
			}
			got, err := board.GetCellNotation(r, c)
			if err != nil {
				t.Error(err)
			}
			for v := 1; v <= 9; v++ {
				if got[v] != true {
					t.Errorf("Row %d:Col %d should have cell notation for %d!", r, c, v)
				}
			}
		}
	}
}

func TestIsNotFull(t *testing.T) {
	board := NewSudoku()
	got, err := board.IsFull()
	if err != nil {
		t.Error(err)
	}
	if got == true {
		t.Errorf("Empty board is complete!")
	}
}

func TestIsFull(t *testing.T) {
	board := NewSudoku()
	for r := 1; r <= 9; r++ {
		for c := 1; c <= 9; c++ {
			board.SetValue(r, c, 1)
		}
	}

	got, err := board.IsFull()
	if err != nil {
		t.Error(err)
	}
	if got == false {
		t.Errorf("Full board is not complete!")
	}
}
