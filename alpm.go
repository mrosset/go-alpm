package alpm

/*
#include <alpm.h>
#include <sys/types.h>
#include <time.h>
#include <stdarg.h>
*/
import "C"


func Initialize() {
  C.alpm_initialize()
}

func Version() string {
  return C.GoString(C.alpm_version())
}
