include $(GOROOT)/src/Make.inc

TARG=github.com/str1ngs/go-alpm
CGOFILES=alpm.go list.go db.go package.go

include $(GOROOT)/src/Make.pkg

examples:
	gomake -C examples

format:
	gofmt -l -w *.go
