package alpm

/*
#include <alpm.h>
*/
import "C"

import (
	"unsafe"
)

type Package struct {
	pmpkg *C.alpm_pkg_t
}

func (pkg Package) Name() string {
	return C.GoString(C.alpm_pkg_get_name(pkg.pmpkg))
}

func (pkg Package) Version() string {
	return C.GoString(C.alpm_pkg_get_version(pkg.pmpkg))
}

func (pkg Package) Description() string {
	return C.GoString(C.alpm_pkg_get_desc(pkg.pmpkg))
}

func (pkg Package) URL() string {
	return C.GoString(C.alpm_pkg_get_url(pkg.pmpkg))
}

func (pkg Package) Packager() string {
	return C.GoString(C.alpm_pkg_get_packager(pkg.pmpkg))
}

// Returns the names of reverse dependencies of a package
func (pkg Package) ComputeRequiredBy() []string {
	result := C.alpm_pkg_compute_requiredby(pkg.pmpkg)
	requiredby := make([]string, 0)
	for i := (*list)(unsafe.Pointer(result)); i != nil; i = i.Next {
		defer C.free(unsafe.Pointer(i))
		if i.Data != nil {
			defer C.free(unsafe.Pointer(i.Data))
			name := C.GoString((*C.char)(unsafe.Pointer(i.Data)))
			requiredby = append(requiredby, name)
		}
	}
	return requiredby
}
