package alpm

/*
#include <alpm.h>
*/
import "C"

type Package struct {
  Pmpkg *[0]uint8
}

func (v *Package) GetName() string {
  return C.GoString(C.alpm_pkg_get_name(v.Pmpkg))
}
