package level3

// #cgo LDFLAGS: -lblis
// #include "blis.h"
import "C"
import (
	"github.com/jmoney/go-blas/mat"
)

/*
void bli_?gemm(

	  trans_t  transa,
	  trans_t  transb,
	  dim_t    m,
	  dim_t    n,
	  dim_t    k,
	  ctype*   alpha,
	  ctype*   a, inc_t rsa, inc_t csa,
	  ctype*   b, inc_t rsb, inc_t csb,
	  ctype*   beta,
	  ctype*   c, inc_t rsc, inc_t csc
	);
*/
func Bli_dgemm(transa mat.TransT, transb mat.TransT, alpha *float64, a mat.General, beta *float64, b mat.General, c mat.General) {
	C.bli_dgemm(C.trans_t(transa), C.trans_t(transb), C.longlong(a.Rows), C.longlong(b.Cols), C.longlong(a.Cols), (*C.double)(alpha), (*C.double)(&a.Data[0]), C.longlong(a.RowsStride), C.longlong(a.ColsStride), (*C.double)(&b.Data[0]), C.longlong(b.RowsStride), C.longlong(b.ColsStride), (*C.double)(beta), (*C.double)(&c.Data[0]), C.longlong(c.RowsStride), C.longlong(c.ColsStride))
}
