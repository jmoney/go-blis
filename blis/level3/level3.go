package level3

// #cgo LDFLAGS: -lblis
// #include "blis.h"
import "C"
import (
	"github.com/jmoney/go-blas/mat"
)

// Bli_dgemm maps to bli_dgemm in BLIS.
//
// Computes the operation: C := (alpha * A') * (beta * B') + C
func Bli_dgemm(transa mat.TransT, transb mat.TransT, alpha *float64, a mat.General, beta *float64, b mat.General, c mat.General) {
	C.bli_dgemm(C.trans_t(transa), C.trans_t(transb),
		C.longlong(a.Rows), C.longlong(b.Cols), C.longlong(a.Cols),
		(*C.double)(alpha), (*C.double)(&a.Data[0]), C.longlong(a.RowsStride), C.longlong(a.ColsStride),
		(*C.double)(&b.Data[0]), C.longlong(b.RowsStride), C.longlong(b.ColsStride), (*C.double)(beta),
		(*C.double)(&c.Data[0]), C.longlong(c.RowsStride), C.longlong(c.ColsStride))
}
