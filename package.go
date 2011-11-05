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

func (pkg Package) MD5Sum() string {
	return C.GoString(C.alpm_pkg_get_md5sum(pkg.pmpkg))
}

func (pkg Package) SHA256Sum() string {
	return C.GoString(C.alpm_pkg_get_sha256sum(pkg.pmpkg))
}

func (pkg Package) BuildDate() int64 {
	t := C.alpm_pkg_get_builddate(pkg.pmpkg)
	return int64(t)
}

func (pkg Package) InstallDate() int64 {
	t := C.alpm_pkg_get_installdate(pkg.pmpkg)
	return int64(t)
}

func (pkg Package) Size() int64 {
	t := C.alpm_pkg_get_size(pkg.pmpkg)
	return int64(t)
}

func (pkg Package) ISize() int64 {
	t := C.alpm_pkg_get_isize(pkg.pmpkg)
	return int64(t)
}

func (pkg Package) DB() *Db {
	ptr := C.alpm_pkg_get_db(pkg.pmpkg)
	if ptr == nil {
		return nil
	}
	return &Db{ptr}
}

func iterateDepends(l *C.alpm_list_t) <-chan Depend {
	out := make(chan Depend)
	go func() {
		defer close(out)
		for i := (*list)(unsafe.Pointer(l)); i != nil; i = i.Next {
			item := (*depend)(unsafe.Pointer(i.Data))
			out <- convertDepend(*item)
		}
	}()
	return out
}

func (pkg Package) Depends() <-chan Depend {
	c_depends := C.alpm_pkg_get_depends(pkg.pmpkg)
	return iterateDepends(c_depends)
}

func (pkg Package) Conflicts() <-chan Depend {
	c_depends := C.alpm_pkg_get_conflicts(pkg.pmpkg)
	return iterateDepends(c_depends)
}

func (pkg Package) Provides() <-chan Depend {
	c_depends := C.alpm_pkg_get_provides(pkg.pmpkg)
	return iterateDepends(c_depends)
}

func (pkg Package) Replaces() <-chan Depend {
	c_depends := C.alpm_pkg_get_replaces(pkg.pmpkg)
	return iterateDepends(c_depends)
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
