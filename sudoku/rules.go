package sudoku

type Rule interface {
	Check(sudoku *Sudoku) (bool, error)
}