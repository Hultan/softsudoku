package sudoku

import (
	softSudoku "github.com/hultan/softsudoku/pkg/sudoku"
	"testing"
)

// TestRowColFail : Test to access all row/columns in an empty game
func TestNewBoard(t *testing.T) {
	board := softSudoku.NewSudoku()
	for r := 1; r <= 9; r++ {
		for c := 1; c <= 9; c++ {
			got, err := board.GetValue(r, c)
			if err != nil {
				t.Error(err)
			}
			if got != 0 {
				t.Errorf("Row %d:Col %d is not 0!", r, c)
			}
		}
	}
}

// TestGetValueFail : Test to access an invalid row/column (GetValue)
func TestGetValueFail(t *testing.T) {
	board := softSudoku.NewSudoku()
	got, err := board.GetValue(1, 10)
	if err != nil && got == 0 {
		// Expected result
		return
	}
	t.Error("Test should have failed!")
}

// TestSetValueFail : Test to set an invalid row/column/value (SetValue)
func TestSetValueFail(t *testing.T) {
	board := softSudoku.NewSudoku()
	err := board.SetValue(1, 8, 11)
	if err != nil {
		// Expected error
		return
	}
	t.Error("Test should have failed!")
}

// TestSetValue : Test to set an invalid row/column/value (GetValue/SetValue)
func TestSetValue(t *testing.T) {
	board := softSudoku.NewSudoku()
	for r := 1; r <= 9; r++ {
		for c := 1; c <= 9; c++ {
			for v := 1; v <= 9; v++ {
				err:=board.SetValue(r, c, v)
				if err!=nil {
					t.Error(err)
				}
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

// TestGiven : Test to get and set a given cell value
func TestGiven(t *testing.T) {
	board := softSudoku.NewSudoku()
	for r := 1; r <= 9; r++ {
		for c := 1; c <= 9; c++ {
			for v := 1; v <= 9; v++ {
				err:=board.SetGiven(r, c, v)
				if err!=nil {
					t.Error(err)
				}
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

// TestSetGivenFail : Test to set an invalid given cell value
func TestSetGivenFail(t *testing.T) {
	board := softSudoku.NewSudoku()
	// Invalid value : 0
	err:=board.SetGiven(5, 5, 0)
	if err==nil {
		t.Errorf("SetGiven(5,5,0) should fail!")
	}
	// Invalid value : 10
	err=board.SetGiven(5, 5, 10)
	if err==nil {
		t.Errorf("SetGiven(5,5,10) should fail!")
	}
	// Invalid row : 0
	err=board.SetGiven(0, 5, 5)
	if err==nil {
		t.Errorf("SetGiven(0,5,5) should fail!")
	}
	// Invalid row : 10
	err=board.SetGiven(10, 5, 5)
	if err==nil {
		t.Errorf("SetGiven(10,5,5) should fail!")
	}
	// Invalid column : 0
	err=board.SetGiven(5, 0, 5)
	if err==nil {
		t.Errorf("SetGiven(5,0,5) should fail!")
	}
	// Invalid column : 10
	err=board.SetGiven(5, 10, 5)
	if err==nil {
		t.Errorf("SetGiven(5,10,5) should fail!")
	}
}

// TestBoxNotations : Test to get and set box notations
func TestBoxNotations(t *testing.T) {
	board := softSudoku.NewSudoku()
	for r := 1; r <= 9; r++ {
		for c := 1; c <= 9; c++ {
			for v := 1; v <= 9; v++ {
				err:=board.ToggleBoxNotation(r, c, v)
				if err!=nil {
					t.Error(err)
				}
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

// TestCellNotations : Test to get and set cell notations
func TestCellNotations(t *testing.T) {
	board := softSudoku.NewSudoku()
	for r := 1; r <= 9; r++ {
		for c := 1; c <= 9; c++ {
			for v := 1; v <= 9; v++ {
				err:=board.ToggleCellNotation(r, c, v)
				if err!=nil {
					t.Error(err)
				}
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

// TestIsNotFull : Test IsFull function on an empty game
func TestIsNotFull(t *testing.T) {
	board := softSudoku.NewSudoku()
	got, err := board.IsFull()
	if err != nil {
		t.Error(err)
	}
	if got == true {
		t.Errorf("Empty board is complete!")
	}
}

// TestIsFull : Test IsFull function on a completed game
func TestIsFull(t *testing.T) {
	board := softSudoku.NewSudoku()
	setTestSudoku(board)

	got, err := board.IsFull()
	if err != nil {
		t.Error(err)
	}
	if got == false {
		t.Errorf("Full board is not complete!")
	}
}
