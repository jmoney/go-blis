package level3

import (
	"testing"

	"github.com/jmoney/go-blas/mat"
)

func Test_Dgemm(t *testing.T) {

	A := mat.General{Rows: 2, Cols: 2, RowsStride: 2, ColsStride: 1, Data: []float64{1.0, 2.0, 3.0, 4.0}}
	B := mat.General{Rows: 2, Cols: 2, RowsStride: 2, ColsStride: 1, Data: []float64{5.0, 6.0, 7.0, 8.0}}
	C := mat.General{Rows: 2, Cols: 2, RowsStride: 2, ColsStride: 1, Data: []float64{0.0, 0.0, 0.0, 0.0}}
	alpha := 1.0
	beta := 0.0
	Bli_dgemm(mat.NoTranspose, mat.NoTranspose, &alpha, A, &beta, B, C)
	C_Hat := mat.General{Rows: 2, Cols: 2, RowsStride: 2, ColsStride: 1, Data: []float64{5.0, 7.0, 25.0, 35.0}}

	for i := 0; i < 4; i++ {
		if C.Data[i] != C_Hat.Data[i] {
			t.Errorf("Test case 1 failed: %f != %f", C.Data[i], C_Hat.Data[i])
		}
	}
}
