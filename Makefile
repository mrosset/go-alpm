include $(GOROOT)/src/Make.inc

TARG=alpm
CGOFILES=alpm.go\
	 error.go\
	 list.go\
	 db.go\
	 package.go

include $(GOROOT)/src/Make.pkg

examples:
	gomake -C examples

format:
	gofmt -l -w *.go
