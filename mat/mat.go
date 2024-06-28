package mat

type TransT uint32

const (
	Transpose TransT = iota
	NoTranspose
)

type DimT int64
type IncT int64

type General struct {
	Rows, Cols             DimT
	RowsStride, ColsStride IncT
	Data                   []float64
}
