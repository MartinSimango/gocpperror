package gocpperror

// #cgo LDFLAGS: -lgoerror
// #include <cpperror/ErrorWrapper.h>
import "C"

const INT_TYPE = int(C.INT_TYPE)
const BOOL_TYPE = int(C.BOOL_TYPE)
const STRING_TYPE = int(C.STRING_TYPE)
const DOUBLE_TYPE = int(C.DOUBLE_TYPE)
