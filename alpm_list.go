package alpm

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

/* allocation */
func (v *AlpmList) Free() {
  C.alpm_list_free(v.alpm_list_t)
}

/* mutators */
func (v *AlpmList) Add(data unsafe.Pointer) *AlpmList {
  return &AlpmList{C.alpm_list_add(v.alpm_list_t, data)}
}

func (v *AlpmList) Join(other *C.alpm_list_t) *AlpmList {
  return &AlpmList{C.alpm_list_join(v.alpm_list_t, other)}
}

/* accessors */
func (v *AlpmList) First() *AlpmList {
  return &AlpmList{C.alpm_list_first(v.alpm_list_t)}
}

func (v *AlpmList) Nth(n C.int) *AlpmList {
  return &AlpmList{C.alpm_list_nth(v.alpm_list_t, n)}
}

func (v *AlpmList) Next() *AlpmList {
  return &AlpmList{C.alpm_list_next(v.alpm_list_t)}
}

func (v *AlpmList) Last() *AlpmList {
  return &AlpmList{C.alpm_list_last(v.alpm_list_t)}
}

func (v *AlpmList) GetData() interface{} {
  return C.alpm_list_getdata(v.alpm_list_t)
}

/* misc */
func (v *AlpmList) Count() uint {
  return uint(C.alpm_list_count(v.alpm_list_t))
}

func (v *AlpmList) FindStr(needle *C.char) string {
  return C.GoString(C.alpm_list_find_str(v.alpm_list_t, needle))
}

func (v *AlpmList) FindPtr(needle unsafe.Pointer) interface{} {
  return C.alpm_list_find_ptr(v.alpm_list_t, needle)
}
