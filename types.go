package alpm

// int dummy;
import "C"

import "unsafe"

type Depend struct {
	Name    string
	Version string
	Mod     DepMod
}

func convertDepend(dep depend) Depend {
	return Depend{
		Name:    C.GoString((*C.char)(unsafe.Pointer(dep.Name))),
		Version: C.GoString((*C.char)(unsafe.Pointer(dep.Version))),
		Mod:     DepMod(dep.Mod)}
}

func (dep Depend) String() string {
	return dep.Name + dep.Mod.String() + dep.Version
}
