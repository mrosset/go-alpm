package alpm

/*
#include <alpm.h>
*/
import "C"
import (
	"os"
	"fmt"
)

var (
	initialized bool = false
	println          = fmt.Println
)

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

func test() os.Error {
	db := GetLocalDB()
	searchlist := GetPkgCache(db)
	for i := uint(0); i < searchlist.Count(); i++ {
		list := searchlist.Nth(i)
		pkg := &Package{list.GetData()}
		name := pkg.GetName()
		fmt.Printf("%v \n", name)
	}
	return nil
}

func GetLocalDB() *[0]uint8 {
	return C.alpm_option_get_localdb()
}

func SetOptions() (err os.Error) {
	if err = SetRoot("/"); err != nil {
		return err
	}
	if err = SetDbPath("/var/lib/pacman"); err != nil {
		return err
	}
	return
}

func prints(prefix string, s *_Ctype_char) {
	fmt.Printf("%v = %v\n", prefix, C.GoString(s))
}

func printT(i interface{}) {
	fmt.Printf("%T = %v\n", i, i)
}
