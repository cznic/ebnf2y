# Copyright 2014 The ebnf2y Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY: all clean

all: scanner.go parser.go
	go fmt
	go build
	./demo

clean:
	@go clean
	rm -f y.output

scanner.go: demo.l
	golex -o $@ $<

parser.go: demo.y
	goyacc -o $@ $<

demo.y: demo.ebnf
	ebnf2y -start Expression -o $@ $<
