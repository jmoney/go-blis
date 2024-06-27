package mat

type TransT int

const (
	Transpose TransT = iota
	NoTranspose
)

type DimT int64
type IncT int64

type General struct {
	Rows, Cols             int
	RowsStride, ColsStride int
	Data                   []float64
}
