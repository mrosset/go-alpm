package alpm

// #cgo LDFLAGS: -lalpm
// #include <alpm.h>
import "C"
import (
	"os"
	"fmt"
)

var (
	initialized bool = false
	println          = fmt.Println
)

// Initialize
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

func Release() os.Error {
	if C.alpm_release() != 0 {
		return LastError()
	}
	initialized = false
	return nil
}


// DB 

// Options
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

// Get the last pm_error
func LastError() os.Error {
	return os.NewError(C.GoString(C.alpm_strerrorlast()))
}

func GetLocalDb() *[0]uint8 {
	return C.alpm_option_get_localdb()
}

// Helper functions
func Version() string {
	return C.GoString(C.alpm_version())
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

// private test functions
func prints(prefix string, s *_Ctype_char) {
	fmt.Printf("%v = %v\n", prefix, C.GoString(s))
}

func printT(i interface{}) {
	fmt.Printf("%T = %v\n", i, i)
}

func test() os.Error {
	db := GetLocalDb()
	searchlist := GetPkgCache(db)
	for i := searchlist.Next(); i.Alpm_list_t != nil; i = i.Next() {
		pkg := &Package{i.GetData()}
		fmt.Printf("%v \n", pkg.GetName())
	}
	return nil
}
