package alpm

/*
#include <alpm.h>
#include <stdio.h>
#include <alpm_list.h>
#include <sys/types.h>
#include <time.h>
#include <stdarg.h>
*/
import "C"

import (
  "os"
  "fmt"
)

// Initializes libalpm
func Init() os.Error {
  if C.alpm_initialize() != 0 {
    return lastError()
  }
  return nil
}

// Release libalpm
func Release() os.Error {
  if C.alpm_release() != 0 {
    return lastError()
  }
  return nil
}

// Returns libalpm version
func Version() string {
  return C.GoString(C.alpm_version())
}

// Get the last pm_error
func lastError() os.Error {
  return os.NewError(C.GoString(C.alpm_strerrorlast()))
}

func Test {
	C.alpm_list_t *i;

	if(C.alpm_initialize() != 0) {
		C.printf("could not int alpm\n")
        return
	}

	if (C.alpm_option_set_root("/") != 0) {
		C.printf("failed setting root option\n")
		C.alpm_release();
		return
	}

  if (C.alpm_option_set_dbpath("/var/lib/pacman") != 0) {
		C.printf("failed setting db path\n")
		C.alpm_release()
		return
	}

	C.printf("root = %s\n", C.alpm_option_get_root());
	C.printf("dbpath = %s\n", C.alpm_option_get_dbpath())

	C.pmdb_t *db_local = C.alpm_db_register_local()
	C.alpm_list_t *searchlist = C.alpm_db_get_pkgcache(db_local)

	/*for(i = searchlist; i ; i = alpm_list_next(i)) {
		pmpkg_t *pkg = alpm_list_getdata(i);
		printf("local/%s %s\n", alpm_pkg_get_name(pkg), alpm_pkg_get_version(pkg));
	}*/

	if(C.alpm_release() != 0) {
		C.printf("could not release alpm\n");
	}
	return
}
