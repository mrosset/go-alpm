package alpm

/*
#include <alpm.h>
*/
import "C"

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
