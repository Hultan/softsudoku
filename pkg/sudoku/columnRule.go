package sudoku

import (
	"errors"
	"fmt"
)

type ColumnRule struct {
	Column int
}

// NewColumnRule : Create a new ColumnRule
func NewColumnRule(col int) (*ColumnRule, error) {
	if col<1 || col>9 {
		return nil, errors.New(fmt.Sprintf("Invalid col : %v", col))
	}

	rule := new(ColumnRule)
	rule.Column = col

	return rule, nil
}

// Check : Check if column n rule is broken
func (r *ColumnRule) Check(sudoku *Sudoku) (bool, error) {
	result := make(map[int]bool)

	for row:=1;row<=9;row++ {
		value, err := sudoku.GetValue(row, r.Column)
		if err!=nil {
			return false, err
		}

		// Continue if the cell does not have a value
		if value == 0 {
			continue
		}

		// Check if the value is duplicate in the column
		if result[value] == true {
			return false, nil
		}

		result[value] = true
	}

	return true, nil
}