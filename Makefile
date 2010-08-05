include $(GOROOT)/src/Make.$(GOARCH)

TARG     = alpm
CGOFILES = alpm.go alpm_list.go

CGO_CFLAGS  = -L/usr/include/ -L./
CGO_LDFLAGS = -lalpm

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O

format:
	gofmt -w=true -tabindent=false -tabwidth=2 *.go
