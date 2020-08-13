package sudoku

import (
	"errors"
	"fmt"
)

type RowRule struct {
	row int
}

// NewRowRule : Create a new RowRule
func NewRowRule(row int) (*RowRule, error) {
	if row<1 || row>9 {
		return nil, errors.New(fmt.Sprintf("invalid row : %v", row))
	}

	rule := new(RowRule)
	rule.row = row

	return rule, nil
}

func (r *RowRule) GetRow() int {
	return r.row
}

// Check : Check if row n rule is broken
func (r *RowRule) Check(sudoku *Sudoku) (bool, error) {
	result := make(map[int]bool)

	for col:=1;col<=9;col++ {
		value, err := sudoku.GetValue(r.row, col)
		if err!=nil {
			return false, err
		}

		// Continue if the cell does not have a value
		if value == 0 {
			continue
		}

		// Check if the value is duplicate in the row
		if result[value] == true {
			return false, nil
		}

		result[value] = true
	}

	return true, nil
}