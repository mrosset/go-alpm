package alpm

/*
#include <alpm.h>
*/
import "C"

import (
	"fmt"
	"io"
	"unsafe"
)

// Opaque structure representing a alpm database.
type Db struct {
	ptr    *C.alpm_db_t
	handle Handle
}

type DbList struct {
	*list
	handle Handle
}

func (l DbList) ForEach(f func(Db) error) error {
	return l.forEach(func(p unsafe.Pointer) error {
		return f(Db{(*C.alpm_db_t)(p), l.handle})
	})
}

func (l DbList) Slice() []Db {
	slice := []Db{}
	l.ForEach(func(db Db) error {
		slice = append(slice, db)
		return nil
	})
	return slice
}

// Returns the local database relative to the given handle.
func (h Handle) LocalDb() (*Db, error) {
	db := C.alpm_get_localdb(h.ptr)
	if db == nil {
		return nil, h.LastError()
	}
	return &Db{db, h}, nil
}

func (h Handle) SyncDbs() (DbList, error) {
	dblist := C.alpm_get_syncdbs(h.ptr)
	if dblist == nil {
		return DbList{nil, h}, h.LastError()
	}
	dblistPtr := unsafe.Pointer(dblist)
	return DbList{(*list)(dblistPtr), h}, nil
}

// SyncDbByName finds a registered database by name.
func (h Handle) SyncDbByName(name string) (db *Db, err error) {
	dblist, err := h.SyncDbs()
	if err != nil {
		return nil, err
	}
	dblist.ForEach(func(b Db) error {
		if b.Name() == name {
			db = &b
			return io.EOF
		}
		return nil
	})
	if db != nil {
		return db, nil
	}
	return nil, fmt.Errorf("database %s not found", name)
}

// Loads a sync database with given name and signature check level.
func (h Handle) RegisterSyncDb(dbname string, siglevel SigLevel) (*Db, error) {
	c_name := C.CString(dbname)
	defer C.free(unsafe.Pointer(c_name))

	db := C.alpm_register_syncdb(h.ptr, c_name, C.alpm_siglevel_t(siglevel))
	if db == nil {
		return nil, h.LastError()
	}
	return &Db{db, h}, nil
}

func (db Db) Name() string {
	return C.GoString(C.alpm_db_get_name(db.ptr))
}

func (db Db) Servers() []string {
	ptr := unsafe.Pointer(C.alpm_db_get_servers(db.ptr))
	return StringList{(*list)(ptr)}.Slice()
}

func (db Db) SetServers(servers []string) {
	C.alpm_db_set_servers(db.ptr, nil)
	for _, srv := range servers {
		C_srv := C.CString(srv)
		defer C.free(unsafe.Pointer(C_srv))
		C.alpm_db_add_server(db.ptr, C_srv)
	}
}

func (db Db) PkgByName(name string) (*Package, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	ptr := C.alpm_db_get_pkg(db.ptr, c_name)
	if ptr == nil {
		return nil,
			fmt.Errorf("Error when retrieving %s from database %s: %s",
				name, db.Name(), db.handle.LastError())
	}
	return &Package{ptr, db.handle}, nil
}

// Returns the list of packages of the database
func (db Db) PkgCache() PackageList {
	pkgcache := (*list)(unsafe.Pointer(C.alpm_db_get_pkgcache(db.ptr)))
	return PackageList{pkgcache, db.handle}
}
