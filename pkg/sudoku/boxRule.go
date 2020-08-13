package sudoku

import (
	"errors"
	"fmt"
)

type BoxRule struct {
	box int
}

// NewBoxRule : Create a new BoxRule
func NewBoxRule(box int) (*BoxRule, error) {
	if box < 1 || box > 9 {
		return nil, errors.New(fmt.Sprintf("invalid box : %v", box))
	}

	rule := new(BoxRule)
	rule.box = box

	return rule, nil
}

func (b *BoxRule) GetBox() int {
	return b.box
}

// Check : Check if box n rule is broken
func (r *BoxRule) Check(sudoku *Sudoku) (bool, error) {
	result := make(map[int]bool)
	startRow := (r.box - 1) / 3 * 3
	startCol := (r.box - 1) % 3 * 3

	for row := 1; row <= 3; row++ {
		for col := 1; col <= 3; col++ {
			value, err := sudoku.GetValue(startRow + row, startCol + col)
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
	}

	return true, nil
}
