package alpm

/*
#include <alpm.h>
#include <sys/types.h>
#include <time.h>
#include <stdarg.h>
*/
import "C"

import (
  "os"
)

// Initializes libalpm
func Initialize() os.Error {
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
