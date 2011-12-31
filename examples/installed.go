package main

import alpm "github.com/remyoudompheng/go-alpm"

import "os"
import "fmt"

func main() {

	h, er := alpm.Init("/", "/var/lib/pacman")
	if er != nil {
		print(er, "\n")
		os.Exit(1)
	}

	db, er := h.LocalDb()
	if er != nil {
		fmt.Println(er)
		os.Exit(1)
	}

	for _, pkg := range db.PkgCache().Slice() {
		fmt.Printf("%s %s\n", pkg.Name(), pkg.Version())
	}

	if h.Release() != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
