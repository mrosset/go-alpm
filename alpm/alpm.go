package alpm

/*
#include <alpm.h>
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

func SetRoot(s string) os.Error {
	cs := C.CString(s)
	if C.alpm_option_set_root(cs) != 0 {
		return LastError()
	}
	return nil
}

func SetDbPath(s string) os.Error {
	cs := C.CString(s)
	if C.alpm_option_set_dbpath(cs) != 0 {
		return LastError()
	}
	return nil
}

func GetRoot() string {
	return C.GoString(C.alpm_option_get_root())
}

func GetDbPath() string {
	return C.GoString(C.alpm_option_get_dbpath())
}

// Returns libalpm version
func Version() string {
	return C.GoString(C.alpm_version())
}

// Get the last pm_error
func LastError() os.Error {
	return os.NewError(C.GoString(C.alpm_strerrorlast()))
}

func test() bool {
	Init()
	defer Release()
	SetRoot("/")
	SetDbPath("/var/lib/pacman")
	db := &DataBase{}
	db.RegisterSync("core")
	db.RegisterSync("community")
	db.RegisterSync("extra")
	searchlist := db.GetPkgCache()
	printT(searchlist)
	for i := uint(0); i < searchlist.Count(); i++ {
		list := searchlist.Nth(i)
		pkg := &Package{list.GetData()}
		println(pkg.GetName())
	}
	fmt.Printf("%v\n", LastError().String())
	return true
}


func prints(prefix string, s *_Ctype_char) {
	fmt.Printf("%v = %v\n", prefix, C.GoString(s))
}

func printT(i interface{}) {
	fmt.Printf("%T = %v\n", i, i)
}
