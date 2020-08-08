package sudoku

import (
	softSudoku "github.com/hultan/softsudoku/pkg/sudoku"
	"testing"
)

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

// TestColumnRules : Add and check all columns in a test sudoku game
func TestColumnRules(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	err:=setTestSudoku(sudoku)
	if err!=nil {
		t.Error(err)
	}

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

// TestColumnRulesFail : Add and check all columns in a test sudoku game
func TestColumnRulesFail(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	err:=setInvalidTestSudoku(sudoku)
	if err!=nil {
		t.Error(err)
	}
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

