package alpm

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
	"time"
)

// Auxiliary formatting
const pkginfo_template = `
Name         : {{ .Name }}
Version      : {{ .Version }}
Description  : {{ .Description }}
URL          : {{ .URL }}
Dependencies : {{ range .Depends.Slice }}{{ . }} {{ end }}
Provides     : {{ range .Provides.Slice }}{{ . }} {{ end }}
Replaces     : {{ range .Replaces.Slice }}{{ . }} {{ end }}
Conflicts    : {{ range .Conflicts.Slice }}{{ . }} {{ end }}
Packager     : {{ .Packager }}
Build Date   : {{ .PrettyBuildDate }}
Install Date : {{ .PrettyInstallDate }}
Package Size : {{ .Size }} bytes
Install Size : {{ .ISize }} bytes
MD5 Sum      : {{ .MD5Sum }}
SHA256 Sum   : {{ .SHA256Sum }}

Required By  : {{ .ComputeRequiredBy }}
Files        : {{ range .Files }}
               {{ .Name }} {{ .Size }}{{ end }}
`

var pkginfo_tpl *template.Template

type PrettyPackage struct {
	Package
}

func (p PrettyPackage) PrettyBuildDate() string {
	date := p.BuildDate()
	t := time.Unix(date, 0)
	return t.Format(time.RFC1123)
}

func (p PrettyPackage) PrettyInstallDate() string {
	date := p.InstallDate()
	t := time.Unix(date, 0)
	return t.Format(time.RFC1123)
}

func init() {
	var er error
	pkginfo_tpl, er = template.New("info").Parse(pkginfo_template)
	if er != nil {
		fmt.Printf("couldn't compile template: %s\n", er)
		panic("template parsing error")
	}
}

// Tests package attribute getters.
func TestPkginfo(t *testing.T) {
	h, er := Init(root, dbpath)
	defer h.Release()
	if er != nil {
		t.Errorf("Failed at alpm initialization: %s", er)
	}

	t.Log("Printing package information for pacman")
	db, _ := h.LocalDb()

	pkg, _ := db.GetPkg("pacman")
	buf := bytes.NewBuffer(nil)
	pkginfo_tpl.Execute(buf, PrettyPackage{*pkg})
	t.Logf("%s...", buf.Bytes()[:1024])

	pkg, _ = db.GetPkg("linux")
	if pkg != nil {
		buf = bytes.NewBuffer(nil)
		pkginfo_tpl.Execute(buf, PrettyPackage{*pkg})
		t.Logf("%s...", buf.Bytes()[:1024])
	}
}
