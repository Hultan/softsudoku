package sudoku

import (
	softSudoku "github.com/hultan/softsudoku/pkg/sudoku"
	"image/color"
	"testing"
)

// TestSumRulesSimple : Test a simple single cell SumRule
func TestSumRulesSimple(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	cells := []softSudoku.Cell{{1,2}}
	background := color.RGBA{R: 255, A: 1}
	rule, err:=softSudoku.NewSumRule(8, cells, false, background)
	errorCheck(err, t)
	if rule.GetSum() != 8 {
		t.Errorf("Rule.GetSum() should have been 8")
	}
	if rule.GetDigitsCanRepeat() != false {
		t.Errorf("Rule.GetDigitsCanRepeat() should have been false")
	}
	if len(rule.GetCells()) != 1 {
		t.Errorf("Length of Rule.GetCells should have been 1")
	}
	if rule.GetCells()[0].Row != 1 {
		t.Errorf("Rule.GetCells()[0].Row should have been 1")
	}
	if rule.GetCells()[0].Column != 2 {
		t.Errorf("Rule.GetCells()[0].Column should have been 2")
	}
	if rule.GetBackgroundColor() != background {
		t.Errorf("Rule.GetBackgroundColor() not equal to background")
	}
	sudoku.AddRule(rule)

	err = sudoku.SetValue(1,2,8)
	errorCheck(err, t)
	result, err := sudoku.CheckRules()
	errorCheck(err, t)
	if result!=true {
		t.Errorf("Rule check should have returned true")
	}

	err = sudoku.SetValue(1,2,5)
	errorCheck(err, t)
	result, err = sudoku.CheckRules()
	errorCheck(err, t)
	if result!=false {
		t.Errorf("Rule check should have returned false")
	}
}

// TestSumRulesAdvanced : Test an advanced multi cell SumRule
func TestSumRulesAdvanced(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	cells := []softSudoku.Cell{{1,2},{3,4},{5,6}}
	background := color.RGBA{R: 255, A: 1}
	rule, err:=softSudoku.NewSumRule(12, cells, true, background)
	errorCheck(err, t)
	sudoku.AddRule(rule)

	err = sudoku.SetValue(1,2,8)
	errorCheck(err, t)
	err = sudoku.SetValue(3,4,1)
	errorCheck(err, t)
	err = sudoku.SetValue(5,6,3)
	errorCheck(err, t)
	result, err := sudoku.CheckRules()
	errorCheck(err, t)
	if result!=true {
		t.Errorf("Rule check should have returned true")
	}

	err = sudoku.SetValue(1,2,8)
	errorCheck(err, t)
	err = sudoku.SetValue(3,4,4)
	errorCheck(err, t)
	err = sudoku.SetValue(5,6,3)
	errorCheck(err, t)
	result, err = sudoku.CheckRules()
	errorCheck(err, t)
	if result!=false {
		t.Errorf("Rule check should have returned false")
	}
}

// TestInvalidSumRules : Test invalid sum rules
func TestInvalidSumRules(t *testing.T) {
	// Test sum too small
	cells := []softSudoku.Cell{{10,2},{10,4}}
	background := color.RGBA{R: 255, A: 1}
	_, err := softSudoku.NewSumRule(1, cells, false, background)
	if err==nil {
		t.Errorf("Invalid new rule should have failed (%e)", err)
	}
	// Test sum too large
	cells = []softSudoku.Cell{{10,2},{10,4}}
	background = color.RGBA{R: 255, A: 1}
	_, err = softSudoku.NewSumRule(19, cells, false, background)
	if err==nil {
		t.Errorf("Invalid new rule should have failed (%e)", err)
	}
	// Test row > 9
	cells = []softSudoku.Cell{{10,2},{10,4}}
	background = color.RGBA{R: 255, A: 1}
	_, err = softSudoku.NewSumRule(12, cells, false, background)
	if err==nil {
		t.Errorf("Invalid new rule should have failed (%e)", err)
	}
	// Test row < 1
	cells = []softSudoku.Cell{{-1,2},{1,4}}
	background = color.RGBA{R: 255, A: 1}
	_, err = softSudoku.NewSumRule(12, cells, false, background)
	if err==nil {
		t.Errorf("Invalid new rule should have failed (%e)", err)
	}
	// Test column > 9
	cells = []softSudoku.Cell{{1,15},{1,4}}
	background = color.RGBA{R: 255, A: 1}
	_, err = softSudoku.NewSumRule(12, cells, false, background)
	if err==nil {
		t.Errorf("Invalid new rule should have failed (%e)", err)
	}
	// Test column < 1
	cells = []softSudoku.Cell{{1,-2},{1,4}}
	background = color.RGBA{R: 255, A: 1}
	_, err = softSudoku.NewSumRule(12, cells, false, background)
	if err==nil {
		t.Errorf("Invalid new rule should have failed (%e)", err)
	}
}

// TestRepeatingDigits : Test repeating digits
func TestRepeatingDigits(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	cells := []softSudoku.Cell{{1,2},{1,3}}
	background := color.RGBA{R: 255, A: 1}
	rule, err:=softSudoku.NewSumRule(8, cells, false, background)
	errorCheck(err, t)
	sudoku.AddRule(rule)
	err = sudoku.SetValue(1,2,8)
	errorCheck(err, t)
	err = sudoku.SetValue(1,3,8)
	errorCheck(err, t)

	result, err := sudoku.CheckRules()
	if result==true {
		t.Errorf("Check rules should have failed (%e)", err)
	}
}