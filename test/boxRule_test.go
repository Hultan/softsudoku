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
		errorCheck(err, t)
		if rule.GetBox()!=box {
			t.Errorf("Rule.GetBox() should have been %d", box)
		}
		sudoku.AddRule(rule)
	}

	got, err := sudoku.CheckRules()
	errorCheck(err, t)
	if got==false {
		t.Errorf("Empty board should not violate the box rule!")
	}
}

// TestBoxRules : Add and check all boxes in a test sudoku game
func TestBoxRules(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	err:=setTestSudoku(sudoku)
	errorCheck(err, t)

	for box:=1;box<=9;box++ {
		rule, err := softSudoku.NewBoxRule(box)
		errorCheck(err, t)

		sudoku.AddRule(rule)
	}

	got, err := sudoku.CheckRules()
	errorCheck(err, t)
	if got==false {
		t.Errorf("Board should not violate the box rule!")
	}
}

// TestBoxRulesFail : Add and check all boxes in a test sudoku game
func TestBoxRulesFail(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	err:=setInvalidTestSudoku(sudoku)
	errorCheck(err, t)
	for box:=1;box<=9;box++ {
		rule, err := softSudoku.NewBoxRule(box)
		errorCheck(err, t)

		sudoku.AddRule(rule)
	}

	got, err := sudoku.CheckRules()
	errorCheck(err, t)
	if got==true {
		t.Errorf("Board should violate the box rule!")
	}
}