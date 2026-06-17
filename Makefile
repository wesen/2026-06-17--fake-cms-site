.PHONY: build run test tidy clean seed serve

BIN ?= fake-cms

build:
	go build -o $(BIN) ./cmd/fake-cms

# build-all excludes the ttmp/ scripts (standalone package main repros).
build-all:
	go build $$(go list ./... | grep -v '/ttmp/')

run: build
	./$(BIN)

test:
	go test $$(go list ./... | grep -v '/ttmp/') -count=1

tidy:
	go mod tidy

seed: build
	./$(BIN) seed --db-path testdata/cms.db

serve: build
	./$(BIN) serve --db-path testdata/cms.db --addr :8080

clean:
	rm -f $(BIN) coverage.out
