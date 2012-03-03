package alpm

/*
#include <alpm.h>
void logCallback(unsigned int level, char *cstring);
void go_alpm_log_cb(alpm_loglevel_t level, const char *fmt, va_list arg);
void go_alpm_set_logging(alpm_handle_t *handle);
*/
import "C"

var DefaultLogLevel = LogWarning

func DefaultLogCallback(lvl uint, s string) {
	if lvl <= DefaultLogLevel {
		print("go-alpm: ", s)
	}
}

var log_callback = DefaultLogCallback

//export logCallback
func logCallback(level uint, cstring *C.char) {
	log_callback(level, C.GoString(cstring))
}

func (h *Handle) SetLogCallback(cb func(uint, string)) {
	log_callback = cb
	C.go_alpm_set_logging(h.ptr)
}
