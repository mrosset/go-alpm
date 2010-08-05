include $(GOROOT)/src/Make.$(GOARCH)

TARG     = list
CGOFILES = list.go

CGO_CFLAGS  = -L/usr/include/
CGO_LDFLAGS = -lalpm

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
