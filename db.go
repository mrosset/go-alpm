package alpm

/*
#include <alpm.h>
*/
import "C"

func RegisterLocalDb() *[0]uint8 {
  return C.alpm_db_register_local()
}

func RegisterSyncDb(s string) *[0]uint8 {
  return C.alpm_db_register_sync(C.CString(s))
}

func GetPkgCache(p *[0]uint8) *AlpmList {
  return &AlpmList{C.alpm_db_get_pkgcache(p)}
}
