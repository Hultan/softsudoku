package sudoku

import (
	"errors"
	"fmt"
	"image/color"
)

type SumRule struct {
	sum int
	cells []Cell
	digitsCanRepeat bool
	backgroundColor color.Color
}

// NewSumRule : Create a new SumRule
func NewSumRule(sum int, cells []Cell, digitsCanRepeat bool, backgroundColor color.Color) (*SumRule, error) {
	rule := new(SumRule)
	if sum <= len(cells) {
		return nil, errors.New(fmt.Sprintf("sum too small for %d cells", len(cells)))
	}
	if sum >= len(cells)*9 {
		return nil, errors.New(fmt.Sprintf("sum too large for %d cells", len(cells)))
	}
	rule.sum = sum
	for index,cell := range cells {
		if cell.Row<1 {
			return nil, errors.New(fmt.Sprintf("row too small in cell %d", index))
		}
		if cell.Row>9 {
			return nil, errors.New(fmt.Sprintf("row too large in cell %d", index))
		}
		if cell.Column<1 {
			return nil, errors.New(fmt.Sprintf("column too small in cell %d", index))
		}
		if cell.Column>9 {
			return nil, errors.New(fmt.Sprintf("column too large in cell %d", index))
		}
	}
	rule.cells = cells
	rule.digitsCanRepeat = digitsCanRepeat
	rule.backgroundColor = backgroundColor

	return rule, nil
}

func (s *SumRule) GetCells() []Cell {
	return s.cells
}

func (s *SumRule) GetSum() int {
	return s.sum
}

func (s *SumRule) GetDigitsCanRepeat() bool {
	return s.digitsCanRepeat
}

func (s *SumRule) GetBackgroundColor() color.Color {
	return s.backgroundColor
}

func (s *SumRule) Check(sudoku *Sudoku) (bool, error) {
	var sum int
	var digit [9]bool

	for _,cell := range s.cells {
		value, err := sudoku.GetValue(cell.Row, cell.Column)
		if err!=nil {
			return false, err
		}
		sum += value

		if s.digitsCanRepeat == false {
			if digit[value] {
				return false, errors.New(fmt.Sprintf("digits may not repeat : %d", value))
			}
			digit[value] = true
		}
	}

	return sum == s.sum, nil
}