package alpm

/*
#include <alpm.h>
*/
import "C"

func DbRegisterLocal() *[0]uint8 {
    return C.alpm_db_register_local() 
}

func DbGetPkgCache(p *[0]uint8) *AlpmList {
    return &AlpmList{C.alpm_db_get_pkgcache(p)}
}

func PkgGetName(p *[0]uint8) string {
    return C.GoString(C.alpm_pkg_get_name(p))
}
