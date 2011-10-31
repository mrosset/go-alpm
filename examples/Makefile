include $(GOROOT)/src/Make.inc

PROG?=installed

DEPS=..

TARG=alpm-$(PROG)

GOFILES=$(PROG).go

include $(GOROOT)/src/Make.cmd

examples:
	$(MAKE) PROG=installed
	$(MAKE) PROG=search

format:
	gofmt -l -s -w *.go

