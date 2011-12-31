package alpm

/*
#include <alpm.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// Opaque structure representing a alpm database.
type Db struct {
	ptr *C.alpm_db_t
}

// Returns the local database relative to the given handle.
func (h Handle) LocalDb() (*Db, error) {
	db := C.alpm_option_get_localdb(h.ptr)
	if db == nil {
		return nil, h.LastError()
	}
	return &Db{db}, nil
}

// Loads a sync database with given name and signature check level.
func (h Handle) RegisterSyncDb(dbname string, siglevel SigLevel) (*Db, error) {
	c_name := C.CString(dbname)
	defer C.free(unsafe.Pointer(c_name))

	db := C.alpm_db_register_sync(h.ptr, c_name, C.alpm_siglevel_t(siglevel))
	if db == nil {
		return nil, h.LastError()
	}
	return &Db{db}, nil
}

func (db Db) Name() string {
	return C.GoString(C.alpm_db_get_name(db.ptr))
}

func (db Db) GetPkg(name string) (*Package, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	ptr := C.alpm_db_get_pkg(db.ptr, c_name)
	if ptr == nil {
		return nil,
			fmt.Errorf("Error when retrieving %s from database %s, see Handle.LastError()",
				name, db.Name())
	}
	return &Package{ptr}, nil
}

// Returns the list of packages of the database
func (db Db) PkgCache() PackageList {
	pkgcache := (*list)(unsafe.Pointer(C.alpm_db_get_pkgcache(db.ptr)))
  return PackageList {pkgcache}
}
