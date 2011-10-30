package alpm

// #include <alpm.h>
import "C"

type Error uint32

func (er Error) String() string {
	return C.GoString(C.alpm_strerror(uint32(er)))
}
