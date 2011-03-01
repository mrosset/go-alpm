package main

import "alpm"
import "os"

func main() {

	if alpm.Init() != nil {
		os.Exit(1)
	}

	if alpm.SetRoot("/") != nil {
		alpm.Release()
		os.Exit(1)
	}

	if alpm.SetDbPath("/var/lib/pacman") != nil {
		alpm.Release()
		os.Exit(1)
	}

	db := &alpm.DataBase{}
	db.RegisterLocal()
	searchlist := db.GetPkgCache()

	for i := uint(0); i < searchlist.Count(); i++ {
		list := searchlist.Nth(i)
		pkg := &alpm.Package{list.GetData()}
		name := pkg.GetName()
		println(name)
	}

	if alpm.Release() != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
