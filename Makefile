# Copyright 2014 The ebnf2y Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY: all clean demo

all: editor
	go vet
	go install
	make todo

editor:
	go fmt
	go test -i
	go test
	go build

todo:
	@grep -n ^[[:space:]]*_[[:space:]]*=[[:space:]][[:alpha:]][[:alnum:]]* *.go || true
	@grep -n TODO *.go || true
	@grep -n BUG *.go || true
	@grep -n println *.go || true

clean:
	@go clean
	rm -f *~ y.output
	make -C demo clean

demo:
	make -C demo
