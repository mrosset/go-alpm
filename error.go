package alpm

// #include <alpm.h>
import "C"

// The Error type represents error codes from libalpm.
type Error uint32

var _ error = Error(0)

// The string representation of an error is given by C function
// alpm_strerror().
func (er Error) Error() string {
	return C.GoString(C.alpm_strerror(uint32(er)))
}
