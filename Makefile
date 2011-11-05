include $(GOROOT)/src/Make.inc

TARG=github.com/remyoudompheng/go-alpm
CGOFILES=alpm.go\
	 handle.go\
	 error.go\
	 list.go\
	 db.go\
	 package.go\
	 types.go

GOFILES=defs.go\
	enums.go

include $(GOROOT)/src/Make.pkg

examples:
	gomake -C examples

format:
	gofmt -l -s -w *.go
