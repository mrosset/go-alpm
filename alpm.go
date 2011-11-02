package alpm

// #cgo LDFLAGS: -lalpm
// #include <alpm.h>
import "C"

import "unsafe"

type Handle struct {
	ptr *C.alpm_handle_t
}

// Initialize
func Init(root, dbpath string) (*Handle, error) {
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

func (h *Handle) Release() error {
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
func (h Handle) LastError() error {
	if h.ptr != nil {
		c_err := C.alpm_errno(h.ptr)
		if c_err != 0 {
			return Error(c_err)
		}
	}
	return nil
}

// Helper functions
func Version() string {
	return C.GoString(C.alpm_version())
}

// Perform version comparison according to Pacman conventions.
// Return value is <0 if and only if v1 is older than v2.
func VerCmp(v1, v2 string) int {
	c1 := C.CString(v1)
	c2 := C.CString(v2)
	defer C.free(unsafe.Pointer(c1))
	defer C.free(unsafe.Pointer(c2))
	result := C.alpm_pkg_vercmp(c1, c2)
	return int(result)
}
