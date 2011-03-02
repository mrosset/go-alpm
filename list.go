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
  Alpm_list_t *C.alpm_list_t
}

/* allocation */
func (v *AlpmList) Free() {
  C.alpm_list_free(v.Alpm_list_t)
}

/* mutators */
func (v *AlpmList) Add(data interface{}) {
  v.Alpm_list_t = C.alpm_list_add(v.Alpm_list_t, unsafe.New(data))
}

func (v *AlpmList) AddSorted(data interface{}, callback *[0]uint8) {
  v = &AlpmList{C.alpm_list_add_sorted(v.Alpm_list_t, unsafe.New(data), callback)}
}

func (v *AlpmList) Join(other *C.alpm_list_t) {
  v = &AlpmList{C.alpm_list_join(v.Alpm_list_t, other)}
}

/* accessors */
func (v *AlpmList) First() *AlpmList {
  return &AlpmList{C.alpm_list_first(v.Alpm_list_t)}
}

func (v *AlpmList) Nth(n uint) *AlpmList {
  return &AlpmList{C.alpm_list_nth(v.Alpm_list_t, C.int(n))}
}

func (v *AlpmList) Next() *AlpmList {
  return &AlpmList{C.alpm_list_next(v.Alpm_list_t)}
}

func (v *AlpmList) Last() *AlpmList {
  return &AlpmList{C.alpm_list_last(v.Alpm_list_t)}
}

func (v *AlpmList) GetData() *[0]uint8 {
  return (*[0]uint8)(v.Alpm_list_t.data)
}

/* misc */
func (v *AlpmList) Count() uint {
  return uint(C.alpm_list_count(v.Alpm_list_t))
}

func (v *AlpmList) FindStr(needle *C.char) string {
  return C.GoString(C.alpm_list_find_str(v.Alpm_list_t, needle))
}

func (v *AlpmList) FindPtr(needle interface{}) interface{} {
  return C.alpm_list_find_ptr(v.Alpm_list_t, unsafe.New(needle))
}
