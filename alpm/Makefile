include $(GOROOT)/src/Make.inc

TARG     = alpm
CGOFILES = alpm.go alpm_list.go db.go

CGO_CFLAGS  = -L/usr/include/ -L./
CGO_LDFLAGS = -lalpm

CLEANFILES+=

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O

format:
	gofmt -l -s -w -tabindent -tabwidth=4 *.go
