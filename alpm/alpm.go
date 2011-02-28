package alpm

/*
#include <alpm.h>
#include <alpm_list.h>
*/
import "C"
import (
	"C"
	"os"
	"fmt"
)

var initialized = false

// Initializes libalpm
func Init() os.Error {
	if initialized {
		return nil
	}

	if C.alpm_initialize() != 0 {
		return LastError()
	}
	initialized = true
	return nil
}

// Release libalpm
func Release() os.Error {
	if C.alpm_release() != 0 {
		return LastError()
	}
	initialized = false
	return nil
}

func OptionSetRoot(s string) os.Error {
	cs := C.CString(s)
	if C.alpm_option_set_root(cs) != 0 {
		return LastError()
	}
	return nil
}

func OptionSetDbpath(s string) os.Error {
	cs := C.CString(s)
	if C.alpm_option_set_dbpath(cs) != 0 {
		return LastError()
	}
	return nil
}

func OptionGetRoot() string {
	return C.GoString(C.alpm_option_get_root())
}

func OptionGetDBpath() string {
	return C.GoString(C.alpm_option_get_dbpath())
}

// Returns libalpm version
func Version() string {
	return C.GoString(C.alpm_version())
}

func ListNext() {
}

// Get the last pm_error
func LastError() os.Error {
	return os.NewError(C.GoString(C.alpm_strerrorlast()))
}

func test() bool {

	if Init() != nil {
		return false
	}

	if OptionSetRoot("/") != nil {
		Release()
		return false
	}

	if OptionSetDbpath("/var/lib/pacman") != nil {
		Release()
		return false
	}

	db_local := C.alpm_db_register_local()
	searchlist := C.alpm_db_get_pkgcache(db_local)

	for i := searchlist; i != nil; i = C.alpm_list_next(i) {
		//printT(C.alpm_list_getdata(i))
		//pkg := (*[0]uint8)(C.alpm_list_getdata(i))
		//printfl(C.GoString(C.alpm_pkg_get_name(pkg)))
	}

	if Release() != nil {
		return false
	}
	return true
}

func printfl(format string, i ...interface{}) {
	fmt.Printf(format+"\n", i...)
}


func printT(i interface{}) {
	printfl("%T = %v", i, i)
}
