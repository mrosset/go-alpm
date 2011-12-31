package alpm

// #cgo CFLAGS: -D_FILE_OFFSET_BITS=64
// #include <alpm.h>
import "C"

import (
	"reflect"
	"unsafe"
)

// Description of a dependency.
type Depend struct {
	Name    string
	Version string
	Mod     DepMod
}

func convertDepend(dep *C.alpm_depend_t) Depend {
	return Depend{
		Name:    C.GoString(dep.name),
		Version: C.GoString(dep.version),
		Mod:     DepMod(dep.mod)}
}

func (dep Depend) String() string {
	return dep.Name + dep.Mod.String() + dep.Version
}

// Description of package files.
type File struct {
	Name string
	Size int64
	Mode uint32
}

func convertFilelist(files *C.alpm_filelist_t) []File {
	size := int(files.count)
	items := make([]File, size)

	raw_items := reflect.SliceHeader{
		Len:  size,
		Cap:  size,
		Data: uintptr(unsafe.Pointer(files.files))}

	c_files := *(*[]C.alpm_file_t)(unsafe.Pointer(&raw_items))

	for i := 0; i < size; i++ {
		items[i] = File{
			Name: C.GoString(c_files[i].name),
			Size: int64(c_files[i].size),
			Mode: uint32(c_files[i].mode)}
	}
	return items
}

// Internal alpm list structure.
type list struct {
	Data unsafe.Pointer
	Prev *list
	Next *list
}

// Iterates a function on a list and stop on error.
func (l *list) forEach(f func(unsafe.Pointer) error) error {
	for ; l != nil; l = l.Next {
		err := f(l.Data)
		if err != nil {
			return err
		}
	}
	return nil
}
