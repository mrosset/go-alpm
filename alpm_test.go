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
