package alpm

import (
	"bytes"
	"reflect"
	"testing"
)

const pacmanConf = `
#
# GENERAL OPTIONS
#
[options]
RootDir     = /
DBPath      = /var/lib/pacman/
CacheDir    = /var/cache/pacman/pkg/ /other/cachedir
LogFile     = /var/log/pacman.log
GPGDir      = /etc/pacman.d/gnupg/
HoldPkg     = pacman glibc
# If upgrades are available for these packages they will be asked for first
SyncFirst   = pacman
#XferCommand = /usr/bin/curl -C - -f %u > %o
XferCommand = /usr/bin/wget --passive-ftp -c -O %o %u
CleanMethod = KeepInstalled
Architecture = x86_64

# Pacman won't upgrade packages listed in IgnorePkg and members of IgnoreGroup
IgnorePkg   = hello world
IgnoreGroup = kde

NoUpgrade   = kernel26
NoExtract   =

# Misc options
UseSyslog
#UseDelta
TotalDownload
CheckSpace
#VerbosePkgLists
ILoveCandy

# PGP signature checking
#SigLevel = Optional

[core]
SigLevel = Required
Server = ftp://ftp.example.com/foobar/$repo/os/$arch/

[custom]
SigLevel = Optional TrustAll
Server = file:///home/custompkgs
`

var pacmanConfRef = PacmanConfig{
	CacheDir:    []string{"/var/cache/pacman/pkg/", "/other/cachedir"},
	HoldPkg:     []string{"pacman", "glibc"},
	SyncFirst:   []string{"pacman"},
	IgnorePkg:   []string{"hello", "world"},
	IgnoreGroup: []string{"kde"},
	NoUpgrade:   []string{"kernel26"},
	NoExtract:   nil,

	RootDir:      "/",
	DBPath:       "/var/lib/pacman/",
	GPGDir:       "/etc/pacman.d/gnupg/",
	LogFile:      "/var/log/pacman.log",
	Architecture: "x86_64",
	XferCommand:  "/usr/bin/wget --passive-ftp -c -O %o %u",
	CleanMethod:  "KeepInstalled",

	Options: ConfUseSyslog | ConfTotalDownload | ConfCheckSpace | ConfILoveCandy,

	Repos: []RepoConfig{
		{Name: "core", Servers: []string{"ftp://ftp.example.com/foobar/$repo/os/$arch/"}},
		{Name: "custom", Servers: []string{"file:///home/custompkgs"}},
	},
}

func detailedDeepEqual(t *testing.T, x, y interface{}) {
	v := reflect.ValueOf(x)
	w := reflect.ValueOf(y)
	if v.Type() != w.Type() {
		t.Errorf("differing types %T vs. %T", x, y)
		return
	}
	for i := 0; i < v.NumField(); i++ {
		v_fld := v.Field(i).Interface()
		w_fld := w.Field(i).Interface()
		if !reflect.DeepEqual(v_fld, w_fld) {
			t.Errorf("field %s differs: got %#v, expected %#v",
				v.Type().Field(i).Name, v_fld, w_fld)
		}
	}
}

func TestPacmanConfigParser(t *testing.T) {
	buf := bytes.NewBufferString(pacmanConf)
	conf, err := ParseConfig(buf)
	if err != nil {
		t.Error(err)
	}

	detailedDeepEqual(t, conf, pacmanConfRef)
}
