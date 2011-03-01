package alpm

/*
#include <alpm.h>
*/
import "C"

type DataBase struct {
	pmdb *[0]uint8
}

func (v *DataBase) RegisterLocal() {
	v.pmdb = C.alpm_db_register_local()
}

func (v *DataBase) RegisterSync(s string) {
	v.pmdb = C.alpm_db_register_sync(C.CString(s))
}

func (v *DataBase) GetPkgCache() *AlpmList {
	return &AlpmList{C.alpm_db_get_pkgcache(v.pmdb)}
}

type Package struct {
	Pmpkg *[0]uint8
}

func (v *Package) GetName() string {
	return C.GoString(C.alpm_pkg_get_name(v.Pmpkg))
}
