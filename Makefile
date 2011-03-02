include $(GOROOT)/src/Make.inc

TARG=github.com/str1ngs/go-alpm
CGOFILES=alpm.go list.go db.go package.go
CGO_LDFLAGS=-lalpm

include $(GOROOT)/src/Make.pkg

examples:
	gomake -C examples

format:
	gofmt -l -spaces=true -tabindent=false -tabwidth=2 -w *.go
