package sudoku

import (
	softSudoku "github.com/hultan/softsudoku/pkg/sudoku"
	"math/rand"
	"testing"
	"time"
)

// setTestSudoku : Fills the sudoku grid with a test sudoku
func setTestSudoku(sudoku *softSudoku.Sudoku) error {

	err:=sudoku.SetSudoku([9][9]int{
		{1,2,3,4,5,6,7,8,9},
		{4,5,6,7,8,9,1,2,3},
		{7,8,9,1,2,3,4,5,6},
		{2,3,4,5,6,7,8,9,1},
		{5,6,7,8,9,1,2,3,4},
		{8,9,1,2,3,4,5,6,7},
		{3,4,5,6,7,8,9,1,2},
		{6,7,8,9,1,2,3,4,5},
		{9,1,2,3,4,5,6,7,8},
	})
	if err!=nil {
		return err
	}
	return nil
}

// setInvalidTestSudoku : Fills the sudoku grid with a test sudoku
func setInvalidTestSudoku(sudoku *softSudoku.Sudoku) error {
	err:=setTestSudoku(sudoku)
	if err!=nil {
		return err
	}

	// Change a digit in the valid sudoku
	rand.Seed(time.Now().UnixNano())
	row:=rand.Intn(8)+1
	col:=rand.Intn(8)+1

	oldValue, _:=sudoku.GetValue(row,col)
	newValue:= oldValue
	for ; oldValue ==newValue; {
		newValue=rand.Intn(8) + 1
	}

	err=sudoku.SetValue(row,col,newValue)
	if err!=nil {
		return err
	}

	return nil
}

// setInvalidDiagonalTestSudoku : Fills the sudoku grid with a invalid test sudoku
func setInvalidDiagonalTestSudoku(sudoku *softSudoku.Sudoku) error {

	err:=sudoku.SetSudoku([9][9]int{
		{1,2,3,4,5,6,7,8,9},
		{4,5,6,7,8,9,1,2,3},
		{7,8,9,1,2,3,4,5,6},
		{2,3,4,5,6,7,8,9,1},
		{5,6,7,8,9,1,2,3,4},
		{8,9,1,2,3,4,5,6,7},
		{3,4,5,6,7,8,9,1,2},
		{6,7,8,9,1,2,3,1,5},
		{9,1,2,3,4,5,6,7,8},
	})
	if err!=nil {
		return err
	}
	return nil
}

// setDiagonalTestSudoku : Fills the sudoku grid with a invalid test sudoku
func setDiagonalTestSudoku(sudoku *softSudoku.Sudoku) error {

	err:=sudoku.SetSudoku([9][9]int{
		{1,2,3,4,5,6,7,8,9},
		{4,2,6,7,8,9,1,8,3},
		{7,8,3,1,2,3,7,5,6},
		{2,3,4,4,6,6,8,9,1},
		{5,6,7,8,5,1,2,3,4},
		{8,9,1,4,3,6,5,6,7},
		{3,4,3,6,7,8,7,1,2},
		{6,2,8,9,1,2,3,8,5},
		{1,1,2,3,4,5,6,7,9},
	})
	if err!=nil {
		return err
	}
	return nil
}

func errorCheck(err error, t *testing.T) {
	if err!=nil {
		t.Error(err)
	}
}