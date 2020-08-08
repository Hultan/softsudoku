package sudoku

import (
	softSudoku "github.com/hultan/softsudoku/pkg/sudoku"
	"testing"
)

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

// TestBoxRules : Add and check all boxes in a test sudoku game
func TestBoxRules(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	err:=setTestSudoku(sudoku)
	if err!=nil {
		t.Error(err)
	}

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

// TestBoxRulesFail : Add and check all boxes in a test sudoku game
func TestBoxRulesFail(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	err:=setInvalidTestSudoku(sudoku)
	if err!=nil {
		t.Error(err)
	}
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