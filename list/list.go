package list

/*
#include <stdlib.h>
#include <alpm_list.h>
*/
import "C"

import (
  "unsafe"
)

type AlpmList struct {
  alpm_list_t *C.alpm_list_t
}

func (v AlpmList) Free() {
  C.alpm_list_free(v.alpm_list_t)
}

func (v AlpmList) Add(data unsafe.Pointer) *AlpmList {
  return &AlpmList{C.alpm_list_add(v.alpm_list_t, data)}
}
