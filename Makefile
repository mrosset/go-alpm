# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

.PHONY: all alpm clean

all: alpm

alpm:
	gomake -C alpm

install: alpm
	gomake -C alpm install

examples:
	gomake -C examples

clean:
	gomake -C alpm clean

test:
	gomake -C alpm test
