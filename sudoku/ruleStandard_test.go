package sudoku

import "testing"

// TestRowRules : Add and check all rows in an empty sudoku game
func TestRowRules(t *testing.T) {
	sudoku := NewSudoku()
	SetTestSudoku(sudoku)
	for row:=1;row<=9;row++ {
		rule, err := NewRowRule(row)
		if err != nil {
			t.Error(err)
		}

		sudoku.AddRule(rule)
	}

	got, err := sudoku.CheckRules()
	if err!=nil {
		t.Error(err)
	}
	if got==false {
		t.Errorf("Board should not violate the row rule!")
	}
}

// TestColumnRules : Add and check all columns in an empty sudoku game
func TestColumnRules(t *testing.T) {
	sudoku := NewSudoku()
	SetTestSudoku(sudoku)
	for col:=1;col<=9;col++ {
		rule, err := NewColumnRule(col)
		if err != nil {
			t.Error(err)
		}

		sudoku.AddRule(rule)
	}

	got, err := sudoku.CheckRules()
	if err!=nil {
		t.Error(err)
	}
	if got==false {
		t.Errorf("Board should not violate the column rule!")
	}
}

// TestBoxRules : Add and check all boxes in an empty sudoku game
func TestBoxRules(t *testing.T) {
	sudoku := NewSudoku()
	SetTestSudoku(sudoku)
	for box:=1;box<=9;box++ {
		rule, err := NewBoxRule(box)
		if err != nil {
			t.Error(err)
		}

		sudoku.AddRule(rule)
	}

	got, err := sudoku.CheckRules()
	if err!=nil {
		t.Error(err)
	}
	if got==false {
		t.Errorf("Board should not violate the box rule!")
	}
}

// SetTestSudoku : Fills the sudoku grid with a test sudoku
func SetTestSudoku(sudoku *Sudoku) {

	sudoku.SetSudoku([9][9]int{
		{1,2,3,4,5,6,7,8,9},
		{4,5,6,7,8,9,1,2,3},
		{7,8,9,1,2,3,4,5,6},
		{2,3,4,5,6,7,8,9,1},
		{5,6,7,8,9,1,2,3,4},
		{8,9,1,2,3,4,5,6,7},
		{3,4,5,6,7,8,9,1,2},
		{6,7,8,9,1,2,3,4,5},
		{9,1,2,3,4,5,6,7,8},
	})
}