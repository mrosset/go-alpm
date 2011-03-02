include $(GOROOT)/src/Make.inc

all: installed search

installed: installed.go
	$(GC) installed.go
	$(LD) -o $@ installed.$(O)

search: search.go
	$(GC) search.go
	$(LD) -o $@ search.$(O)

clean:
	rm -f -r *.8 *.6 *.o */*.8 */*.6 */*.o */_obj installed search

format:
	gofmt -l -s -w -tabindent -tabwidth=4 *.go

test: all
