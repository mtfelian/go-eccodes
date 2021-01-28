// +build windows,amd64

package native

/*
#cgo LDFLAGS: -leccodes -lpng -laec -ljasper -lz
*/
import "C"

type Cint = int32
type Clong = int32
type Culong = uint32
type Cdouble = float64
type CsizeT = int64
