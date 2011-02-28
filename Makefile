include $(GOROOT)/src/Make.inc

TARG     = alpm
CGOFILES = alpm.go alpm_list.go

CGO_CFLAGS  = -L/usr/include/ -L./
CGO_LDFLAGS = -lalpm

CLEANFILES+=main

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O

main: main.c
	gcc -lalpm main.c -o main
format:
	gofmt -l -s -w -tabindent -tabwidth=4 *.go

all: main format
