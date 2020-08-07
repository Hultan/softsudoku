package sudoku

import (
	softSudoku "github.com/hultan/softsudoku/pkg/sudoku"
	"testing"
)

// TestRowRulesFullFail : Add and check all rows in a test sudoku game
func TestRowRulesFullFail(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	setInvalidTestSudoku(sudoku)
	for row:=1;row<=9;row++ {
		rule, err := softSudoku.NewRowRule(row)
		if err != nil {
			t.Error(err)
		}

		sudoku.AddRule(rule)
	}

	got, err := sudoku.CheckRules()
	if err!=nil {
		t.Error(err)
	}
	if got==true {
		t.Errorf("Board should violate the row rule!")
	}
}

// TestColumnRulesFullFail : Add and check all columns in a test sudoku game
func TestColumnRulesFullFail(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	setInvalidTestSudoku(sudoku)
	for col:=1;col<=9;col++ {
		rule, err := softSudoku.NewColumnRule(col)
		if err != nil {
			t.Error(err)
		}

		sudoku.AddRule(rule)
	}

	got, err := sudoku.CheckRules()
	if err!=nil {
		t.Error(err)
	}
	if got==true {
		t.Errorf("Board should violate the column rule!")
	}
}

// TestBoxRulesFullFail : Add and check all boxes in a test sudoku game
func TestBoxRulesFullFail(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	setInvalidTestSudoku(sudoku)
	for box:=1;box<=9;box++ {
		rule, err := softSudoku.NewBoxRule(box)
		if err != nil {
			t.Error(err)
		}

		sudoku.AddRule(rule)
	}

	got, err := sudoku.CheckRules()
	if err!=nil {
		t.Error(err)
	}
	if got==true {
		t.Errorf("Board should violate the box rule!")
	}
}