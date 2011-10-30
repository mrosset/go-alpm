package alpm

/*
#include <alpm.h>
*/
import "C"

func (h Handle) RegisterSyncDb(s string, siglevel uint32) *[0]uint8 {
	return C.alpm_db_register_sync(h.ptr, C.CString(s), C.alpm_siglevel_t(siglevel))
}

func GetPkgCache(p *[0]uint8) *AlpmList {
	return &AlpmList{C.alpm_db_get_pkgcache(p)}
}
