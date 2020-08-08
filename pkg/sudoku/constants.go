package sudoku

type Corner int

//goland:noinspection GoUnusedConst
const (
	UpperLeftCorner  Corner = iota
	UpperRightCorner
	LowerRightCorner
	LowerLeftCorner
)

func (c Corner) String() string {
	return [...]string{"UpperLeft","UpperRight","LowerRight","LowerLeft"}[c]
}

