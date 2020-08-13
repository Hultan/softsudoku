package sudoku

import (
	"errors"
	"fmt"
)

type DiagonalRule struct {
	startingCorner Corner
}

// NewDiagonalRule : Create a new DiagonalRule
func NewDiagonalRule(startingCorner Corner) (*DiagonalRule, error) {
	if startingCorner < 0 || startingCorner > 3 {
		return nil, errors.New(fmt.Sprintf("invalid starting corner (valid: 0-3) : %v", startingCorner))
	}

	rule := new(DiagonalRule)
	rule.startingCorner = startingCorner

	return rule, nil
}

func (d *DiagonalRule) GetStartingCorner() Corner {
 	return d.startingCorner
}

// Check : Check if column n rule is broken
func (r *DiagonalRule) Check(sudoku *Sudoku) (bool, error) {
	result := make(map[int]bool)
	var value int
	var err error

	for i := 1; i <= 9; i++ {
		if r.startingCorner == UpperLeftCorner || r.startingCorner == LowerRightCorner {
			value, err = sudoku.GetValue(i, i)
		} else {
			value, err = sudoku.GetValue(i, 10-i)
		}
		if err != nil {
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
