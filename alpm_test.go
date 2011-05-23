package alpm

import (
	"fmt"
	"os"
	"testing"
)

const (
	root    = "/"
	dbpath  = "/var/lib/pacman"
	version = "6.0.2"
)

func init() {
	err := Init()
	if err != nil {
		println("Init() failed w can not continue from here")
		os.Exit(1)
	}
}

func TestSetDBPath(t *testing.T) {
	if err := SetDbPath(dbpath); err != nil {
		t.Error(err)
	}
}

func TestSetRoot(t *testing.T) {
	if err := SetRoot(root); err != nil {
		t.Error(err)
	}
}

func TestVersion(t *testing.T) {
	if Version() != version {
		t.Error("verion's do not match")
	}
}

func TestLocalDB(t *testing.T) {
	defer func() {
		if recover() != nil {
			t.Errorf("local db failed")
		}
	}()
	db := GetLocalDb()
	searchlist := GetPkgCache(db)
	for i := searchlist.Next(); i.Alpm_list_t != nil; i = i.Next() {
		pkg := &Package{i.GetData()}
		fmt.Printf("%v \n", pkg.GetName())
	}
}

func TestRelease(t *testing.T) {
	if !initialized {
		Init()
	}
	if err := Release(); err != nil {
		t.Error(err)
	}
}
