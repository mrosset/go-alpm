package alpm

import (
	"fmt"
	"os"
	"testing"
)

const (
	root    = "/"
	dbpath  = "/var/lib/pacman"
	version = "7.0.0"
)

var h *Handle

func init() {
	var err os.Error
	h, err = Init("/", "/var/lib/pacman")
	if err != nil {
		fmt.Printf("failed to Init(): %s", err)
		os.Exit(1)
	}
}

func TestVersion(t *testing.T) {
	if Version() != version {
		t.Error("verion's do not match")
	}
}

func TestVercmp(t *testing.T) {
	x := VerCmp("1.0-2", "2.0-1")
	if x >= 0 {
		t.Error("failed at checking 2.0-1 is newer than 1.0-2")
	}
	x = VerCmp("1:1.0-2", "2.0-1")
	if x <= 0 {
		t.Error("failed at checking 2.0-1 is older than 1.0-2")
	}
	x = VerCmp("2.0.2-2", "2.0.2-2")
	if x != 0 {
		t.Error("failed at checking 2.0.2-2 is equal to itself")
	}
}

func TestLocalDB(t *testing.T) {
	defer func() {
		if recover() != nil {
			t.Errorf("local db failed")
		}
	}()
	db, _ := h.GetLocalDb()
	searchlist := db.GetPkgCache()
	for i := searchlist.Next(); i.Alpm_list_t != nil; i = i.Next() {
		pkg := &Package{i.GetData()}
		fmt.Printf("%v \n", pkg.Name())
	}
}

func TestRelease(t *testing.T) {
	if err := h.Release(); err != nil {
		t.Error(err)
	}
}
