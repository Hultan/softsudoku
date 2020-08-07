package sudoku

import (
	softSudoku "github.com/hultan/softsudoku/pkg/sudoku"
	"testing"
)

// TestRowRulesFull : Add and check all rows in a test sudoku game
func TestRowRulesFull(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	setTestSudoku(sudoku)
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
	if got==false {
		t.Errorf("Board should not violate the row rule!")
	}
}

// TestColumnRulesFull : Add and check all columns in a test sudoku game
func TestColumnRulesFull(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	setTestSudoku(sudoku)
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
	if got==false {
		t.Errorf("Board should not violate the column rule!")
	}
}

// TestBoxRulesFull : Add and check all boxes in a test sudoku game
func TestBoxRulesFull(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	setTestSudoku(sudoku)
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
	if got==false {
		t.Errorf("Board should not violate the box rule!")
	}
}
