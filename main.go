package main

import (
	"fmt"
	sudoku "github.com/hultan/softsudoku/sudoku"
)

func main() {
	//sudoku := sudoku.NewSudoku()
	//spew.Dump(sudoku)

	TestBoxNotations()
	TestAddRowRules()

	for box:=0; box<9;box ++{
		startRow := box/3 * 3
		startCol := box%3 * 3
		fmt.Printf("Box : %d ; {%d,%d}\n", box, startRow, startCol)
	}

}

func TestAddRowRules() {
	board := sudoku.NewSudoku()
	for row:=1;row<=9;row++ {
		rule, err := sudoku.NewRowRule(row)
		if err != nil {
			fmt.Errorf("%e", err)
		}

		board.AddRule(rule)
	}

	got, err := board.CheckRules()
	if err!=nil {
		fmt.Errorf("%e", err)
	}
	if got==false {
		fmt.Errorf("empty board should not violate a rule")
	}
}

func TestBoxNotations() {
	board := sudoku.NewSudoku()
	//spew.Dump(board)
	for r := 1; r <= 9; r++ {
		for c := 1; c <= 9; c++ {
			for v := 1; v <= 9; v++ {
				board.ToggleBoxNotation(r, c, v)
			}
			got, err := board.GetBoxNotation(r, c)
			if err!=nil {
				fmt.Errorf("%e", err)
			}
			for v := 1; v <= 9; v++ {
				if got[v]!=true {
					fmt.Errorf("row %d:Col %d should have value %d but have value %v", r, c, v, got)
				}
			}
		}
	}
}