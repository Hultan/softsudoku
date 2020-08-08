package sudoku

import (
	softSudoku "github.com/hultan/softsudoku/pkg/sudoku"
	"testing"
)

//
// NewSudoku function
//

// TestNewBoard : Test a new sudoku game
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

//
// SetValue function
//

// TestGetSetValue : Test to set and get an row/column/value
func TestGetSetValue(t *testing.T) {
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

// TestGetValueFail : Test to get an invalid row/column
func TestGetValueFail(t *testing.T) {
	board := softSudoku.NewSudoku()
	_,err := board.GetValue(0, 5)
	if err == nil {
		t.Error("GetValue(0,5) should have failed!")
	}
	_,err = board.GetValue(10, 5)
	if err == nil {
		t.Error("GetValue(10,5) should have failed!")
	}
	_,err = board.GetValue(5, 0)
	if err == nil {
		t.Error("GetValue(5,0) should have failed!")
	}
	_,err = board.GetValue(5, 10)
	if err == nil {
		t.Error("GetValue(5,10) should have failed!")
	}
}

// TestSetValueFail : Test to set an invalid row/column/value
func TestSetValueFail(t *testing.T) {
	board := softSudoku.NewSudoku()
	err := board.SetValue(0, 5, 5)
	if err == nil {
		t.Error("SetValue(0,5,5) should have failed!")
	}
	err = board.SetValue(10, 5, 5)
	if err == nil {
		t.Error("SetValue(10,5,5) should have failed!")
	}
	err = board.SetValue(5, 0, 5)
	if err == nil {
		t.Error("SetValue(5,0,5) should have failed!")
	}
	err = board.SetValue(5, 10, 5)
	if err == nil {
		t.Error("SetValue(5,10,5) should have failed!")
	}
	err = board.SetValue(5, 5, 0)
	if err == nil {
		t.Error("SetValue(5,5,0) should have failed!")
	}
	err = board.SetValue(5, 5, 10)
	if err == nil {
		t.Error("SetValue(5,5,10) should have failed!")
	}
}

//
// GetGiven/SetGiven functions
//

// TestGetSetGiven : Test to get and set a given cell value
func TestGetSetGiven(t *testing.T) {
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

// TestGetGivenFail : Test to get a given cell value with invalid row/column
func TestGetGivenFail(t *testing.T) {
	board := softSudoku.NewSudoku()
	// Invalid row : 0
	_,err:=board.GetGiven(0, 5)
	if err==nil {
		t.Errorf("GetGiven(0,5) should fail!")
	}
	// Invalid row : 10
	_,err=board.GetGiven(10, 5)
	if err==nil {
		t.Errorf("GetGiven(10,5) should fail!")
	}
	// Invalid column : 0
	_,err=board.GetGiven(5, 0)
	if err==nil {
		t.Errorf("GetGiven(5,0) should fail!")
	}
	// Invalid column : 10
	_,err=board.GetGiven(5, 10)
	if err==nil {
		t.Errorf("GetGiven(5,10) should fail!")
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

//
// ToggleBoxNotation function
//

// TestBoxNotations : Test to toggle box notations
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

// TestBoxNotationsFail : Test to toggle box notations with invalid parameters
func TestBoxNotationsFail(t *testing.T) {
	board := softSudoku.NewSudoku()
	err:=board.ToggleBoxNotation(0, 1, 2)
	if err==nil {
		t.Errorf("Change box notations with invalid row/column should have failed!")
	}
	err=board.ToggleBoxNotation(10, 1, 2)
	if err==nil {
		t.Errorf("Change box notations with invalid row/column should have failed!")
	}
	err=board.ToggleBoxNotation(1, 0, 2)
	if err==nil {
		t.Errorf("Change box notations with invalid row/column should have failed!")
	}
	err=board.ToggleBoxNotation(1, 10, 2)
	if err==nil {
		t.Errorf("Change box notations with invalid row/column should have failed!")
	}
	err=board.ToggleBoxNotation(1, 2, 0)
	if err==nil {
		t.Errorf("Change box notations with invalid value should have failed!")
	}
	err=board.ToggleBoxNotation(1, 2, 10)
	if err==nil {
		t.Errorf("Change box notations with invalid value should have failed!")
	}
}

//
// ToggleCellNotation function
//

// TestCellNotations : Test to toggle cell notations
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

// TestCellNotationsFail : Test to toggle cell notations with invalid parameters
func TestCellNotationsFail(t *testing.T) {
	board := softSudoku.NewSudoku()
	err:=board.ToggleCellNotation(0, 1, 2)
	if err==nil {
		t.Errorf("Change cell notations with invalid row/column should have failed!")
	}
	err=board.ToggleCellNotation(10, 1, 2)
	if err==nil {
		t.Errorf("Change cell notations with invalid row/column should have failed!")
	}
	err=board.ToggleCellNotation(1, 0, 2)
	if err==nil {
		t.Errorf("Change cell notations with invalid row/column should have failed!")
	}
	err=board.ToggleCellNotation(1, 10, 2)
	if err==nil {
		t.Errorf("Change cell notations with invalid row/column should have failed!")
	}
	err=board.ToggleCellNotation(1, 2, 0)
	if err==nil {
		t.Errorf("Change cell notations with invalid value should have failed!")
	}
	err=board.ToggleCellNotation(1, 2, 10)
	if err==nil {
		t.Errorf("Change cell notations with invalid value should have failed!")
	}
}

//
// IsFull function
//

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
	err:=setTestSudoku(board)
	if err!=nil {
		t.Error(err)
	}

	got, err := board.IsFull()
	if err != nil {
		t.Error(err)
	}
	if got == false {
		t.Errorf("Full board is not complete!")
	}
}
