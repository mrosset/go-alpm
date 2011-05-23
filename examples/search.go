package main

import alpm "github.com/str1ngs/go-alpm"
import "fmt"

func main() {
	alpm.Init()
	defer alpm.Release()
	alpm.SetRoot("/")
	alpm.SetDbPath("/var/lib/pacman")
	db := alpm.RegisterSyncDb("core")
	alpm.RegisterSyncDb("community")
	alpm.RegisterSyncDb("extra")
	searchlist := alpm.GetPkgCache(db)
	for i := uint(0); i < searchlist.Count(); i++ {
		list := searchlist.Nth(i)
		pkg := &alpm.Package{list.GetData()}
		fmt.Printf("%s\n", pkg.GetName())
	}
}
