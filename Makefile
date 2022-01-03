build:
	go build -o neuBi ./assembler
	go build -o neuVM ./interpreter

test:
	go test ./...

bench:
	go test ./core -run=all -bench=.
