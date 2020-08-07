package sudoku

import (
	softSudoku "github.com/hultan/softsudoku/pkg/sudoku"
	"testing"
)

// TestAddStandardRulesEmpty : Test add all standard rules to an empty sudoku game
func TestAddStandardRulesEmpty(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	for row:=1;row<=9;row++ {
		rule, err :=  softSudoku.NewRowRule(row)
		if err != nil {
			t.Error(err)
		}

		sudoku.AddRule(rule)
	}
	for col:=1;col<=9;col++ {
		rule, err := softSudoku.NewColumnRule(col)
		if err != nil {
			t.Error(err)
		}

		sudoku.AddRule(rule)
	}
	for box:=1;box<=9;box++ {
		rule, err := softSudoku.NewBoxRule(box)
		if err != nil {
			t.Error(err)
		}

		sudoku.AddRule(rule)
	}
}

// TestRowRulesEmpty : Add and check all rows in an empty sudoku game
func TestRowRulesEmpty(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
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
		t.Errorf("Empty board should not violate the row rule!")
	}
}

// TestColumnRulesEmpty : Add and check all columns in an empty sudoku game
func TestColumnRulesEmpty(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
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
		t.Errorf("Empty board should not violate the column rule!")
	}
}

// TestBoxRulesEmpty : Add and check all boxes in an empty sudoku game
func TestBoxRulesEmpty(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
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
		t.Errorf("Empty board should not violate the box rule!")
	}
}
