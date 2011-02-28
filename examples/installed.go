package main

import "alpm"
import "os"

func main() {

	if alpm.Init() != nil {
		os.Exit(1)
	}

	if alpm.OptionSetRoot("/") != nil {
		alpm.Release()
		os.Exit(1)
	}

	if alpm.OptionSetDbpath("/var/lib/pacman") != nil {
		alpm.Release()
		os.Exit(1)
	}

	db_local := alpm.DbRegisterLocal()
	searchlist := alpm.DbGetPkgCache(db_local)

	for i := uint(0); i < searchlist.Count(); i++ {
		list := searchlist.Nth(i)
		pkg := list.GetData()
		name := alpm.PkgGetName(pkg)
		println(name)
	}

	if alpm.Release() != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
