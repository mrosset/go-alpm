package alpm

// #cgo LDFLAGS: -lalpm
// #include <alpm.h>
import "C"

import (
	"os"
	"unsafe"
)

type Handle struct {
	ptr *C.alpm_handle_t
}

// Initialize
func Init(root, dbpath string) (*Handle, os.Error) {
	c_root := C.CString(root)
	defer C.free(unsafe.Pointer(c_root))
	c_dbpath := C.CString(dbpath)
	defer C.free(unsafe.Pointer(c_dbpath))
	var c_err C.enum__alpm_errno_t
	h := C.alpm_initialize(c_root, c_dbpath, &c_err)

	if c_err != 0 {
		return nil, Error(c_err)
	}

	return &Handle{h}, nil
}

func (h *Handle) Release() os.Error {
	if er := C.alpm_release(h.ptr); er != 0 {
		return Error(er)
	}
	h.ptr = nil
	return nil
}

// DB 

func (h Handle) GetRoot() string {
	return C.GoString(C.alpm_option_get_root(h.ptr))
}

func (h Handle) GetDbPath() string {
	return C.GoString(C.alpm_option_get_dbpath(h.ptr))
}

// Get the last pm_error
func (h Handle) LastError() os.Error {
	c_err := C.alpm_errno(h.ptr)
	if c_err != 0 {
		return Error(c_err)
	}
	return nil
}

func (h Handle) GetLocalDb() *[0]uint8 {
	return C.alpm_option_get_localdb(h.ptr)
}

// Helper functions
func Version() string {
	return C.GoString(C.alpm_version())
}
