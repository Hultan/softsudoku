package sudoku

import (
	"errors"
	"fmt"
	"image/color"
)

type ThermoRule struct {
	cells           []Cell
	increasing      bool
	strict          bool
	backgroundColor color.Color
	thermoColor     color.Color
}

// NewThermoRule : Create a new ThermoRule
func NewThermoRule(cells []Cell, increasing, strict bool, backgroundColor color.Color, thermoColor color.Color) (*ThermoRule, error) {
	rule := new(ThermoRule)
	rule.increasing = increasing
	rule.strict = strict
	var lastCell *Cell
	for index, cell := range cells {
		if lastCell != nil {
			// Check that cells are consecutive
			if absInt(cell.Row-lastCell.Row) >= 1 || absInt(cell.Column-lastCell.Column) >= 1 {
				return nil, errors.New(fmt.Sprintf("Cells must be consecutive : cell %d", index+1))
			}
		}
		lastCell = &cell
	}
	rule.cells = cells
	rule.backgroundColor = backgroundColor
	rule.thermoColor = thermoColor

	return rule, nil
}

func (t *ThermoRule) GetCells() []Cell {
	return t.cells
}

func (t *ThermoRule) GetIncreasing() bool {
	return t.increasing
}

func (t *ThermoRule) GetStrict() bool {
	return t.strict
}

func (t *ThermoRule) GetBackgroundColor() color.Color {
	return t.backgroundColor
}

func (t *ThermoRule) GetThermoColor() color.Color {
	return t.thermoColor
}

func (t *ThermoRule) Check(sudoku *Sudoku) (bool, error) {
	var lastValue int = -1

	for _, cell := range t.cells {
		value, err := sudoku.GetValue(cell.Row, cell.Column)
		if lastValue > 0 {
			if err != nil {
				return false, err
			}
			dValue := value - lastValue
			if t.increasing && t.strict && dValue <= 0 {
				return false, nil
			} else if t.increasing && t.strict == false && dValue < 0 {
				return false, nil
			} else if t.increasing==false && t.strict && dValue >= 0 {
				return false, nil
			} else if t.increasing==false && t.strict == false && dValue > 0 {
				return false, nil
			}
		}
		lastValue = value
	}

	return true, nil
}

// absInt returns the absolute value of x.
func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
