package alpm

/*
#include <alpm.h>
*/
import "C"

import (
	"unsafe"
)

type Package struct {
	pmpkg  *C.alpm_pkg_t
	handle Handle
}

type PackageList struct {
	*list
	handle Handle
}

func (l PackageList) ForEach(f func(Package) error) error {
	return l.forEach(func(p unsafe.Pointer) error {
		return f(Package{(*C.alpm_pkg_t)(p), l.handle})
	})
}

func (l PackageList) Slice() []Package {
	slice := []Package{}
	l.ForEach(func(p Package) error {
		slice = append(slice, p)
		return nil
	})
	return slice
}

type DependList struct{ *list }

func (l DependList) ForEach(f func(Depend) error) error {
	return l.forEach(func(p unsafe.Pointer) error {
		dep := convertDepend((*C.alpm_depend_t)(p))
		return f(dep)
	})
}

func (l DependList) Slice() []Depend {
	slice := []Depend{}
	l.ForEach(func(dep Depend) error {
		slice = append(slice, dep)
		return nil
	})
	return slice
}

func (pkg Package) Name() string {
	return C.GoString(C.alpm_pkg_get_name(pkg.pmpkg))
}

func (pkg Package) Version() string {
	return C.GoString(C.alpm_pkg_get_version(pkg.pmpkg))
}

func (pkg Package) Architecture() string {
	return C.GoString(C.alpm_pkg_get_arch(pkg.pmpkg))
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
	return &Db{ptr, pkg.handle}
}

func (pkg Package) Files() []File {
	c_files := C.alpm_pkg_get_files(pkg.pmpkg)
	return convertFilelist(c_files)
}

func (pkg Package) Depends() DependList {
	ptr := unsafe.Pointer(C.alpm_pkg_get_depends(pkg.pmpkg))
	return DependList{(*list)(ptr)}
}

func (pkg Package) Conflicts() DependList {
	ptr := unsafe.Pointer(C.alpm_pkg_get_conflicts(pkg.pmpkg))
	return DependList{(*list)(ptr)}
}

func (pkg Package) Provides() DependList {
	ptr := unsafe.Pointer(C.alpm_pkg_get_provides(pkg.pmpkg))
	return DependList{(*list)(ptr)}
}

func (pkg Package) Replaces() DependList {
	ptr := unsafe.Pointer(C.alpm_pkg_get_replaces(pkg.pmpkg))
	return DependList{(*list)(ptr)}
}

func (pkg Package) Groups() StringList {
	ptr := unsafe.Pointer(C.alpm_pkg_get_groups(pkg.pmpkg))
	return StringList{(*list)(ptr)}
}

func (pkg Package) Licenses() StringList {
	ptr := unsafe.Pointer(C.alpm_pkg_get_licenses(pkg.pmpkg))
	return StringList{(*list)(ptr)}
}

func (pkg Package) Reason() PkgReason {
	reason := C.alpm_pkg_get_reason(pkg.pmpkg)
	return PkgReason(reason)
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
