package main

import (
	"fmt"
)

func main() {
	//sudoku := sudoku.NewSudoku()
	//spew.Dump(sudoku)

	for box:=0; box<9;box ++{
		startRow := box/3 * 3
		startCol := box%3 * 3
		fmt.Printf("Box : %d ; {%d,%d}\n", box, startRow, startCol)
	}

}
