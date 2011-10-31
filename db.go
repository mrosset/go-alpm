package alpm

/*
#include <alpm.h>
*/
import "C"

import (
	"os"
	"unsafe"
)

// Opaque structure representing a alpm database.
type Db struct {
	ptr *C.alpm_db_t
}

// Returns the local database relative to the given handle.
func (h Handle) GetLocalDb() (*Db, os.Error) {
	db := C.alpm_option_get_localdb(h.ptr)
	if db == nil {
		return nil, h.LastError()
	}
	return &Db{db}, nil
}

// Loads a sync database with given name and signature check level.
func (h Handle) RegisterSyncDb(dbname string, siglevel uint32) (*Db, os.Error) {
	c_name := C.CString(dbname)
	defer C.free(unsafe.Pointer(c_name))

	db := C.alpm_db_register_sync(h.ptr, c_name, C.alpm_siglevel_t(siglevel))
	if db == nil {
		return nil, h.LastError()
	}
	return &Db{db}, nil
}

// Returns the list of packages of the database
func (db Db) GetPkgCache() <-chan *Package {
	pkgcache := &AlpmList{C.alpm_db_get_pkgcache(db.ptr)}
	output := make(chan *Package)
	go func() {
		defer close(output)
		for i := pkgcache; i.Alpm_list_t != nil; i = i.Next() {
			pkg := &Package{i.GetData()}
			output <- pkg
		}
	}()
	return output
}
