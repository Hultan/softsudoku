package sudoku

import (
	softSudoku "github.com/hultan/softsudoku/pkg/sudoku"
	"image/color"
	"testing"
)

// TestThermoRulesSimple : Test a simple ThermoRule
func TestThermoRulesSimple(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	cells := []softSudoku.Cell{{1,2}}
	background := color.RGBA{R: 255, A: 1}
	thermo := color.RGBA{G: 255, A: 1}
	rule, err:=softSudoku.NewThermoRule(cells, true, true,  background, thermo)
	errorCheck(err, t)
	if rule.GetIncreasing() != true {
		t.Errorf("Rule.GetIncreasin() should have been true")
	}
	if rule.GetStrict() != true {
		t.Errorf("Rule.GetStrict() should have been true")
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
	if rule.GetThermoColor() != thermo {
		t.Errorf("Rule.GetBackgroundColor() not equal to background")
	}
	sudoku.AddRule(rule)

	result, err := sudoku.CheckRules()
	errorCheck(err, t)
	if result!=true {
		t.Errorf("Rule check should have returned true")
	}
}

// TestThermoRulesIncreasingStrict : Test a ThermoRule that is strict increasing
func TestThermoRulesIncreasingStrict(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	cells := []softSudoku.Cell{{1,2},{2,2},{2,3}, {3,3}}
	background := color.RGBA{R: 255, A: 1}
	thermo := color.RGBA{G: 255, A: 1}
	rule, err:=softSudoku.NewThermoRule(cells, true, true,  background, thermo)
	errorCheck(err, t)

	sudoku.AddRule(rule)

	err = sudoku.SetValue(1,2,1)
	errorCheck(err, t)
	err = sudoku.SetValue(2,2,2)
	errorCheck(err, t)
	err = sudoku.SetValue(2,3,3)
	errorCheck(err, t)
	err = sudoku.SetValue(3,3,4)
	errorCheck(err, t)
	result, err := sudoku.CheckRules()
	errorCheck(err, t)
	if result!=true {
		t.Errorf("Rule check should have returned true")
	}
}

// TestThermoRulesIncreasing : Test a ThermoRule that is increasing
func TestThermoRulesIncreasing(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	cells := []softSudoku.Cell{{1,2},{2,2},{2,3}, {3,4}}
	background := color.RGBA{R: 255, A: 1}
	thermo := color.RGBA{G: 255, A: 1}
	rule, err:=softSudoku.NewThermoRule(cells, true, false,  background, thermo)
	errorCheck(err, t)

	sudoku.AddRule(rule)

	err = sudoku.SetValue(1,2,1)
	errorCheck(err, t)
	err = sudoku.SetValue(2,2,2)
	errorCheck(err, t)
	err = sudoku.SetValue(2,3,3)
	errorCheck(err, t)
	err = sudoku.SetValue(3,4,3)
	errorCheck(err, t)
	result, err := sudoku.CheckRules()
	errorCheck(err, t)
	if result!=true {
		t.Errorf("Rule check should have returned true")
	}
}

// TestThermoRulesDecreasingStrict : Test a ThermoRule that is strict decreasing
func TestThermoRulesDecreasingStrict(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	cells := []softSudoku.Cell{{1,2},{2,2},{2,3}, {3,3}}
	background := color.RGBA{R: 255, A: 1}
	thermo := color.RGBA{G: 255, A: 1}
	rule, err:=softSudoku.NewThermoRule(cells, false, true,  background, thermo)
	errorCheck(err, t)

	sudoku.AddRule(rule)

	err = sudoku.SetValue(1,2,4)
	errorCheck(err, t)
	err = sudoku.SetValue(2,2,3)
	errorCheck(err, t)
	err = sudoku.SetValue(2,3,2)
	errorCheck(err, t)
	err = sudoku.SetValue(3,3,1)
	errorCheck(err, t)
	result, err := sudoku.CheckRules()
	errorCheck(err, t)
	if result!=true {
		t.Errorf("Rule check should have returned true")
	}
}

// TestThermoRulesDecreasing : Test a ThermoRule that is decreasing
func TestThermoRulesDecreasing(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	cells := []softSudoku.Cell{{1,2},{2,2},{2,3}, {3,4}}
	background := color.RGBA{R: 255, A: 1}
	thermo := color.RGBA{G: 255, A: 1}
	rule, err:=softSudoku.NewThermoRule(cells, false, false,  background, thermo)
	errorCheck(err, t)

	sudoku.AddRule(rule)

	err = sudoku.SetValue(1,2,3)
	errorCheck(err, t)
	err = sudoku.SetValue(2,2,2)
	errorCheck(err, t)
	err = sudoku.SetValue(2,3,1)
	errorCheck(err, t)
	err = sudoku.SetValue(3,4,1)
	errorCheck(err, t)
	result, err := sudoku.CheckRules()
	errorCheck(err, t)
	if result!=true {
		t.Errorf("Rule check should have returned true")
	}
}

// TestThermoRulesIncreasingStrictFail : Test a ThermoRule that is strict increasing (should fail)
func TestThermoRulesIncreasingStrictFail(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	cells := []softSudoku.Cell{{1,2},{2,2},{2,3}, {3,3}}
	background := color.RGBA{R: 255, A: 1}
	thermo := color.RGBA{G: 255, A: 1}
	rule, err:=softSudoku.NewThermoRule(cells, true, true,  background, thermo)
	errorCheck(err, t)

	sudoku.AddRule(rule)

	err = sudoku.SetValue(1,2,1)
	errorCheck(err, t)
	err = sudoku.SetValue(2,2,2)
	errorCheck(err, t)
	err = sudoku.SetValue(2,3,2)
	errorCheck(err, t)
	err = sudoku.SetValue(3,3,3)
	errorCheck(err, t)
	result, err := sudoku.CheckRules()
	errorCheck(err, t)
	if result!=false {
		t.Errorf("Rule check should have returned false")
	}
}

// TestThermoRulesIncreasingFail : Test a ThermoRule that is increasing (should fail)
func TestThermoRulesIncreasingFail(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	cells := []softSudoku.Cell{{1,2},{2,2},{2,3}, {3,4}}
	background := color.RGBA{R: 255, A: 1}
	thermo := color.RGBA{G: 255, A: 1}
	rule, err:=softSudoku.NewThermoRule(cells, true, false,  background, thermo)
	errorCheck(err, t)

	sudoku.AddRule(rule)

	err = sudoku.SetValue(1,2,1)
	errorCheck(err, t)
	err = sudoku.SetValue(2,2,2)
	errorCheck(err, t)
	err = sudoku.SetValue(2,3,1)
	errorCheck(err, t)
	err = sudoku.SetValue(3,4,3)
	errorCheck(err, t)
	result, err := sudoku.CheckRules()
	errorCheck(err, t)
	if result!=false {
		t.Errorf("Rule check should have returned false")
	}
}

// TestThermoRulesDecreasingStrictFail : Test a ThermoRule that is strict decreasing (should fail)
func TestThermoRulesDecreasingStrictFail(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	cells := []softSudoku.Cell{{1,2},{2,2},{2,3}, {3,3}}
	background := color.RGBA{R: 255, A: 1}
	thermo := color.RGBA{G: 255, A: 1}
	rule, err:=softSudoku.NewThermoRule(cells, false, true,  background, thermo)
	errorCheck(err, t)

	sudoku.AddRule(rule)

	err = sudoku.SetValue(1,2,4)
	errorCheck(err, t)
	err = sudoku.SetValue(2,2,3)
	errorCheck(err, t)
	err = sudoku.SetValue(2,3,3)
	errorCheck(err, t)
	err = sudoku.SetValue(3,3,1)
	errorCheck(err, t)
	result, err := sudoku.CheckRules()
	errorCheck(err, t)
	if result!=false {
		t.Errorf("Rule check should have returned false")
	}
}

// TestThermoRulesDecreasingFail : Test a ThermoRule that is decreasing (should fail)
func TestThermoRulesDecreasingFail(t *testing.T) {
	sudoku := softSudoku.NewSudoku()
	cells := []softSudoku.Cell{{1,2},{2,2},{2,3}, {3,4}}
	background := color.RGBA{R: 255, A: 1}
	thermo := color.RGBA{G: 255, A: 1}
	rule, err:=softSudoku.NewThermoRule(cells, false, false,  background, thermo)
	errorCheck(err, t)

	sudoku.AddRule(rule)

	err = sudoku.SetValue(1,2,3)
	errorCheck(err, t)
	err = sudoku.SetValue(2,2,2)
	errorCheck(err, t)
	err = sudoku.SetValue(2,3,3)
	errorCheck(err, t)
	err = sudoku.SetValue(3,4,1)
	errorCheck(err, t)
	result, err := sudoku.CheckRules()
	errorCheck(err, t)
	if result!=false {
		t.Errorf("Rule check should have returned false")
	}
}
