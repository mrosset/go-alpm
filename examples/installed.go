package main

import "alpm"
import "os"
import "fmt"

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

	db := alpm.RegisterLocalDb()
	searchlist := alpm.GetPkgCache(db)

	for i := uint(0); i < searchlist.Count(); i++ {
		list := searchlist.Nth(i)
		pkg := &alpm.Package{list.GetData()}
		name := pkg.GetName()
		fmt.Printf("%v \n",name)
	}

	if alpm.Release() != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
