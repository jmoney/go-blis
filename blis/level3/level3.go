package level3

// #cgo LDFLAGS: -lblis
// #include "blis.h"
import "C"
import (
	"fmt"

	"github.com/jmoney/go-lapack/mat"
)

// Bli_dgemm maps to bli_dgemm in BLIS.
//
// Computes the double precision operation: C := (alpha * A') * (beta * B') + C
func Bli_dgemm(transa mat.TransT, alpha *float64, a mat.General, transb mat.TransT, beta *float64, b mat.General, c mat.General) error {

	if err := validateDimensions(transa, a, transb, b, c); err != nil {
		return err
	}

	C.bli_dgemm(C.trans_t(transa), C.trans_t(transb),
		C.longlong(a.Rows), C.longlong(b.Cols), C.longlong(a.Cols),
		(*C.double)(alpha), (*C.double)(&a.Data[0]), C.longlong(a.RowsStride), C.longlong(a.ColsStride),
		(*C.double)(&b.Data[0]), C.longlong(b.RowsStride), C.longlong(b.ColsStride), (*C.double)(beta),
		(*C.double)(&c.Data[0]), C.longlong(c.RowsStride), C.longlong(c.ColsStride))

	return nil
}

func validateDimensions(transa mat.TransT, a mat.General, transb mat.TransT, b mat.General, c mat.General) error {
	if transa == mat.NoTranspose && transb == mat.NoTranspose {
		if a.Cols != b.Rows || a.Rows != c.Rows || b.Cols != c.Cols {
			return fmt.Errorf("dimensions do not match")
		}
	} else if transa == mat.Transpose && transb == mat.NoTranspose {
		if a.Rows != b.Rows || a.Cols != c.Rows || b.Cols != c.Cols {
			return fmt.Errorf("dimensions do not match")
		}
	} else if transa == mat.NoTranspose && transb == mat.Transpose {
		if a.Cols != b.Cols || a.Rows != c.Rows || b.Rows != c.Cols {
			return fmt.Errorf("dimensions do not match")
		}
	} else if transa == mat.Transpose && transb == mat.Transpose {
		if a.Rows != b.Cols || a.Cols != c.Rows || b.Rows != c.Cols {
			return fmt.Errorf("dimensions do not match")
		}
	}

	return nil
}
