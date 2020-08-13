package sudoku

import (
	"errors"
	"fmt"
)

type Cell struct {
	Row int
	Column int
}

type Sudoku struct {
	grid     [9][9]int
	given    [9][9]bool
	notation [9][9]*notation
	rules    []Rule
}

type notation struct {
	box  map[int]bool
	cell map[int]bool
}

// NewSudoku : Create a new Sudoku
func NewSudoku() *Sudoku {
	sudoku := new(Sudoku)
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			notation := new(notation)
			notation.box = make(map[int]bool)
			notation.cell = make(map[int]bool)
			sudoku.notation[r][c] = notation
		}
	}
	return sudoku
}

// IsFull : Is the sudoku full?
func (s *Sudoku) IsFull() (bool, error) {
	for r := 1; r <= 9; r++ {
		for c := 1; c <= 9; c++ {
			value, err := s.GetValue(r, c)
			if err != nil {
				return false, err
			}
			if value == 0 {
				return false, nil
			}
		}
	}
	return true, nil
}

// GetValue : Gets the value of the cell in the specified row and column
func (s *Sudoku) GetValue(row, col int) (int, error) {
	err := s.validateRowColumn(row, col)
	if err != nil {
		return 0, err
	}

	return s.grid[row-1][col-1], nil
}

// SetValue : Sets the value of the cell in the specified row and column to the specified value
func (s *Sudoku) SetValue(row, col, value int) error {
	err := s.validateRowColumnValue(row, col, value)
	if err != nil {
		return err
	}

	s.grid[row-1][col-1] = value
	return nil
}

// GetGiven : Gets if the specified row and column is a given value
func (s *Sudoku) GetGiven(row, col int) (bool, error) {
	err := s.validateRowColumn(row, col)
	if err != nil {
		return false, err
	}
	return s.given[row-1][col-1], nil
}

// SetGiven : Sets the specified row and column as a given value
func (s *Sudoku) SetGiven(row, col, value int) error {
	err := s.validateRowColumnValue(row, col, value)
	if err != nil {
		return err
	}

	s.grid[row-1][col-1] = value
	s.given[row-1][col-1] = true
	return nil
}

// GetBoxNotation : Get the box notations for the specified cell
func (s *Sudoku) GetBoxNotation(row, col int) (map[int]bool, error) {
	err := s.validateRowColumn(row, col)
	if err != nil {
		return nil, err
	}

	// Copy the notation map
	result := make(map[int]bool)
	for k, v := range s.notation[row-1][col-1].box {
		result[k] = v
	}

	return result, nil
}

// ToggleBoxNotation : Toggle the box notation for the specified cell for value
func (s *Sudoku) ToggleBoxNotation(row, col, value int) error {
	err := s.validateRowColumnValue(row, col, value)
	if err != nil {
		return err
	}

	// Toggle the notation value
	s.notation[row-1][col-1].box[value] = !s.notation[row-1][col-1].box[value]

	return nil
}

// GetCellNotation : Get the cell notations for the specified cell
func (s *Sudoku) GetCellNotation(row, col int) (map[int]bool, error) {
	err := s.validateRowColumn(row, col)
	if err != nil {
		return nil, err
	}

	// Copy the notation map
	result := make(map[int]bool)
	for k, v := range s.notation[row-1][col-1].cell {
		result[k] = v
	}

	return result, nil
}

// ToggleCellNotation : Toggle the cell notation for the specified cell for value
func (s *Sudoku) ToggleCellNotation(row, col, value int) error {
	err := s.validateRowColumnValue(row, col, value)
	if err != nil {
		return err
	}

	// Toggle the notation value
	s.notation[row-1][col-1].cell[value] = !s.notation[row-1][col-1].cell[value]

	return nil
}

// SetSudoku : Fills the sudoku grid with the values provided
func (s *Sudoku) SetSudoku(grid [9][9]int) error {
	for r := 1; r <= 9; r++ {
		for c := 1; c <= 9; c++ {
			err:=s.SetValue(r,c,grid[r-1][c-1])
			if err!=nil {
				return err
			}
		}
	}

	return nil
}

// RULE FUNCTIONS

// AddRule : Add a rule to the sudoku rule set
func (s *Sudoku) AddRule(rule Rule) {
	s.rules = append(s.rules, rule)
}

// CheckRules : Check all rules
func (s *Sudoku) CheckRules() (bool, error) {
	result := true
	for _, value := range s.rules {
		ruleResult, err := value.Check(s)
		if err != nil {
			return false, err
		}
		result = result && ruleResult
	}
	return result, nil
}

//
// PRIVATE FUNCTIONS
//

// validateRowColumn : Validates the row and col parameters and returns an error
// 					   if they are not between 1 and 9
func (s *Sudoku) validateRowColumn(row, col int) error {
	if row < 1 || row > 9 {
		return errors.New(fmt.Sprintf("Invalid row : %d", row))
	}
	if col < 1 || col > 9 {
		return errors.New(fmt.Sprintf("Invalid column : %d", col))
	}
	return nil
}

// validateRowColumnValue : Validates the row, col and value parameters and returns an error
// 					   		if they are not between 1 and 9
func (s *Sudoku) validateRowColumnValue(row, col, value int) error {
	if row < 1 || row > 9 {
		return errors.New(fmt.Sprintf("Invalid row : %d", row))
	}
	if col < 1 || col > 9 {
		return errors.New(fmt.Sprintf("Invalid column : %d", col))
	}
	if value < 1 || value > 9 {
		return errors.New(fmt.Sprintf("Invalid value : %d", value))
	}

	return nil
}
