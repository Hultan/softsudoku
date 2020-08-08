package sudoku

import (
	softSudoku "github.com/hultan/softsudoku/pkg/sudoku"
	"testing"
)

// TestDiagonalRulesEmpty : Add and check all rows in an empty sudoku game
func TestDiagonalRulesEmpty(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	rule, err := softSudoku.NewDiagonalRule(softSudoku.UpperLeftCorner)
	if err != nil {
		t.Error(err)
	}
	sudoku.AddRule(rule)
	rule, err = softSudoku.NewDiagonalRule(softSudoku.LowerRightCorner)
	if err != nil {
		t.Error(err)
	}
	sudoku.AddRule(rule)

	got, err := sudoku.CheckRules()
	if err!=nil {
		t.Error(err)
	}
	if got==false {
		t.Errorf("Empty board should not violate the row rule!")
	}
}

// TestDiagonalRule : Add and check diagonals in a test sudoku game
func TestDiagonalRule(t *testing.T) {
	// Set up sudoku
	sudoku := softSudoku.NewSudoku()
	err:=setDiagonalTestSudoku(sudoku)
	if err!=nil {
		t.Error(err)
	}

	// Add two diagonal rules
	rule, err := softSudoku.NewDiagonalRule(softSudoku.UpperLeftCorner)
	if err != nil {
		t.Error(err)
	}
	sudoku.AddRule(rule)
	rule, err = softSudoku.NewDiagonalRule(softSudoku.UpperRightCorner)
	if err != nil {
		t.Error(err)
	}
	sudoku.AddRule(rule)

	// Check rules
	got, err := sudoku.CheckRules()
	if err!=nil {
		t.Error(err)
	}
	if got==false {
		t.Errorf("Sudoku should not violate the diagonal rule!")
	}
}


// TestDiagonalRuleFail : Add and check diagonals in an invalid test sudoku game
func TestDiagonalRuleFail(t *testing.T) {
	// Set up sudoku
	sudoku := softSudoku.NewSudoku()
	err:=setInvalidDiagonalTestSudoku(sudoku)
	if err!=nil {
		t.Error(err)
	}

	// Add two diagonal rules
	rule, err := softSudoku.NewDiagonalRule(softSudoku.UpperLeftCorner)
	if err != nil {
		t.Error(err)
	}
	sudoku.AddRule(rule)
	rule, err = softSudoku.NewDiagonalRule(softSudoku.UpperRightCorner)
	if err != nil {
		t.Error(err)
	}
	sudoku.AddRule(rule)

	// Check rules
	got, err := sudoku.CheckRules()
	if err!=nil {
		t.Error(err)
	}
	if got==true {
		t.Errorf("Invalid sudoku should violate the diagonal rule!")
	}
}
