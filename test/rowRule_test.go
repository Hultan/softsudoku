package sudoku

import (
	softSudoku "github.com/hultan/softsudoku/pkg/sudoku"
	"testing"
)

// TestRowRulesEmpty : Add and check all rows in an empty sudoku game
func TestRowRulesEmpty(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	for row:=1;row<=9;row++ {
		rule, err := softSudoku.NewRowRule(row)
		errorCheck(err, t)
		if rule.GetRow()!=row {
			t.Errorf("Rule.GetRow() should have been %d", row)
		}
		sudoku.AddRule(rule)
	}

	got, err := sudoku.CheckRules()
	errorCheck(err, t)
	if got==false {
		t.Errorf("Empty board should not violate the row rule!")
	}
}

// TestRowRules : Add and check all rows in a test sudoku game
func TestRowRules(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	err:=setTestSudoku(sudoku)
	errorCheck(err, t)
	for row:=1;row<=9;row++ {
		rule, err := softSudoku.NewRowRule(row)
		errorCheck(err, t)

		sudoku.AddRule(rule)
	}

	got, err := sudoku.CheckRules()
	errorCheck(err, t)
	if got==false {
		t.Errorf("Board should not violate the row rule!")
	}
}

// TestRowRulesFail : Add and check all rows in a test sudoku game
func TestRowRulesFail(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	err:=setInvalidTestSudoku(sudoku)
	errorCheck(err, t)
	for row:=1;row<=9;row++ {
		rule, err := softSudoku.NewRowRule(row)
		errorCheck(err, t)

		sudoku.AddRule(rule)
	}

	got, err := sudoku.CheckRules()
	errorCheck(err, t)
	if got==true {
		t.Errorf("Board should violate the row rule!")
	}
}
