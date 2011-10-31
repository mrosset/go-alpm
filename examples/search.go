package main

import alpm "github.com/remyoudompheng/go-alpm"

import "fmt"

func main() {
	h, er := alpm.Init("/", "/var/lib/pacman")
	if er != nil {
		fmt.Println(er)
		return
	}
	defer h.Release()

	db, _ := h.RegisterSyncDb("core", 0)
	h.RegisterSyncDb("community", 0)
	h.RegisterSyncDb("extra", 0)

	for pkg := range db.GetPkgCache() {
		fmt.Printf("%s %s\n  %s\n",
			pkg.Name(), pkg.Version(), pkg.Description())
	}
}
