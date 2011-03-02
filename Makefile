include $(GOROOT)/src/Make.inc

TARG     = alpm
CGOFILES = alpm.go list.go db.go package.go

CGO_CFLAGS  = -L/usr/include/ -L./
CGO_LDFLAGS = -lalpm

CLEANFILES+=

include $(GOROOT)/src/Make.pkg

#%: install %.go
#	$(GC) $*.go
#	$(LD) -o $@ $*.$O

examples:
	gomake -C examples

format:
	gofmt -l -s -w -tabindent -tabwidth=2 *.go
