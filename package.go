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

func (v *Package) Name() string {
	return C.GoString(C.alpm_pkg_get_name(v.pmpkg))
}

func (v *Package) Version() string {
	return C.GoString(C.alpm_pkg_get_version(v.pmpkg))
}

func (v *Package) Description() string {
	return C.GoString(C.alpm_pkg_get_desc(v.pmpkg))
}

func (v *Package) URL() string {
	return C.GoString(C.alpm_pkg_get_url(v.pmpkg))
}

func (v *Package) Packager() string {
	return C.GoString(C.alpm_pkg_get_packager(v.pmpkg))
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
