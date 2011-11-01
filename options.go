package alpm

// #include <alpm.h>
import "C"

import (
 "os"
)

func (h Handle) GetUseSyslog() bool {
	value := C.alpm_option_get_usesyslog(h.ptr)
	return (value != 0)
}

func (h Handle) SetUseSyslog(value bool) os.Error {
	var int_value C.int
	if value {
		int_value = 1
	} else {
		int_value = 0
	}
	ok := C.alpm_option_set_usesyslog(h.ptr, int_value)
	if ok < 0 {
		return h.LastError()
	}
	return nil
}
