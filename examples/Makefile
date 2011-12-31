all: alpm-installed alpm-search

alpm-installed: installed.go
	go build -x -o $@ $<

alpm-search: search.go
	go build -x -o $@ $<

clean:
	rm -f alpm-installed alpm-search
