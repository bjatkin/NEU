build:
	go build -o neuBi ./assembler
	go build -o neuClient ./interpreter

test:
	go test ./...