package sudoku

import (
	softSudoku "github.com/hultan/softsudoku/pkg/sudoku"
	"testing"
)

// TestDiagonalRulesEmpty : Add and check all rows in an empty sudoku game
func TestDiagonalRulesEmpty(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	rule, err := softSudoku.NewDiagonalRule(softSudoku.UpperLeftCorner)
	errorCheck(err, t)
	if rule.GetStartingCorner()!=softSudoku.UpperLeftCorner {
		t.Errorf("Rule.GetStartingCorner() should have been UpperLeftCorner")
	}
	sudoku.AddRule(rule)
	rule, err = softSudoku.NewDiagonalRule(softSudoku.LowerRightCorner)
	errorCheck(err, t)
	if rule.GetStartingCorner()!=softSudoku.LowerRightCorner {
		t.Errorf("Rule.GetStartingCorner() should have been LowerRightCorner")
	}
	sudoku.AddRule(rule)

	got, err := sudoku.CheckRules()
	errorCheck(err, t)
	if got==false {
		t.Errorf("Empty board should not violate the row rule!")
	}
}

// TestDiagonalRule : Add and check diagonals in a test sudoku game
func TestDiagonalRule(t *testing.T) {
	// Set up sudoku
	sudoku := softSudoku.NewSudoku()
	err:=setDiagonalTestSudoku(sudoku)
	errorCheck(err, t)

	// Add two diagonal rules
	rule, err := softSudoku.NewDiagonalRule(softSudoku.UpperLeftCorner)
	errorCheck(err, t)
	sudoku.AddRule(rule)
	rule, err = softSudoku.NewDiagonalRule(softSudoku.UpperRightCorner)
	errorCheck(err, t)
	sudoku.AddRule(rule)

	// Check rules
	got, err := sudoku.CheckRules()
	errorCheck(err, t)
	if got==false {
		t.Errorf("Sudoku should not violate the diagonal rule!")
	}
}


// TestDiagonalRuleFail : Add and check diagonals in an invalid test sudoku game
func TestDiagonalRuleFail(t *testing.T) {
	// Set up sudoku
	sudoku := softSudoku.NewSudoku()
	err:=setInvalidDiagonalTestSudoku(sudoku)
	errorCheck(err, t)

	// Add two diagonal rules
	rule, err := softSudoku.NewDiagonalRule(softSudoku.UpperLeftCorner)
	errorCheck(err, t)
	sudoku.AddRule(rule)
	rule, err = softSudoku.NewDiagonalRule(softSudoku.UpperRightCorner)
	errorCheck(err, t)
	sudoku.AddRule(rule)

	// Check rules
	got, err := sudoku.CheckRules()
	errorCheck(err, t)
	if got==true {
		t.Errorf("Invalid sudoku should violate the diagonal rule!")
		
	}
}
