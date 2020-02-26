all: cli

cli:
	pkger -o ./cmd/compareto
	go build ./cmd/compareto

run:
	./compareto

clean:
	go clean ./cmd/compareto