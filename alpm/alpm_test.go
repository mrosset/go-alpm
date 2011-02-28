package alpm

import "os"
import "testing"

const (
	root    = "/"
	dbpath  = "/var/lib/pacman"
	version = "5.0.3"
)

func init() {
	err := Init()
	if err != nil {
		println("Init() failed w can not continue from here")
		os.Exit(1)
	}
}

func TestOptionSetDBpath(t *testing.T) {
	if err := OptionSetDbpath(dbpath); err != nil {
		t.Error(err)
	}
}

func TestOptionSetRoot(t *testing.T) {
	if err := OptionSetRoot(root); err != nil {
		t.Error(err)
	}
}

func TestVersion(t *testing.T) {
	if Version() != version {
		t.Error("verion's do not match")
	}
}

func TestRun(t *testing.T) {
	if !test() {
		t.Error("test() failed. note are objective is to get ride of this")
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
