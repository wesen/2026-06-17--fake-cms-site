.PHONY: build run test tidy clean seed serve

BIN ?= fake-cms

build:
	go build -o $(BIN) ./cmd/fake-cms

run: build
	./$(BIN)

test:
	go test ./... -count=1

tidy:
	go mod tidy

seed: build
	./$(BIN) seed --db-path testdata/cms.db

serve: build
	./$(BIN) serve --db-path testdata/cms.db --addr :8080

clean:
	rm -f $(BIN) coverage.out
