package sudoku

import (
	softSudoku "github.com/hultan/softsudoku/pkg/sudoku"
	"math/rand"
	"time"
)

// setTestSudoku : Fills the sudoku grid with a test sudoku
func setTestSudoku(sudoku *softSudoku.Sudoku) {

	sudoku.SetSudoku([9][9]int{
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
}
// setInvalidTestSudoku : Fills the sudoku grid with a test sudoku
func setInvalidTestSudoku(sudoku *softSudoku.Sudoku) {
	setTestSudoku(sudoku)

	// Change a digit in the valid sudoku
	rand.Seed(time.Now().UnixNano())
	row:=rand.Intn(8)+1
	col:=rand.Intn(8)+1

	oldValue, _:=sudoku.GetValue(row,col)
	newValue:= oldValue
	for ; oldValue ==newValue; {
		newValue=rand.Intn(8) + 1
	}
	err:=sudoku.SetValue(row,col,newValue)
	if err!=nil {
		// Ignore
	}
}
