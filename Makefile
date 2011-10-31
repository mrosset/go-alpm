include $(GOROOT)/src/Make.inc

TARG=alpm
CGOFILES=alpm.go\
	 error.go\
	 list.go\
	 db.go\
	 package.go

GOFILES=defs.go

include $(GOROOT)/src/Make.pkg

defs.go: defs.c
	godefs -g alpm $< | sed s/ALPM_// | gofmt > $@

examples:
	gomake -C examples

format:
	gofmt -l -w *.go
