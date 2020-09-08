// +build linux,amd64

package native

/*
#cgo LDFLAGS: -leccodes -lpng -laec -lpthread -fopenmp -lz -lm
*/
import "C"

type Cint = int32
type Clong = int64
type Culong = uint64
type Cdouble = float64
type CsizeT = int64
