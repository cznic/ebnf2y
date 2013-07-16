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
	go tool yacc -o $@ $<

demo.y: demo.ebnf
	ebnf2y -start Expression -o $@ $<
