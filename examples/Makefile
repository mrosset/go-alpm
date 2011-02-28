include $(GOROOT)/src/Make.inc

TARG=installed
GOFILES=installed.go

include $(GOROOT)/src/Make.cmd

test: format all
	./installed

format:
	gofmt -l -s -w -tabindent -tabwidth=4 *.go
