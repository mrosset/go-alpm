package alpm

// #include <alpm.h>
import "C"

import "os"

// The Error type represents error codes from libalpm.
type Error uint32

var _ os.Error = Error(0)

// The string representation of an error is given by C function
// alpm_strerror().
func (er Error) String() string {
	return C.GoString(C.alpm_strerror(uint32(er)))
}
