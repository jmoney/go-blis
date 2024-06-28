package level3

import (
	"testing"

	"github.com/jmoney/go-lapack/mat"
)

type TestCase struct {
	A, B, C, C_Hat mat.General
	A_Transpose    mat.TransT
	B_Transpose    mat.TransT
	alpha          float64
	beta           float64
	assertError    bool
}

func Test_Dgemm(t *testing.T) {
	tests := map[string]TestCase{
		"A*B=C": {
			A:           mat.General{Rows: 2, Cols: 2, RowsStride: 2, ColsStride: 1, Data: []float64{1.0, 2.0, 3.0, 4.0}},
			B:           mat.General{Rows: 2, Cols: 2, RowsStride: 2, ColsStride: 1, Data: []float64{5.0, 6.0, 7.0, 8.0}},
			C:           mat.General{Rows: 2, Cols: 2, RowsStride: 2, ColsStride: 1, Data: []float64{0.0, 0.0, 0.0, 0.0}},
			C_Hat:       mat.General{Rows: 2, Cols: 2, RowsStride: 2, ColsStride: 1, Data: []float64{5.0, 7.0, 25.0, 35.0}},
			A_Transpose: mat.NoTranspose,
			B_Transpose: mat.NoTranspose,
			alpha:       1.0,
			beta:        0.0,
			assertError: false,
		},
		"5A*B=C": {
			A:           mat.General{Rows: 2, Cols: 2, RowsStride: 2, ColsStride: 1, Data: []float64{1.0, 2.0, 3.0, 4.0}},
			B:           mat.General{Rows: 2, Cols: 2, RowsStride: 2, ColsStride: 1, Data: []float64{5.0, 6.0, 7.0, 8.0}},
			C:           mat.General{Rows: 2, Cols: 2, RowsStride: 2, ColsStride: 1, Data: []float64{0.0, 0.0, 0.0, 0.0}},
			C_Hat:       mat.General{Rows: 2, Cols: 2, RowsStride: 2, ColsStride: 1, Data: []float64{25.0, 35.0, 125.0, 175.0}},
			A_Transpose: mat.NoTranspose,
			B_Transpose: mat.NoTranspose,
			alpha:       5.0,
			beta:        0.0,
			assertError: false,
		},
		"A*B=C Dimensions do not match": {
			A:           mat.General{Rows: 2, Cols: 2, RowsStride: 2, ColsStride: 1, Data: []float64{1.0, 2.0, 3.0, 4.0}},
			B:           mat.General{Rows: 3, Cols: 1, RowsStride: 2, ColsStride: 1, Data: []float64{5.0, 6.0, 7.0}},
			C:           mat.General{Rows: 2, Cols: 2, RowsStride: 2, ColsStride: 1, Data: []float64{0.0, 0.0, 0.0, 0.0}},
			C_Hat:       mat.General{Rows: 2, Cols: 2, RowsStride: 2, ColsStride: 1, Data: []float64{0.0, 0.0, 0.0, 0.0}},
			A_Transpose: mat.NoTranspose,
			B_Transpose: mat.NoTranspose,
			alpha:       1.0,
			beta:        0.0,
			assertError: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := Bli_dgemm(test.A_Transpose, &test.alpha, test.A, test.B_Transpose, &test.beta, test.B, test.C)

			if test.assertError && err == nil {
				t.Errorf("%s failed: %v", t.Name(), err)
			}

			if !test.assertError {
				for i := 0; i < len(test.C.Data); i++ {
					if test.C.Data[i] != test.C_Hat.Data[i] {
						t.Errorf("%s failed: %f != %f", t.Name(), test.C.Data[i], test.C_Hat.Data[i])
					}
				}
			}
		})
	}
}
