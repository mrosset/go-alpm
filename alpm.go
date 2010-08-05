package alpm

/*
#include <alpm.h>
#include <sys/types.h>
#include <time.h>
#include <stdarg.h>
*/
import "C"


func Initialize() {
	C.alpm_initialize()
}

func Version() string {
	p := C.alpm_version()
	return C.GoString(p)
}
